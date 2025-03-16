# Simple Text Editor

A lightweight command-line text editor written in Go.

## Features
- Read and edit text files from the terminal
- Simple interface with minimal commands
- Supports any text file format

## Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/cli-txt-editor.git

# Navigate to the project directory
cd cli-txt-editor

# Build the project
go build -o editor
```

## Usage

```bash
# Open a file for editing
./editor filename.txt
```

## How to Use
1. Run the program with a filename as argument
2. If the file exists, its contents will be displayed
3. Enter new text for the file (each line is entered separately)
4. Type 'SAVE' on a new line to save and exit

## Example

```
$ ./editor notes.txt
Welcome to text editor!
Editing file: notes.txt
[content of notes.txt displayed here]
Enter new text for file (Type 'SAVE' in new line to save the file)
This is a new line for my notes
Another line
SAVE
File saved successfully.
```

## Dependencies
- Standard Go libraries only

## License
MIT
