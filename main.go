package main

import (
	"fmt"
	"os"
	"strings"
)

// main function: entry point of the program
func main() {
	// Validate command-line arguments
	inputString, err := validateInput()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(inputString) == 0 {
		return
	}
	if inputString == "\\n" {
		fmt.Println()
		return
	}

	// Read the standard.txt file
	baseFormat, err := readFile("standard.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	// Process and print ASCII art for the input string
	printASCIIArt(inputString, baseFormat)
}

// validateInput function: checks if the command-line input is valid
func validateInput() (string, error) {
	if len(os.Args) != 2 {
		return "", fmt.Errorf("the format of input is incorrect. you should use <go run . \"input\">")
	}
	return os.Args[1], nil
}

// readFile function: reads the content of a file and returns it as a string
func readFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// printASCIIArt function: converts the input string to ASCII art and prints it
func printASCIIArt(inputString string, baseFormat string) {
	// Constants for ASCII character properties
	const asciiHeight = 8  // Number of rows for each ASCII character
	const asciiOffset = 32 // ASCII offset for printable characters

	// Split input into lines based on "\\n" and base format into lines
	inputLines := strings.Split(inputString, "\\n")
	asciiLines := strings.Split(baseFormat, "\n")

	// Iterate over each line of input
	for _, inputLine := range inputLines {
		if inputLine == "" {
			fmt.Println() // Print an empty line for blank input lines
			continue
		}

		// For each row of the ASCII character representation (1 to 8)
		for row := 1; row <= asciiHeight; row++ {
			var lineData strings.Builder

			// Process each character in the input line
			for _, char := range inputLine {
				asciiIndex := int(char) - asciiOffset // Calculate the index for the character in ASCII art
				lineNumber := (asciiIndex * (asciiHeight + 1)) + row

				// Append the corresponding line of the ASCII character to the output
				if lineNumber < len(asciiLines) {
					lineData.WriteString(asciiLines[lineNumber])
				}
			}

			// Print the resulting line if it contains any data
			if lineData.Len() > 0 {
				fmt.Println(lineData.String())
			}
		}
	}
}
