package day11

import (
	"aoc2024/internal/utils"
	"fmt"
	"math"
	"strconv"
)

// https://adventofcode.com/2024/day/11
var multiplier int = 2024

func Solve() {
	testNumber := 3
	turns := 60
	var inputFile = fmt.Sprintf("puzzles/day11/test%d.txt", testNumber)
	inputContents, _ := utils.ReadFile(inputFile)

	stoneSeq := utils.ReadAsNumberSeq(inputContents)
	fmt.Println(stoneSeq)

	//approximateMaxStoneValue(stoneSeq, multiplier, turns)

	stoneSeq = blinkTimes(turns, stoneSeq)
	fmt.Println(len(stoneSeq))
}

// Approximate max size of stones.
// Result = doesn't fit in uint32
func approximateMaxStoneValue(stoneSeq []int, turns int) {
	for _, val := range stoneSeq {
		elementSize := approximateElementSize(val, turns)
		if elementSize > float64(^uint32(0)) {
			fmt.Printf("Element size after %d turns exceeds uint32 limit: %f\n", turns, elementSize)
		} else {
			fmt.Printf("Element size after %d turns: %f (fits in uint32)\n", turns, elementSize)
		}
	}
}

func approximateElementSize(initialValue int, turns int) float64 {
	return float64(initialValue) * math.Pow(float64(multiplier), float64(turns))
}

func blinkTimes(turns int, stoneSeq []int) []int {
	var bufferSize = int(math.Pow(2, float64(turns/2)))
	var batchSize = 1000000
	buffer1 := make([]int, 0, bufferSize)
	buffer2 := make([]int, 0, bufferSize)

	current := &buffer1
	next := &buffer2

	*current = append(*current, stoneSeq...)
	for i := 0; i < turns; i++ {
		fmt.Println("Turn: ", i, " | Seq size: ", len(*current))
		fmt.Println("Batches: ", 1+len(*current)/batchSize)
		for i := 0; i < len(*current); i += batchSize {
			end := min(i+batchSize, len(*current))
			*next = (*next)[:0]
			newQ := (*current)[i:end]
			*current, *next = blink(&newQ, (*next)[i:end])
			*current, *next = blink((*current)[i:end], (*next)[i:end])
		}
	}
	return *current
}

func blink(currStoneSeq *[]int, nextStoneSeq *[]int) ([]int, []int) {
	for i := 0; i < len(*currStoneSeq); i++ {
		if (*currStoneSeq)[i] == 0 {
			*nextStoneSeq = append(*nextStoneSeq, 1)
			continue
		}
		strStone := strconv.Itoa((*currStoneSeq)[i])
		if len(strStone)%2 == 0 {
			newStones := splitStone((*currStoneSeq)[i], len(strStone))
			*nextStoneSeq = append(*nextStoneSeq, newStones...)

		} else {
			newValue := (*currStoneSeq)[i] * multiplier
			/*if !fitsUint16(newValue) {
				if !fitsUint32(newValue) {
					fmt.Println("Value bigger than 32bit: ", newValue)
				}
				fmt.Println("Value bigger than 16bit: ", newValue)
			}*/
			*nextStoneSeq = append(*nextStoneSeq, newValue)
		}
	}
	return *nextStoneSeq, *currStoneSeq
}

func splitStone(stone int, stoneLength int) []int {
	// to string, split in half, atoi to int, will get rid of leading 0's
	stoneStr := strconv.Itoa(stone)
	stone1, _ := strconv.Atoi(stoneStr[:stoneLength/2])
	stone2, _ := strconv.Atoi(stoneStr[stoneLength/2:])

	stones := []int{stone1, stone2}
	return stones
}

// check how big stones get
func fitsUint16(n int) bool {
	return n >= 0 && n <= math.MaxUint16
}
func fitsUint32(n int) bool {
	return n >= 0 && n <= math.MaxUint32
}
