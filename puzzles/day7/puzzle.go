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

	operatorPermutations := make(map[int][]string)
	for _, v := range valMap {
		if operatorPermutations[len(v)] != nil {
			operatorPermutations[len(v)] = make([]string, 0, len(v)-1)
		}
	}

	generateAllOperators()

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

	// currOperators - I need them for each iteration/depth,
	// so somehow have to remove last/clean after reaching end?
	currOperators := make([]string, 0, len(operands)-1)
	generateOperators(len(operands)-2, &operators, currOperators, 0)

	return false
}

// TODO: 2 ideas:
//  1. create operators for every line
//  2. create map for number of operands in each line, so if there are multiple entries of 3/4/5 etc. operands
//     I only generate operators list once, before even testing if results are viable
func generateOperators(depth int, operators *[][]string, currOperators []string, index int) {
	// I need
	// 	a) depth - essentially operators left
	// 	b) operators string/array to add them in recursive calls;
	// 	c) iterator, so depending on depth I can add currOperators to operators 2d array on specific index
	for i, _ := range availableOperators {
		// will it behave like copy when recursive calls are "returning"
		// so I complete one full recursive call, get e.g. "+++"
		// then it goes level up, will it be "++" or still "+++"?
		// same with index, maybe assigning to new value/using that as argument to recursive call would work if this approach doesnt
		currOperators = append(currOperators, string(availableOperators[i]))
		index = index + i*int(math.Pow(2, float64(depth)))
		if depth > 0 {
			generateOperators(depth-1, operators, currOperators, index)
		} else {
			// maybe it would be better as a map...
			(*operators)[index] = currOperators
		}
	}
}
