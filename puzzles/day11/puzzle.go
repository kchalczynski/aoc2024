package day11

import (
	"aoc2024/internal/utils"
	"fmt"
	"math"
	"strconv"
)

// https://adventofcode.com/2024/day/11
// Stones = numbers
// Turns = blinks

// Default iterations value
var turns = 25
var multiplier = 2024

var cache = make(map[string]int)

var cacheRetrieveCount = 0
var totalRecursiveCallCount = 0
var cachedPerTurn = make(map[int]int)
var recursiveCallsPerTurn = make(map[int]int)

func initBenchmarkMap(turns int, benchmarkMap map[int]int) {
	for i := 0; i < turns; i++ {
		benchmarkMap[i] = 0
	}
}

// Solve function extracts "iterations" as "turns" if available
func Solve(testFile string, params map[string]interface{}) {

	// Override if provided
	if val, ok := params["iterations"].(int); ok {
		turns = val
	}

	var inputFile = testFile
	inputContents, _ := utils.ReadFile(inputFile)

	stoneSeq := utils.ReadAsNumberSeq(inputContents)
	fmt.Println(stoneSeq)

	// For Part 1 we can generate whole sequence of stones
	if turns < 10 {
		stoneSeq = blinkTimes(turns, stoneSeq)
		fmt.Println(len(stoneSeq))
	} else {
		initBenchmarkMap(turns, cachedPerTurn)
		initBenchmarkMap(turns, recursiveCallsPerTurn)
		totalElements := countTotalElements(stoneSeq, turns)
		fmt.Println("Total elements after", turns, "turns:", totalElements)
		for i := 0; i < turns; i++ {
			fmt.Println(fmt.Sprintf("Turn %d: Cached %d/%d total cache calls;"+
				" %d/%d recursive calls out of total.",
				i, cachedPerTurn[i], cacheRetrieveCount, recursiveCallsPerTurn[i], totalRecursiveCallCount))

		}
	}

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

// After approximately ~45 turns/iterations we run out of memory.
// Full array of generated stones takes ~6GB of RAM at this point.
// Storing it on disk is not viable as it grows further,
// but for Part2 we only need total count of stones after N turns
func blinkTimes(turns int, stoneSeq []int) []int {

	//	Tried to save memory by splitting input into chunks, working with two
	//	slices only, processing it sequentially. Works fine for processing each part,
	//	but resulting sequence becomes too big by ~turn 45 to store it in memory.
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

// currStoneSeq = current batch
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

// Recursively process a single stone and count total number of stones generated
func dfsCount(stone int, turnsLeft int) int {

	currentTurn := turns - turnsLeft
	//	Last turn, count element as 1
	if turnsLeft == 0 {
		return 1
	}

	//	Check if current stone value was previously computed
	key := fmt.Sprintf("%d_%d", stone, turnsLeft)
	if result, found := cache[key]; found {
		cacheRetrieveCount += 1
		cachedPerTurn[currentTurn] += 1
		return result
	} else {
		totalRecursiveCallCount += 1
		recursiveCallsPerTurn[currentTurn] += 1
	}

	var result int
	if stone == 0 {
		result = dfsCount(1, turnsLeft-1)
	} else {
		strStone := strconv.Itoa(stone)
		if len(strStone)%2 == 0 {
			newStones := splitStone(stone, len(strStone))
			result = dfsCount(newStones[0], turnsLeft-1) + dfsCount(newStones[1], turnsLeft-1)
		} else {
			result = dfsCount(stone*2024, turnsLeft-1)
		}
	}
	cache[key] = result

	return result
}
