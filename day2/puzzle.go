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
	var conditionalSafeCounter int = 0

	countSafeRecords(inputFile, &safeCounter, &conditionalSafeCounter)

	fmt.Println(safeCounter)

}

func countSafeRecords(filename string, safeCounter *int, conditionalSafeCounter *int) {
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
		} else if isReportSafe(line[:len(line)-1]) || isReportSafe(line[1:len(line)]) {
			*conditionalSafeCounter += 1
		} else if isReportSafe2(line) {
			*conditionalSafeCounter += 1
		}

	}
}
// first puzzle, levels in report must be in order (asc or desc)
// diff between levels must be between 1 and 3

func isReportSafe(report []int) bool {
	var startOrder bool = report[1] - report[0] > 0
	var currentOrder bool = true // true = asc; false = desc

	diff := 0 
	for i := 1; i < len(report); i++ {
		diff = report[i]-report[i-1]

		
		if diff == 0 || diff > 3 || diff < -3 {
			return false
		}

		if diff < 0 {
			currentOrder = false
		} else if diff > 0 {
			currentOrder = true
		}

		if startOrder != currentOrder {
			return false
		}
	}

	return true
}

/* 
 second puzzle
 same as previous, but
 if excluding one level makes it safe, whole report is safe
 edge cases: first and last level in report are tricky to handle

 idea: if not safe, truncate from start and/or end to check if any of those is safe
 if not, execute modified algo
 how to not do too many checks tho
 i could probably do like in first algo, but not return immediatelly, instead count cases when order is wrong
 or diff to big; if its only one, its fine;
 but would have to compare order/diff ommiting one value, so holding 3 in memory
*/
func isReportSafe2(report []int) bool {
	var startOrder bool = report[1] - report[0] > 0
	var currentOrder bool = true // true = asc; false = desc
	var offenceCounter int = 0
	
	diff := 0 
	for i := 1; i < len(report); i++ {
		diff = report[i]-report[i-1]

		
		if diff == 0 || diff > 3 || diff < -3 {
			return false
		}

		if diff < 0 {
			currentOrder = false
		} else if diff > 0 {
			currentOrder = true
		}

		if startOrder != currentOrder {
			return false
		}
	}

	return true
}
