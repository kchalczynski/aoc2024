package main

import (
	"aoc2024/puzzles/day1"
	"aoc2024/puzzles/day10"
	"aoc2024/puzzles/day11"
	"aoc2024/puzzles/day2"
	"aoc2024/puzzles/day3"
	"aoc2024/puzzles/day4"
	"aoc2024/puzzles/day5"
	"aoc2024/puzzles/day6"
	"aoc2024/puzzles/day7"
	"aoc2024/puzzles/day8"
	"aoc2024/puzzles/day9"
	"flag"
	"fmt"
	_ "net/http/pprof" // Register pprof handlers
	"os"
	"strconv"
)

var solvers = map[int]func(string, map[string]interface{}){
	1:  day1.Solve,
	2:  day2.Solve,
	3:  day3.Solve,
	4:  day4.Solve,
	5:  day5.Solve,
	6:  day6.Solve,
	7:  day7.Solve,
	8:  day8.Solve,
	9:  day9.Solve,
	10: day10.Solve,
	11: day11.Solve,
}

func main() {

	// Define flags for optional parameters
	testFileNumber := flag.String("test", "", "Specify which test file to use (e.g., --test 2 for test2.txt)")
	iterations := flag.Int("iterations", -1, "Specify number of iterations (used only in some puzzles)")
	output := flag.String("output", "", "Specify output file name (used only in some puzzles)")
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <day number>")
		return
	}

	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("Usage: go run . <day_number> [--test N] [--iterations X] [--output Y]")
		return
	}

	day, _ := strconv.Atoi(args[0])
	solveFunc, exists := solvers[day]
	if !exists {
		fmt.Printf("No puzzle found for Day %s\n", day)
		return
	}

	// If testFileNumber is empty, Solve() will handle the default internally
	testFile := ""
	defaultTestFile := "test1.txt"
	if *testFileNumber != "" {
		testFile = fmt.Sprintf("test%s.txt", *testFileNumber)
	} else {
		testFile = defaultTestFile
	}

	params := make(map[string]interface{})

	// Only pass "iterations" if explicitly provided
	if *iterations != -1 {
		params["iterations"] = *iterations
	}

	// Only pass "output" if explicitly provided
	if *output != "" {
		params["output"] = *output
	}

	// Call the Solve function with test file and optional parameters
	solveFunc(testFile, params)
}
