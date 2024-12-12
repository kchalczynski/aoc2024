package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var inputFile string = "input.txt"
	var safeCounter int = 0

	readFileIntoLists(inputFile, &safeCounter)

}

func readFileIntoLists(filename string, safeCounter *int) {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := []int{}

		numbers := strings.Fields(scanner.Text())
		fmt.Println(numbers)

		for _, el := range numbers {
			val1, _ := strconv.Atoi(el)
			line = append(line, val1)
		}

		if isReportSafe(line) {
			*safeCounter += 1
		}

	}
}

func isReportSafe(report []int) bool {

	return true
}
