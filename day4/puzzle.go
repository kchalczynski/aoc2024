package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	// "strconv"
)

func main() {

	var inputFile string = "input.txt"
	var inputContent string
	var inputLinesCount int
	inputContent, inputLinesCount = readFile(inputFile)

	fmt.Println(inputContent, inputLinesCount)
	searchPhrase := "XMAS"
	var wordCount int = countWordOccurence(inputContent, inputLinesCount, searchPhrase)

	fmt.Println("Word `", searchPhrase, "` appears: ", wordCount, " times")
}

func countWordOccurence(input string, linesCount int, searchPhrase string) int {

	letterMatrix := readContentIntoMatrix(input, linesCount)

	count := searchWordsInMatrix(letterMatrix, searchPhrase)

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

			// n
			if canGoUp {
				if checkN(input, searchPhrase, i, j) {
					count++
				}
			}

			// ne
			if canGoUp && canGoRight {
				if checkNE(input, searchPhrase, i, j) {
					count++
				}
			}
			// e
			if canGoRight {
				if checkE(input, searchPhrase, i, j) {
					count++
				}
			}
			// se
			if canGoDown && canGoRight {
				if checkSE(input, searchPhrase, i, j) {
					count++
				}
			}
			// s
			if canGoDown {
				if checkS(input, searchPhrase, i, j) {
					count++
				}
			}
			// sw
			if canGoDown && canGoLeft {
				if checkSW(input, searchPhrase, i, j) {
					count++
				}
			}
			// w
			if canGoLeft {
				if checkW(input, searchPhrase, i, j) {
					count++
				}
			}
			// nw
			if canGoUp && canGoLeft {
				if checkNW(input, searchPhrase, i, j) {
					count++
				}
			}

		}
	}

	return count
}

func checkN(input [][]string, searchPhrase string, row int, column int) bool {
	for i, el := range string(searchPhrase) {
		if input[row-i][column] != string(el) {
			return false
		}
	}
	return true
}

func checkNE(input [][]string, searchPhrase string, row int, column int) bool {
	for i, el := range string(searchPhrase) {
		if input[row-i][column+i] != string(el) {
			return false
		}
	}
	return true
}

func checkE(input [][]string, searchPhrase string, row int, column int) bool {
	for i, el := range string(searchPhrase) {
		if input[row][column+i] != string(el) {
			return false
		}
	}
	return true
}

func checkSE(input [][]string, searchPhrase string, row int, column int) bool {
	for i, el := range string(searchPhrase) {
		if input[row+i][column+i] != string(el) {
			return false
		}
	}
	return true
}

func checkS(input [][]string, searchPhrase string, row int, column int) bool {
	for i, el := range string(searchPhrase) {
		if input[row+i][column] != string(el) {
			return false
		}
	}
	return true
}

func checkSW(input [][]string, searchPhrase string, row int, column int) bool {
	for i, el := range string(searchPhrase) {
		if input[row+i][column-i] != string(el) {
			return false
		}
	}
	return true
}

func checkW(input [][]string, searchPhrase string, row int, column int) bool {
	for i, el := range string(searchPhrase) {
		if input[row][column-i] != string(el) {
			return false
		}
	}
	return true
}

func checkNW(input [][]string, searchPhrase string, row int, column int) bool {
	for i, el := range string(searchPhrase) {
		if input[row-i][column-i] != string(el) {
			return false
		}
	}
	return true
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
