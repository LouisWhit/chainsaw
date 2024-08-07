
<pre>
   ___ _           _                           
  / __\ |__   __ _(_)_ __  ___  __ ___      __ 
 / /  | '_ \ / _` | | '_ \/ __|/ _` \ \ /\ / / 
/ /___| | | | (_| | | | | \__ \ (_| |\ V  V /  
\____/|_| |_|\__,_|_|_| |_|___/\__,_| \_/\_/   

</pre>
<p align="center">
  <strong>Slice through your HTML with precision!</strong>
</p>


<p align="center">
  <a href="#installation">Installation</a> •
  <a href="#usage">Usage</a> •
  <a href="#features">Features</a> •
  <a href="#contributing">Contributing</a> •
  <a href="#license">License</a>
</p>

---

## 🌟 About

Chainsaw is a powerful command-line tool that allows you to make surgical changes to your HTML files across entire project directories. With Chainsaw, you can easily find and replace text within specific HTML elements, saving you time and reducing the risk of manual errors.

## 🚀 Features

- 🔍 Search for specific HTML IDs across multiple files
- 🔄 Find and replace text within targeted HTML elements
- 📁 Recursively process entire project directories
- 📝 Detailed logging of all changes
- 💅 Stylish CLI interface powered by [Bubble Tea](https://github.com/charmbracelet/bubbletea)

## 🛠 Installation

### macOS

Run the following command in your terminal:

```bash
curl -sSL https://raw.githubusercontent.com/yourusername/chainsaw/main/install_chainsaw.sh | bash
```
### Other platforms

Ensure you have Go installed (version 1.16+)
Run:
Copygo install github.com/yourusername/chainsaw@latest


## 🔧 Usage
bashCopychainsaw <project_folder> <html_id> <find_string> <replace_string>
Example:
bashCopychainsaw ./my-project "header-nav" "Contact Us" "Get in Touch"
This command will:

Search through all files in ./my-project
Find elements with the ID header-nav
Replace "Contact Us" with "Get in Touch" within those elements

## 📊 Output
Chainsaw provides real-time feedback in your terminal and generates a detailed log file in your home directory:
Copy~/chainsaw_log-2024-08-07_14-30-45.txt

## 🤝 Contributing
Contributions, issues, and feature requests are welcome! Feel free to check the issues page.

## 📜 License
This project is good luck licensed.

