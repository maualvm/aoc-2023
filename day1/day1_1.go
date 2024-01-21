package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

// getLinesFromFile reads a file and returns a slice of strings, each string
// being a line from the file.
func getLinesFromFile(fileName string) ([]string, error) {
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("There was an error opening the file:", err)
		return nil, err
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("There was an error scanning the file:", err)
		return nil, err
	}

	return lines, nil
}

// getFirstDigit takes a string and returns the first digit in the string. It
// returns -1 if no digit is found.
func getFirstDigit(line string) rune {
	for _, char := range line {
		if unicode.IsDigit(char) {
			return char
		}
	}

	return -1
}

// getLastDigit takes a string and returns the last digit in the string. It
// returns -1 if no digit is found.
func getLastDigit(line string) rune {
	for i := len(line) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(line[i])) {
			return rune(line[i])
		}
	}

	return -1
}

func main() {

	if len(os.Args) < 2 {
		fmt.Printf("Usage: go run day1_1.go <filename>\n")
		return
	}

	filename := os.Args[1]
	lines, err := getLinesFromFile(filename)

	if err != nil {
		return
	}

	totalSum := 0
	for _, line := range lines {
		stringInLine := string(getFirstDigit(line)) + string(getLastDigit(line))
		numberInLine, _ := strconv.Atoi(stringInLine)
		totalSum += numberInLine
	}

	fmt.Println("Total sum:", totalSum)
}
