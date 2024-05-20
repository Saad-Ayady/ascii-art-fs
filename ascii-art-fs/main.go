package main

import (
	"fmt"
	"os"
	"strings"
)

// GetAscii returns the ASCII representation of a character based on an offset
// in the provided file content.
func GetAscii(char int, fileContent string) string {
	fileContent = strings.ReplaceAll(fileContent, "\r\n", "\n")
	asciiOffset := char - 32
	lines := strings.Split(fileContent, "\n\n")
	return lines[asciiOffset]
}

// VirInput parses a string, separating it into lines using "\n" as a delimiter.
func VirInput(input string) []string {
	for _, char := range input {
		if char < 32 || char > 126 {
			fmt.Println("Pleas Get a logic string (╯°□°)╯!")
			os.Exit(1)
		}
	}
	lines := strings.Split(input, "\\n")
	return lines
}

// printH prints the ASCII representation of text lines horizontally.
func printH(lines [][]string) {
	for i := 0; i < 8; i++ {
		for j := 0; j < len(lines); j++ {
			if i < len(lines[j]) {
				fmt.Print(lines[j][i])
			}
		}
		fmt.Println()
	}
}

// isEmpty checks if a slice of strings is empty.
func isEmpty(input []string) bool {
	for _, char := range input {
		if char != "" {
			return false
		}
	}
	return true
}

// printTxt prints the ASCII representation of text.
func printTxt(input, fileContent string) {
	var nStr [][]string
	lines := VirInput(input)
	for i, line := range lines {
		input = string(line)
		if input == "" {
			if i != 0 || !isEmpty(lines) {
				fmt.Println()
			}
			continue
		}
		for _, char := range line {
			asciiStr := GetAscii(int(char), fileContent)
			if asciiStr != "" {
				lines := strings.Split(asciiStr, "\n")
				nStr = append(nStr, lines)
			} else {
				nStr = append(nStr, []string{})
			}
		}
		printH(nStr)
		nStr = nil
	}
}

func main() {
	if len(os.Args) >= 4 || len(os.Args) < 2 {
		fmt.Println("Usage <3: go run . [STRING] [BANNER]")
		os.Exit(1)
	}
	styleBanner := ""
	if len(os.Args) == 3 {
		styleBanner = os.Args[2]
	} else {
		styleBanner = "standard"
	}

	if strings.Contains(styleBanner, ".txt") {
		styleBanner = string(styleBanner)
	} else {
		styleBanner = string(styleBanner + ".txt")
	}

	content, err := os.ReadFile(styleBanner)
	if err != nil {
		fmt.Printf("Error reading file: %v", err)
		return
	}
	content = content[1:]
	inputUser := os.Args[1]
	printTxt(inputUser, string(content))
}
