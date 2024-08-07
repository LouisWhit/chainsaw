#!/bin/bash

# Chainsaw Installation Script for Mac

# Set the GitHub repository URL
REPO_URL="https://github.com/LouisWhit/chainsaw"

# Check if Go is installed
if ! command -v go &> /dev/null
then
    echo "Go is not installed. Please install Go first."
    echo "Visit https://golang.org/dl/ for installation instructions."
    exit 1
fi

# Create a temporary directory
TEMP_DIR=$(mktemp -d)
echo "Created temporary directory: $TEMP_DIR"

# Clone the repository
echo "Cloning the Chainsaw repository..."
git clone $REPO_URL $TEMP_DIR
if [ $? -ne 0 ]; then
    echo "Failed to clone the repository. Please check the URL and your internet connection."
    exit 1
fi

# Navigate to the cloned directory
cd $TEMP_DIR

# Initialize Go module
go mod init chainsaw

# Install dependencies
echo "Installing dependencies..."
go get github.com/charmbracelet/bubbles/spinner
go get github.com/charmbracelet/bubbletea

# Build the application
echo "Building Chainsaw..."
go build -o chainsaw

# Move the binary to /usr/local/bin
echo "Installing Chainsaw..."
sudo mv chainsaw /usr/local/bin/

# Clean up
cd ..
rm -rf $TEMP_DIR

echo "Chainsaw has been successfully installed!"
echo "You can now use it by running 'chainsaw' in your terminal."
echo "Usage: chainsaw <project_folder> <html_id> <find_string> <replace_string>"