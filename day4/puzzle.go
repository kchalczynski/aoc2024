package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {

	var inputFile string = "input.txt"
	var searchPhrase string = "XMAS"
	var inputContent string
	var inputLinesCount int

	inputContent, inputLinesCount = readFile(inputFile)

	// fmt.Println(inputContent, inputLinesCount)
	var wordCount int = countWordOccurrence(inputContent, inputLinesCount, searchPhrase)

	fmt.Println("Word `", searchPhrase, "` appears: ", wordCount, " times")
	var countP2 int = countXShapedMASOccurrence(inputContent, inputLinesCount)

	fmt.Println("MAS IN SHAPE OF X CROSS appears: ", countP2, " times")
}

func countWordOccurrence(input string, linesCount int, searchPhrase string) int {

	letterMatrix := readContentIntoMatrix(input, linesCount)

	count := searchWordsInMatrix(letterMatrix, searchPhrase)

	return count
}

func countXShapedMASOccurrence(input string, linesCount int) int {

	letterMatrix := readContentIntoMatrix(input, linesCount)

	count := searchWordsInMatrix2(letterMatrix)

	return count
}

func readContentIntoMatrix(input string, linesCount int) [][]string {

	letterMatrix := make([][]string, linesCount)

	for i, val := range strings.Split(input, "\n") {
		letterMatrix[i] = strings.Split(val, "")
	}
	return letterMatrix
}

func searchWordsInMatrix(input [][]string, searchPhrase string) int {

	// search by directions: east/west/south/north and diagonals
	// possible directions: n; ne; e; se; s; sw; w; nw;
	// limiting index: len = len(searchPhrase) ; n[len][]; e[][size-len];s[size-len][];w[][len]
	// diagonals combine both requirements

	var count int = 0

	var canGoUp bool = false
	var canGoDown bool = false
	var canGoRight bool = false
	var canGoLeft bool = false

	phraseLength := len(searchPhrase)
	rows := len(input)
	columns := len(input[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {

			if input[i][j] != string(searchPhrase[0]) {
				continue
			}

			canGoUp = i >= phraseLength-1
			canGoDown = i <= rows-phraseLength
			canGoRight = j <= columns-phraseLength
			canGoLeft = j >= phraseLength-1

			if canGoUp {
				if checkDirection("N", input, searchPhrase, i, j) {
					count++
				}
			}

			if canGoUp && canGoRight {
				if checkDirection("NE", input, searchPhrase, i, j) {
					count++
				}
			}

			if canGoRight {
				if checkDirection("E", input, searchPhrase, i, j) {
					count++
				}
			}

			if canGoDown && canGoRight {
				if checkDirection("SE", input, searchPhrase, i, j) {
					count++
				}
			}

			if canGoDown {
				if checkDirection("S", input, searchPhrase, i, j) {
					count++
				}
			}

			if canGoDown && canGoLeft {
				if checkDirection("SW", input, searchPhrase, i, j) {
					count++
				}
			}

			if canGoLeft {
				if checkDirection("W", input, searchPhrase, i, j) {
					count++
				}
			}

			if canGoUp && canGoLeft {
				if checkDirection("NW", input, searchPhrase, i, j) {
					count++
				}
			}

		}
	}

	return count
}

func checkDirection(direction string, input [][]string, searchPhrase string, row int, column int) bool {

	var dir_x, dir_y int = 0, 0

	switch {
	case direction == "N":
		dir_x = -1
	case direction == "NE":
		dir_x = -1
		dir_y = 1
	case direction == "E":
		dir_y = 1
	case direction == "SE":
		dir_x = 1
		dir_y = 1
	case direction == "S":
		dir_x = 1
	case direction == "SW":
		dir_x = 1
		dir_y = -1
	case direction == "W":
		dir_y = -1
	case direction == "NW":
		dir_x = -1
		dir_y = -1
	}

	for i, el := range string(searchPhrase) {
		if input[row+i*dir_x][column+i*dir_y] != string(el) {
			return false
		}
	}
	return true
}

// Part2, X = shape, "A" in the middle, "M" and "S" on either side diagonally
// but both diagonals must have M and S

func searchWordsInMatrix2(input [][]string) int {
	var count int = 0
	rows := len(input)
	columns := len(input[0])

	//check for "A"'s, last/first row/column excluded
	for i := 1; i < rows-1; i++ {
		for j := 1; j < columns-1; j++ {
			if input[i][j] != "A" {
				continue
			}

			if searchDiagonals(input, i, j) {
				count++
			}

		}
	}

	return count
}

func searchDiagonals(input [][]string, row int, column int) bool {
	diag1 := []string{input[row-1][column-1], input[row+1][column+1]}
	diag2 := []string{input[row-1][column+1], input[row+1][column-1]}

	if slices.Contains(diag1, "M") && slices.Contains(diag1, "S") &&
		slices.Contains(diag2, "M") && slices.Contains(diag2, "S") {
		return true
	}
	return false
}

func readFile(fileName string) (string, int) {
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error opening file: ", err)
		return "", 0
	}
	var content string = ""

	defer file.Close()

	var lineCount int = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineCount += 1
		content += line + "\n"
	}

	return strings.TrimSuffix(content, "\n"), lineCount
}
