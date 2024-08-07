package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	spinner       spinner.Model
	projectPath   string
	htmlID        string
	findString    string
	replaceString string
	done          bool
	err           error
}

func initialModel() model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	return model{spinner: s}
}

func (m model) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	case processFinishedMsg:
		m.done = true
		return m, tea.Quit
	case errorMsg:
		m.err = msg
		return m, tea.Quit
	}
	return m, nil
}

func (m model) View() string {
	if m.err != nil {
		return fmt.Sprintf("Error: %v\n", m.err)
	}
	if m.done {
		return "Processing complete. Check the log file in your home directory.\n"
	}
	return fmt.Sprintf("%s Processing files...\n", m.spinner.View())
}

type processFinishedMsg struct{}
type errorMsg error

func main() {
	if len(os.Args) != 5 {
		fmt.Println("Usage: chainsaw <project_folder> <html_id> <find_string> <replace_string>")
		os.Exit(1)
	}

	m := initialModel()
	m.projectPath = os.Args[1]
	m.htmlID = os.Args[2]
	m.findString = os.Args[3]
	m.replaceString = os.Args[4]

	p := tea.NewProgram(m)

	go func() {
		err := processFiles(m.projectPath, m.htmlID, m.findString, m.replaceString)
		if err != nil {
			p.Send(errorMsg(err))
		} else {
			p.Send(processFinishedMsg{})
		}
	}()

	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}

func processFiles(projectPath, htmlID, findString, replaceString string) error {
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	logFileName := fmt.Sprintf("chainsaw_log-%s.txt", timestamp)
	logFile, err := os.Create(filepath.Join(os.Getenv("HOME"), logFileName))
	if err != nil {
		return err
	}
	defer logFile.Close()

	return filepath.Walk(projectPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		return processFile(path, htmlID, findString, replaceString, logFile)
	})
}

func processFile(filePath, htmlID, findString, replaceString string, logFile *os.File) error {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	lines := strings.Split(string(content), "\n")
	insideID := false
	modified := false

	for i, line := range lines {
		if strings.Contains(line, htmlID) {
			insideID = true
		}
		if insideID && strings.Contains(line, findString) {
			lines[i] = strings.ReplaceAll(line, findString, replaceString)
			logChange(logFile, filePath, i+1)
			modified = true
		}
		if insideID && strings.Contains(line, ">") {
			insideID = false
		}
	}

	if modified {
		return ioutil.WriteFile(filePath, []byte(strings.Join(lines, "\n")), 0644)
	}
	return nil
}

func logChange(logFile *os.File, filePath string, lineNumber int) {
	logEntry := fmt.Sprintf("Changed file: %s, line: %d\n", filePath, lineNumber)
	logFile.WriteString(logEntry)
	fmt.Print(logEntry)
}
