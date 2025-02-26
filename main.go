package main

import (
	"aoc2024/puzzles/day1"
	"aoc2024/puzzles/day2"
	"aoc2024/puzzles/day3"
	"aoc2024/puzzles/day4"
	"aoc2024/puzzles/day5"
	"aoc2024/puzzles/day6"
	"aoc2024/puzzles/day7"
	"fmt"
	"os"
	"strconv"
)

var problemMap = map[int]func(){
	1: day1.Solve,
	2: day2.Solve,
	3: day3.Solve,
	4: day4.Solve,
	5: day5.Solve,
	6: day6.Solve,
	7: day7.Solve,
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <day number>")
		return
	}

	day, err := strconv.Atoi(os.Args[1])
	if err != nil || day < 1 || day > 25 {
		fmt.Println("Invalid day number. Please enter a number between 1 and 25.")
	}

	if fn, exists := problemMap[day]; exists {
		fn()
	} else {
		fmt.Println("Problem not implemented.")
	}
}
