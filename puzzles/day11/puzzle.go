package day11

import (
	"aoc2024/internal/utils"
	"fmt"
	"strconv"
)

// https://adventofcode.com/2024/day/11

func Solve() {
	testNumber := 2
	turns := 60
	var inputFile = fmt.Sprintf("puzzles/day11/test%d.txt", testNumber)
	inputContents, _ := utils.ReadFile(inputFile)

	stoneSeq := utils.ReadAsNumberSeq(inputContents)
	fmt.Println(stoneSeq)
	stoneSeq = blinkTimes(turns, stoneSeq)
	fmt.Println(len(stoneSeq))
}

func blinkTimes(turns int, stoneSeq []int) []int {
	currStoneSeq := stoneSeq
	nextStoneSeq := make([]int, 0, 2*len(stoneSeq))
	for i := 0; i < turns; i++ {
		nextStoneSeq = nextStoneSeq[:0]
		stoneSeq = blink(&currStoneSeq, &nextStoneSeq)
		//if len(stoneSeq) > 10 {
		//	midIdx := len(stoneSeq) / 2
		//	newSeq := make([]int, 0)
		//	resultPart1 := blink(stoneSeq[0:midIdx])
		//	newSeq = append(newSeq, resultPart1...)
		//	resultPart2 := blink(stoneSeq[midIdx:])
		//	newSeq = append(newSeq, resultPart2...)
		//	stoneSeq = newSeq
		//} else {
		//	stoneSeq = blink(stoneSeq)
		//}
	}
	return stoneSeq
}

func blink(stoneSeq *[]int, newStoneSeq *[]int) []int {
	newStoneSeq := make([]int, 0, 2*len(stoneSeq))
	for i := 0; i < len(stoneSeq); i++ {
		if stoneSeq[i] == 0 {
			newStoneSeq = append(newStoneSeq, 1)
		} else if len(strconv.Itoa(stoneSeq[i]))%2 == 0 {
			newStones := splitStone(stoneSeq[i])
			newStoneSeq = append(newStoneSeq, newStones...)

		} else {
			newStoneSeq = append(newStoneSeq, stoneSeq[i]*2024)
		}
	}
	return newStoneSeq
}

func splitStone(stone int) []int {
	// to string, split in half, atoi to int, will get rid of leading 0's
	stoneStr := strconv.Itoa(stone)
	stone1, _ := strconv.Atoi(stoneStr[:len(stoneStr)/2])
	stone2, _ := strconv.Atoi(stoneStr[len(stoneStr)/2:])

	stones := []int{stone1, stone2}
	return stones
}
