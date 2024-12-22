package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {

	var inputFile string = "input.txt"
	var inputContent string = readFile(inputFile)
	operandTupleList := cleanInput(inputContent)

	var sum int = calculateSumOfMultiplications(operandTupleList)

	fmt.Print(sum)

}

func readFile(fileName string) string {
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error opening file: ", err)
		return ""
	}
	var content string = ""

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		content += line
	}
	return content

}

func cleanInput(content string) [][]int {

	pattern := regexp.MustCompile(`mul\(\d+,\d+\)`)
	mulArray := pattern.FindAllString(content, -1)
	pattern2 := regexp.MustCompile(`\d+`)

	operandArray := make([][]int, len(mulArray))

	for i := range operandArray {
		operandArray[i] = make([]int, 2)
	}

	for i, val := range mulArray {
		var intOperands [2]int
		stringOperands := pattern2.FindAllString(val, 2)
		val1, _ := strconv.Atoi(stringOperands[0])
		intOperands[0] = val1
		val2, _ := strconv.Atoi(stringOperands[1])
		intOperands[1] = val2

		operandArray[i][0] = intOperands[0]
		operandArray[i][1] = intOperands[1]
	}

	return operandArray
}

func calculateSumOfMultiplications(operands [][]int) int {

	var sum int = 0

	for i := range operands {
		result := operands[i][0] * operands[i][1] 
		sum += result
	}

	return sum
}
