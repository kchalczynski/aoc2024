package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	visitedMark     = "X"
	notVisitedMark  = "."
	obstructionMark = "#"
	dirN            = "^"
	dirE            = ">"
	dirS            = "v"
	dirW            = "<"
)

func readFile(fileName string) (string, int) {
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error opening file: ", err)
		return "", 0
	}
	var content = ""
	var lineCount = 0

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineCount += 1
		content += line + "\n"
	}

	return strings.TrimSuffix(content, "\n"), lineCount
}

func readContentIntoMatrix(input string, linesCount int) [][]string {

	letterMatrix := make([][]string, linesCount)

	for i, val := range strings.Split(input, "\n") {
		letterMatrix[i] = strings.Split(val, "")
	}
	return letterMatrix
}
