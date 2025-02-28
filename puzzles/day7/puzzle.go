package day7

import (
	"aoc2024/internal/utils"
	"fmt"
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

	// For each input length, store 2d slice of all possible operator permutations in map, so it is generated only once
	operatorPermutations := make(map[int][][]string)
	initOperatorPermutationMap(operatorPermutations, valMap)

	pretty.Println(valMap)
	generateAllOperators(operatorPermutations)
	pretty.Println(operatorPermutations)

	//posResults := checkForResults(&testVals)
	//
	//sumValues(&posResults, &testVals)
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

func initOperatorPermutationMap(opMap map[int][][]string, valMap map[int][]int) {
	for _, v := range valMap {
		if opMap[len(v)] == nil {
			opMap[len(v)] = make([][]string, int(math.Pow(2, float64(len(v)-1))))
			for i, _ := range opMap[len(v)] {
				opMap[len(v)][i] = make([]string, len(v)-1)
			}
		}

	}
}

func generateAllOperators(operatorPermutations map[int][][]string) {
	for k, v := range operatorPermutations {
		currOperators := make([]string, 0, k-1)
		// operationPermutations[2] <-- possible operator permutations for 2 OPERANDS,
		//	maybe would be better to map it for X operators instead
		generateOperators(k-2, &v, currOperators, 0)
	}
}

// 	Create map for number of operands in each line, so if there are multiple entries of 3/4/5 etc. operands
//     I only generate operators list once, before even testing if results are viable

// I need
// 	a) depth - essentially operators left, it would be either first to last or last to first,
// 		but it doesn't matter as long as order is preserved across all permutations
// 	b) operators string/array to add them in recursive calls;
// 	c) iterator, so depending on depth I can add currOperators to operators 2d array on specific index

func generateOperators(depth int, operators *[][]string, currOperators []string, operatorIdx int) {
	tempOperators := make([]string, 0, cap(currOperators))
	for i, _ := range availableOperators {
		fmt.Println(i, string(availableOperators[i]))
	}
	for i, _ := range availableOperators {

		// will it behave like copy when recursive calls are "returning"
		// so I complete one full recursive call, get e.g. "+++"
		// then it goes level up, will it be "++" or still "+++"?
		// same with index, maybe assigning to new value/using that as argument to recursive call would work if this approach doesnt

		// problem: second iteration on the same depth used same slice, so it operator was added without "resetting"
		// perhaps temp structure can fix it
		copy(tempOperators, currOperators)
		tempOperators = append(currOperators, string(availableOperators[i]))
		operatorIdx = operatorIdx + i*int(math.Pow(2, float64(depth)))
		if depth > 0 {
			generateOperators(depth-1, operators, tempOperators, operatorIdx)
		} else {
			copy((*operators)[operatorIdx], tempOperators)
		}
	}
}
