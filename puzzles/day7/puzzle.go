package day7

import (
	"aoc2024/internal/utils"
	"fmt"
	"math"
	"strconv"
)

var valMap = map[int][]int{}
var availableOperators = []string{"+", "*", "||"}

func Solve(testFile string, params map[string]interface{}) {
	var inputFile = testFile
	inputContent, inputLength := utils.ReadFile(inputFile)

	valMap = splitInputToMap(utils.SplitStringByLines(inputContent))

	testVals := make([]int, 0, inputLength)
	for k := range valMap {
		testVals = append(testVals, k)
	}

	// For each input length (k), generate all possible operator permutations only once,
	// and store it in 2d slice (v) in map
	operatorPermutations := make(map[int][][]string)
	initOperatorPermutationMap(operatorPermutations, valMap)
	generateAllOperators(operatorPermutations)

	posResults := checkForResults(&testVals, operatorPermutations)
	totalSumOfValid := sumValues(&posResults, &testVals)
	fmt.Println(totalSumOfValid)
}

func sumValues(posResults *[]int, testVals *[]int) int {
	sum := 0
	for _, v := range *posResults {
		sum += (*testVals)[v]
	}

	return sum
}

func checkForResults(testVals *[]int, opPermMap map[int][][]string) []int {
	posResults := make([]int, 0, len(*testVals))
	for i, v := range *testVals {
		if checkForResult(v, opPermMap) {
			posResults = append(posResults, i)
		}
	}

	return posResults
}

func checkForResult(testVal int, opPermMap map[int][][]string) bool {
	operands := valMap[testVal]
	operators := opPermMap[len(operands)]
	for i, _ := range operators {
		tempResult := operands[0]
		for j, v := range operators[i] {
			tempResult = applyOperand(tempResult, operands[j+1], v)
			if tempResult > testVal {
				break
			}
		}
		if tempResult == testVal {
			return true
		}
	}

	return false
}

// TODO: store available operators in map or some struct, to not list them in this function explicitly
func applyOperand(op1 int, op2 int, operator string) int {
	if operator == "+" {
		return op1 + op2
	}
	if operator == "*" {
		return op1 * op2
	}
	if operator == "||" {
		tempResult, _ := strconv.Atoi(strconv.Itoa(op1) + strconv.Itoa(op2))
		return tempResult
	}
	return 0
}

func initOperatorPermutationMap(opMap map[int][][]string, valMap map[int][]int) {
	for _, v := range valMap {
		if opMap[len(v)] == nil {
			opMap[len(v)] = make([][]string, int(math.Pow(float64(len(availableOperators)), float64(len(v)-1))))
			for i, _ := range opMap[len(v)] {
				opMap[len(v)][i] = make([]string, len(v)-1)
			}
		}

	}
}

func generateAllOperators(operatorPermutations map[int][][]string) {
	for operandCount, v := range operatorPermutations {
		currOperators := make([]string, 0, operandCount-1)
		generateOperators(operandCount-2, &v, currOperators, 0)
	}
}

// Create map for number of operands in each line, so if there are multiple entries of 3/4/5 etc. operands.
//
// Generate operators list only once, before evaluating puzzle problem.
//
// Required parameters:
//
//	a) Depth: operators left, it would be either first to last or last to first,
//	   but it doesn't matter as long as order is preserved across all permutations
//	b) 2d slice: storage of all possible operator permutations
//	c) Operators string slice: to add them in recursive calls
//	d) Permutation Index: iteration/depth unique combination
//	   To add currOperators to operators 2d slice on specific index
func generateOperators(depth int, operators *[][]string, currOperators []string, permutationIdx int) {
	tempOperators := make([]string, 0, cap(currOperators))
	tempPermutationIdx := permutationIdx
	for opIdx, _ := range availableOperators {

		// problem: second iteration on the same recursion level
		// should use currOperators and permutationIdx values from one level up
		// fix: tempOperator/tempPermutationIdx slice to reset to original state
		copy(tempOperators, currOperators)
		tempOperators = append(currOperators, availableOperators[opIdx])
		tempPermutationIdx = permutationIdx + opIdx*int(math.Pow(float64(len(availableOperators)), float64(depth)))
		if depth > 0 {
			generateOperators(depth-1, operators, tempOperators, tempPermutationIdx)
		} else {
			copy((*operators)[tempPermutationIdx], tempOperators)
		}
	}
}
