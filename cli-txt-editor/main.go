package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	fmt.Println("Welcome to text editor! ")
	if len(os.Args) < 2 {
		fmt.Println("Please provide a filename")
		os.Exit(1)
	}

	filename := os.Args[1]
	fmt.Printf("Editing file: %s\n ", filename)

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading this file: ", err)
		os.Exit(1)
	}
	fmt.Println(string(contents))

	fmt.Println("Enter new text for file (Type 'SAVE' in new line to save the file)")
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		if line == "SAVE" {
			break
		}
		lines = append(lines, line)
	}
	newContents := []byte{}

	for _, line := range lines {
		newContents = append(newContents, []byte(line)...)
		newContents = append(newContents, '\n')
	}
	err = ioutil.WriteFile(filename, newContents, 0644)
	if err != nil {
		fmt.Println("Error writing to file: ", err)
		os.Exit(1)
	}
	fmt.Println("File saved successfuly.")
}
