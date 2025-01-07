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
				if checkDirection("N", input, searchPhrase, i, j) {
					count++
				}
			}

			// ne
			if canGoUp && canGoRight {
				if checkDirection("NE", input, searchPhrase, i, j) {
					count++
				}
			}
			// e
			if canGoRight {
				if checkDirection("E", input, searchPhrase, i, j) {
					count++
				}
			}
			// se
			if canGoDown && canGoRight {
				if checkDirection("SE", input, searchPhrase, i, j) {
					count++
				}
			}
			// s
			if canGoDown {
				if checkDirection("S", input, searchPhrase, i, j) {
					count++
				}
			}
			// sw
			if canGoDown && canGoLeft {
				if checkDirection("SW", input, searchPhrase, i, j) {
					count++
				}
			}
			// w
			if canGoLeft {
				if checkDirection("W", input, searchPhrase, i, j) {
					count++
				}
			}
			// nw
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
