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
	fmt.Println(conditionalSafeCounter)
	fmt.Println(safeCounter + conditionalSafeCounter)

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
		// fmt.Println(numbers)

		for _, el := range numbers {
			val1, _ := strconv.Atoi(el)
			line = append(line, val1)
		}
		if isReportSafe(line) {
			*safeCounter += 1
		} else if isReportConditionalySafe(line) {
			// fmt.Println("ConditionalySafe: ", numbers)
			*conditionalSafeCounter += 1
		} else {
			fmt.Println("Not Safe: ", line)
		}

	}
}

// first puzzle, levels in report must be in order (asc or desc)
// diff between levels must be between 1 and 3

func isReportSafe(report []int) bool {
	var startOrder bool = report[1]-report[0] > 0
	var currentOrder bool = true // true = asc; false = desc

	diff := 0
	for i := 1; i < len(report); i++ {
		diff = report[i] - report[i-1]

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
func isReportConditionalySafe(report []int) bool {

	// check removing first and last element, might not be most efficient

	if isReportSafe(report[:len(report)-1]) || isReportSafe(report[1:]) {
		// fmt.Println("Condtitionaly safe by cutting first or last")
		return true
	}

	// assures what order should report be based on first 2 elements
	// if elements are equal, one of them must be removed for report to be safe
	// removing first element was checked before
	// @ values: true = asc; false = desc

	var startOrder bool

	if report[1] == report[0] {
		return isReportSafe(append(report[:1], report[2:]...))
	} else {
		startOrder = report[1]-report[0] > 0
	}

	var currentOrder bool = startOrder
	var errorCounter int = 0
	var isCurrentValueCorrect bool
	var diff int

	for i := 1; i < len(report); i++ {
		diff = report[i] - report[i-1]
		currentOrder = diff > 0
		isCurrentValueCorrect = true

		if !isDiffGood(diff) || startOrder != currentOrder {
			isCurrentValueCorrect = false
		}

		if !isCurrentValueCorrect && i < len(report)-1 {
			diffPrevToNext := report[i+1] - report[i-1]
			if !isDiffGood(diffPrevToNext) {
				return isReportSafe(append(report[:i-1], report[i:]...))
			}

			orderPrevToNext := report[i+1]-report[i-1] > 0
			if startOrder != orderPrevToNext {
				return isReportSafe(append(report[:i-1], report[i:]...))
			}
		}

		if !isCurrentValueCorrect {
			errorCounter += 1

			if errorCounter > 1 {
				fmt.Println("Stopped at index: ", i, " with value", report[i], ": error count > 1")
				return false
			}

			i += 1
		}
	}

	return true

}

func isDiffGood(diff int) bool {
	return !(diff == 0 || diff > 3 || diff < -3)
}
