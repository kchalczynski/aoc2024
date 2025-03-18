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
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof" // Register pprof handlers
	"os"
	"strconv"
)

var problemMap = map[int]func(){
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
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <day number>")
		return
	}

	day, err := strconv.Atoi(os.Args[1])
	if err != nil || day < 1 || day > 25 {
		fmt.Println("Invalid day number. Please enter a number between 1 and 25.")
	}

	// TODO: number of test input to use as second argument
	/*	test, err := strconv.Atoi(os.Args[2])
		if err != nil || test < 1 || day > 25 {
			fmt.Println("Invalid test number. Please enter a number between 1 and X.")
		}*/

	if fn, exists := problemMap[day]; exists {
		go func() {
			log.Println(http.ListenAndServe("localhost:6060", nil))
		}()

		fn()
	} else {
		fmt.Println("Problem not implemented.")
	}
}
