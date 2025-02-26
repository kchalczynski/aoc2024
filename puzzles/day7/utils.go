package day7

import (
	"aoc2024/internal/utils"
	"strconv"
	"strings"
)

func splitInputToMap(input []string) map[int][]int {
	// format: "test_val: val1 val2 val3 ..."
	// test_val is a candidate for sum/product of val args
	// key: test_val, value: slice of val args

	valMap := make(map[int][]int)

	for _, line := range input {
		vals := strings.Split(line, ": ")
		valArgs := strings.Split(
			strings.TrimSpace(
				strings.ReplaceAll(vals[1], " ", ";")),
			";")
		key, _ := strconv.Atoi(vals[0])
		values := utils.StringsToInts(valArgs)
		valMap[key] = append(valMap[key], values...)
	}

	return valMap
}
