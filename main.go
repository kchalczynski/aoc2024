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

var dayNumber int
var testFileNumber int
var iterations int
var output string

func init() {
	// Define flags for optional parameters
	flag.IntVar(&dayNumber, "day", -1, "Specify which day puzzle to solve (e.g., -d 2 or --day 2 for day2)")
	flag.IntVar(&dayNumber, "d", -1, "Alias for --day")
	flag.IntVar(&testFileNumber, "test", -1, "Specify which test file to use (e.g., --test 2 for test2.txt)")
	flag.IntVar(&testFileNumber, "t", -1, "Alias for --test")
	flag.IntVar(&iterations, "iterations", -1, "Specify number of iterations (used only in some puzzles)")
	flag.IntVar(&iterations, "i", -1, "Alias for --iterations")
	flag.StringVar(&output, "output", "", "Specify output file name (used only in some puzzles)")
	flag.StringVar(&output, "o", "", "Alias for --output")

}

func main() {

	fmt.Println("Raw Args:", os.Args)

	//w Workaround to parse flags after first arg
	flag.Parse()

	// Print parsed flag values (for debugging)
	fmt.Println("Parsed Flags:")
	fmt.Println("  -t/--test:", testFileNumber)
	fmt.Println("  -d/--day:", dayNumber)
	fmt.Println("  -i/--iterations:", iterations)
	fmt.Println("  -o/--output:", output)

	args := flag.CommandLine.Args()
	fmt.Println(args)

	var day int
	if len(args) < 1 {
		if dayNumber == -1 {
			fmt.Println("Usage: go run . [-d N]/[--day N] [--test N] [--iterations X] [--output Y] day_number" +
				" [optional, can be specified via flag -d/--day or argument]")
			fmt.Println("Day number must be specified either via \"-d\" or \"--day\", or explicitely (after flags)")
			return
		} else {
			day = dayNumber
		}
	} else {
		dayArg, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("Invalid day number %s\n", args[0])
			return
		} else {
			day = dayArg
		}
	}

	// Print day number (for debugging)
	fmt.Println("Day argument:", day)

	solveFunc, exists := solvers[day]
	if !exists {
		fmt.Printf("No puzzle found for Day %s\n", day)
		return
	}

	// If testFileNumber is empty, Solve() will handle the default internally
	testFile := fmt.Sprintf("puzzles/day%d/test1.txt", day)

	if testFileNumber != -1 {
		testFile = fmt.Sprintf("puzzles/day%d/test%d.txt", day, testFileNumber)
	}

	// Print final test file selection (for debugging)
	fmt.Println("Test file selected:", testFile)

	params := make(map[string]interface{})

	// Only pass "iterations" if explicitly provided
	if iterations != -1 {
		params["iterations"] = iterations
	}

	// Only pass "output" if explicitly provided
	if output != "" {
		params["output"] = output
	}

	// Call the Solve function with test file and optional parameters
	solveFunc(testFile, params)
}
