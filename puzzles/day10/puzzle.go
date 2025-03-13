package day10

import (
	"aoc2024/internal/utils"
	"fmt"
	"strconv"
)

// https://adventofcode.com/2024/day/10

type Grid [][]*Position

var directions = []struct{ dy, dx int }{
	{0, 1},  // Up
	{0, -1}, // Down
	{1, 0},  // Right
	{-1, 0}, // Left
}

func Solve() {
	testNumber := 3
	var inputFile = fmt.Sprintf("puzzles/day10/test%d.txt", testNumber)
	inputContents, linesCount := utils.ReadFile(inputFile)

	var grid Grid = convertInputToPosition(utils.ReadIntoMatrixByCharacter(inputContents, linesCount))
	trailHeads, availableSteps := createMapOfSteps(grid)
	validTrails := searchValidTrails(trailHeads, availableSteps)
	fmt.Println(sumTrailScores(validTrails))

	distinctTrails := searchDistinctTrails(trailHeads, availableSteps)
	fmt.Println(sumTrailRatings(distinctTrails))
}

// Used to print out the input while solving the puzzle
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
func createMapOfSteps(grid [][]*Position) ([]*Position, map[*Position][]*Position) {
	rowSize, colSize := len(grid), len(grid[0])
	trailHeads := make([]*Position, 0)
	nextSteps := make(map[*Position][]*Position)
	for y, line := range grid {
		for x, currentPos := range line {

			if currentPos.height == 0 {
				trailHeads = append(trailHeads, currentPos)
			}
			availableSteps := make([]*Position, 0, 4)

			for _, dir := range directions {
				newY, newX := y+dir.dy, x+dir.dx

				if newY >= 0 && newY < rowSize && newX >= 0 && newX < colSize {
					neighbor := grid[newY][newX]

					if neighbor.height == currentPos.height+1 {
						availableSteps = append(availableSteps, neighbor)
					}
				}
			}

			nextSteps[currentPos] = availableSteps
		}
	}
	return trailHeads, nextSteps
}

// Depth First Search: consider every route from Start to Peak by `height` increment of 1 (positive).
//
//	Start: Position with `height` = 0
//	Peak: Position with `height` = 9
//
// For `Score` route is considered unique by every Start and Peak combination.
// Different positions on route from same combination of Start and Peak are not considered different route.
func dfs(start *Position, availableSteps map[*Position][]*Position, visited map[*Position]bool) []*Position {
	stack := []*Position{start}
	trails := make([]*Position, 0)
	for len(stack) > 0 {
		currentPos := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if visited[currentPos] {
			continue
		}

		visited[currentPos] = true
		// I don't have to reset visited while going back (checking another neighbor),
		// because checking already visited Position would not result in different End
		if currentPos.height == 9 {
			trails = append(trails, currentPos)
			continue
		}

		for _, neighbor := range availableSteps[currentPos] {
			if !visited[neighbor] && neighbor.height == currentPos.height+1 {
				stack = append(stack, neighbor)
			}
		}
	}
	return trails
}

// Same as dfs, but for `Rating` route is considered unique by `distinct` sequence of steps from Start to Peak,
// So there might be multiple distinct routes for the same Start and Peak combination
func dfsDistinct(start *Position, availableSteps map[*Position][]*Position) [][10]*Position {
	stack := []*Position{start}
	distinctTrails := make([][10]*Position, 0)
	trail := [10]*Position{}
	for len(stack) > 0 {
		currentPos := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		trail[currentPos.height] = currentPos

		if currentPos.height == 9 {
			var distinctTrail = trail
			distinctTrails = append(distinctTrails, distinctTrail)
			continue
		}

		for _, neighbor := range availableSteps[currentPos] {
			if neighbor.height == currentPos.height+1 {
				stack = append(stack, neighbor)
			}
		}
	}
	return distinctTrails
}

// Find all valid trails for each of the trail heads
//
// Valid Trail is considered by unique combination of Start and Peak
func searchValidTrails(trailHeads []*Position, availableSteps map[*Position][]*Position) map[*Position][]*Position {
	visited := make(map[*Position]bool)
	validTrails := make(map[*Position][]*Position)

	for _, trailHead := range trailHeads {
		visited = make(map[*Position]bool)

		reachablePeaks := dfs(trailHead, availableSteps, visited)
		if len(reachablePeaks) > 0 {
			validTrails[trailHead] = reachablePeaks
		}
	}

	return validTrails
}

// Find all distinct valid trails for each of the trail heads
//
// Valid trail is considered by unique sequence of steps from Start to Peak
func searchDistinctTrails(trailHeads []*Position, availableSteps map[*Position][]*Position) map[*Position][][10]*Position {
	distinctTrails := make(map[*Position][][10]*Position)

	for _, trailHead := range trailHeads {
		distinctTrailsByHead := dfsDistinct(trailHead, availableSteps)
		if len(distinctTrailsByHead) > 0 {
			distinctTrails[trailHead] = distinctTrailsByHead
		}
	}

	return distinctTrails
}

func sumTrailScores(validTrails map[*Position][]*Position) int {
	trailScoreSum := 0
	for _, validTrail := range validTrails {
		trailScore := len(validTrail)
		trailScoreSum += trailScore
	}
	return trailScoreSum
}

func sumTrailRatings(distinctTrails map[*Position][][10]*Position) int {
	trailRatingSum := 0
	for _, validTrail := range distinctTrails {
		trailRating := len(validTrail)
		trailRatingSum += trailRating
	}
	return trailRatingSum
}
