package main

import (
	"os"
	"testing"
)

// TestValidateInput tests the validateInput function
func TestValidateInput(t *testing.T) {
	os.Args = []string{"cmd", "input"}
	input, err := validateInput()
	if err != nil || input != "input" {
		t.Errorf("Expected 'input', got '%s', error: %v", input, err)
	}

	os.Args = []string{"cmd"}
	_, err = validateInput()
	if err == nil {
		t.Error("Expected an error for missing argument, got none")
	}

	os.Args = []string{"cmd", "input", "extra"}
	_, err = validateInput()
	if err == nil {
		t.Error("Expected an error for too many arguments, got none")
	}
}

// TestReadFile tests the readFile function
func TestReadFile(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "testfile.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	content := "Hello, World!"
	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tmpFile.Close()

	readContent, err := readFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if readContent != content {
		t.Errorf("Expected content '%s', got '%s'", content, readContent)
	}

	_, err = readFile("non_existent_file.txt")
	if err == nil {
		t.Error("Expected an error for non-existent file, got none")
	}
}

// TestPrintASCIIArt tests the printASCIIArt function
func TestPrintASCIIArt(t *testing.T) {
	inputMap := map[string]string{
		"1":  "hello",
		"2":  "HELLO",
		"3":  "HeLlo HuMaN",
		"4":  "1Hello 2There",
		"5":  "Hello\\nThere",
		"6":  "Hello\\n\\nThere",
		"7":  "{Hello & There #}",
		"8":  "hello There 1 to 2!",
		"9":  "MaD3IrA&LiSboN",
		"10": "1a\"#FdwHywR&/()=",
		"11": "{|}~",
		"12": "[\\]^_ 'a",
		"13": "RGB",
		"14": ":;<=>?@",
		"15": "\\!\" #$%&'()*+,-./",
		"16": "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"17": "abcdefghijklmnopqrstuvwxyz",
		"18": "AaBbCcDdEe",
		"19": "zxcvb 33",
		"20": "A%@#",
		"21": "xc  3%^WSE",
	}

	baseFormat, err := readFile("standard.txt")
	if err != nil {
		t.Errorf("Error reading file: %v", err)
	}

	for inputIndex, inputString := range inputMap {
		// Create a temporary file to capture output
		tmpFile, err := os.CreateTemp("", "output.txt")
		if err != nil {
			t.Fatalf("Failed to create temp file: %v", err)
		}
		defer os.Remove(tmpFile.Name()) // Clean up after the test

		// Redirect stdout to the temporary file
		originalStdout := os.Stdout
		os.Stdout = tmpFile

		// Run the function to generate ASCII art
		printASCIIArt(inputString, baseFormat)

		// Reset stdout back to original
		os.Stdout = originalStdout
		tmpFile.Close()

		// Read the content from the temporary file
		outputBytes, err := os.ReadFile(tmpFile.Name())
		if err != nil {
			t.Fatalf("Failed to read temp file: %v", err)
		}
		output := string(outputBytes)

		// Read the expected output for comparison
		expectedOutputBytes, err := os.ReadFile("test_expected_results/" + inputIndex + ".txt")
		if err != nil {
			t.Fatalf("Failed to read expected output file: %v", err)
		}
		expectedOutput := string(expectedOutputBytes)

		// Compare the actual output with the expected output
		if output != expectedOutput {
			t.Errorf("Test case %s failed.\nExpected output:\n%v\nGot:\n%v", inputIndex, expectedOutput, output)
		}
	}
}
