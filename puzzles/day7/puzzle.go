package day7

import (
	"aoc2024/internal/utils"
	"github.com/kr/pretty"
	"math"
)

var valMap = map[int][]int{}
var availableOperators = "+*"

func Solve() {
	var inputFile = "puzzles/day7/test1.txt"
	inputContent, inputLength := utils.ReadFile(inputFile)

	valMap = splitInputToMap(utils.SplitStringByLines(inputContent))

	testVals := make([]int, 0, inputLength)
	for k := range valMap {
		testVals = append(testVals, k)
	}

	pretty.Println(valMap)
	posResults := checkForResults(&testVals)

	sumValues(&posResults, &testVals)
}

func sumValues(posResults *[]int, testVals *[]int) int {
	sum := 0
	for _, v := range *posResults {
		sum += (*testVals)[v]
	}

	return sum
}

func checkForResults(testVals *[]int) []int {
	posResults := make([]int, 0, len(*testVals))
	for i, v := range *testVals {
		if !checkForResult(v) {
			posResults = append(posResults, i)
		}
	}

	return posResults
}

func checkForResult(testVal int) bool {
	operands := valMap[testVal]
	operators := make([][]string, int(math.Pow(2, float64(len(operands)-1))))
	for i, _ := range operators {
		operators[i] = make([]string, len(operands)-1)
	}

	generateOperators(len(operands)-1, &operators)

}

func generateOperators(operatorsLeft int, operators *[][]string, currOperators []string) {
	// I need
	// 	a) depth
	// 	b) operators string/array to add them in recursive calls;
	// 	c) iterator, so depending on depth I can add currOperators to operators 2d array on specific index
	for i, _ := range availableOperators {
		if operatorsLeft > 0 {
			generateOperators(operatorsLeft-1, operators)
		}
		if (i+1)%2 != 0 {

		}
	}
}
