package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadFile(fileName string) (string, int) {
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
		content += strings.TrimSpace(line) + "\n"
	}

	return strings.TrimSuffix(content, "\n"), lineCount
}

func SplitStringByLines(input string) []string {
	lines := strings.Split(input, "\n")
	return lines
}

func ReadIntoMatrixByCharacter(input string, linesCount int) [][]string {
	letterMatrix := make([][]string, linesCount)

	for i, val := range strings.Split(input, "\n") {
		letterMatrix[i] = strings.Split(val, "")
	}
	return letterMatrix
}

func ReadAsNumberSeq(input string) []int {
	numbers := make([]string, 0)

	for _, val := range strings.Split(input, " ") {
		numbers = append(numbers, val)
	}
	return StringsToInts(numbers)
}

func StringsToInts(input []string) []int {
	result := make([]int, len(input))
	for i, s := range input {
		result[i], _ = strconv.Atoi(s)
	}
	return result
}
