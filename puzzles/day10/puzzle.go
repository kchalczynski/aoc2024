package day10

import (
	"aoc2024/internal/utils"
	"fmt"
	"strconv"
)

func Solve() {
	testNumber := 1
	var inputFile = fmt.Sprintf("puzzles/day9/test%d.txt", testNumber)
	inputContents, linesCount := utils.ReadFile(inputFile)

	inputArray := convertInputToInt(utils.ReadIntoMatrixByCharacter(inputContents, linesCount))
	trailHeads, nextSteps := createMapOfSteps(inputArray)

	trailHeadScores := make(map[Position]int)
	for _, pos := range trailHeads {
		trailHeadScores[pos] = 0
	}
}

func convertInputToInt(input [][]string) [][]int {
	intArray := make([][]int, len(input))
	for i := 0; i < len(input); i++ {
		intArray[i] = make([]int, len(input[i]))
		for j := 0; j < len(input[i]); j++ {
			intArray[i][j], _ = strconv.Atoi(input[i][j])
		}
	}
	return intArray
}

func convertInputToPosition(input [][]string) [][]*Position {
	positionArray := make([][]*Position, len(input))
	for i := 0; i < len(input); i++ {
		positionArray[i] = make([]*Position, len(input[i]))
		for j := 0; j < len(input[i]); j++ {
			height, _ := strconv.Atoi(input[i][j])
			pos := new(Position)
			pos.y, pos.x, pos.height = i, j, height
			positionArray[i][j] = pos
		}
	}
	return positionArray
}

// Create a map of viable next steps for every position
func createMapOfSteps(input [][]*Position) ([]*Position, map[*Position][]*Position) {
	rowSize, colSize := len(input), len(input[0])
	trailHeads := make([]*Position, 0)
	nextSteps := make(map[*Position][]*Position)
	for y, line := range input {
		for x, height := range line {
			currentPosition := Position{y: y, x: x, value: height}
			if height == 0 {
				trailHeads = append(trailHeads, currentPosition)
			}
			availableSteps := make([]Position, 0, 4)
			if y > 0 && input[y-1][x] == height+1 {
				// I don't need height in available steps list for now, as every step will be height+1
				// and don't need references for position objects? only coordinates
				availableSteps = append(availableSteps, Position{y: y - 1, x: x})
			}
			if y < rowSize-1 && input[y+1][x] == height+1 {
				availableSteps = append(availableSteps, Position{y: y + 1, x: x})
			}
			if x > 0 && input[y][x-1] == height+1 {
				availableSteps = append(availableSteps, Position{y: y, x: x - 1})
			}
			if x < colSize-1 && input[y][x+1] == height+1 {
				availableSteps = append(availableSteps, Position{y: y, x: x + 1})
			}

			nextSteps[currentPosition] = append(nextSteps[currentPosition], availableSteps...)
		}
	}
	return trailHeads, nextSteps
}
