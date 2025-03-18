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
	turns := 43
	var inputFile = fmt.Sprintf("puzzles/day11/test%d.txt", testNumber)
	inputContents, _ := utils.ReadFile(inputFile)

	stoneSeq := utils.ReadAsNumberSeq(inputContents)
	fmt.Println(stoneSeq)

	//approximateMaxStoneValue(stoneSeq, multiplier, turns)

	if turns < 35 {
		stoneSeq = blinkTimes(turns, stoneSeq)
		fmt.Println(len(stoneSeq))
	}

	totalElements := countTotalElements(stoneSeq, turns)
	fmt.Println("Total elements after", turns, "turns:", totalElements)
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

// After approximately ~45 blinks/iterations we run out of memory.
// Full array of generated stones takes ~6GB of RAM
func blinkTimes(turns int, stoneSeq []int) []int {
	var bufferSize = int(math.Pow(2, float64(turns/10))) * len(stoneSeq)
	var batchSize = 100000
	current := make([]int, 0, bufferSize)
	next := make([]int, 0, bufferSize)

	current = append(current, stoneSeq...)
	for i := 0; i < turns; i++ {
		fmt.Println("Turn: ", i, " | Seq size: ", len(current))
		fmt.Println("Batches: ", 1+len(current)/batchSize)
		next = (next)[:0]

		for i := 0; i < len(current); i += batchSize {
			end := min(i+batchSize, len(current))
			blink(current[i:end], &next)
		}
		current, next = next, current
	}
	return current
}

// currStoneSeq may be just current batch
func blink(currStoneSeq []int, nextStoneSeq *[]int) {
	for i := 0; i < len(currStoneSeq); i++ {
		if (currStoneSeq)[i] == 0 {
			*nextStoneSeq = append(*nextStoneSeq, 1)
			continue
		}
		strStone := strconv.Itoa((currStoneSeq)[i])
		if len(strStone)%2 == 0 {
			newStones := splitStone((currStoneSeq)[i], len(strStone))
			*nextStoneSeq = append(*nextStoneSeq, newStones...)

		} else {
			newValue := (currStoneSeq)[i] * multiplier
			*nextStoneSeq = append(*nextStoneSeq, newValue)
		}
	}
}

func splitStone(stone int, stoneLength int) []int {
	// to string, split in half, atoi to int, will get rid of leading 0's
	stoneStr := strconv.Itoa(stone)
	stone1, _ := strconv.Atoi(stoneStr[:stoneLength/2])
	stone2, _ := strconv.Atoi(stoneStr[stoneLength/2:])

	stones := []int{stone1, stone2}
	return stones
}

func countTotalElements(numbers []int, turnsTotal int) int {
	total := 0
	for _, number := range numbers {
		total += dfsCount(number, turnsTotal)
	}
	return total
}

// dfsCount recursively processes a single number and counts its total contributions
func dfsCount(stone int, turnsLeft int) int {
	if turnsLeft == 0 {
		return 1 // At the end, it counts as a single element
	}
	if stone == 0 {
		// 0 turns into 1 in one step
		return dfsCount(1, turnsLeft-1)
	}

	strStone := strconv.Itoa(stone)
	if len(strStone)%2 == 0 {
		// Even length: Split into two
		newStones := splitStone(stone, len(strStone))
		return dfsCount(newStones[0], turnsLeft-1) + dfsCount(newStones[1], turnsLeft-1)
	} else {
		// Odd length: Multiply
		return dfsCount(stone*2024, turnsLeft-1)
	}
}

// check how big stones get
func fitsUint16(n int) bool {
	return n >= 0 && n <= math.MaxUint16
}
func fitsUint32(n int) bool {
	return n >= 0 && n <= math.MaxUint32
}
