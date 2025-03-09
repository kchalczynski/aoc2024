package day8

import (
	"aoc2024/internal/utils"
	"fmt"
	"maps"
	"math"
)

var isPart2 = false

func Solve( /*testNumber int*/ ) {
	testNumber := 2
	if testNumber == 0 {
		testNumber = 1
	}
	var inputFile = fmt.Sprintf("puzzles/day8/test%d.txt", testNumber)

	inputContent, lineCount := utils.ReadFile(inputFile)

	inputArray := utils.ReadIntoMatrixByCharacter(inputContent, lineCount)
	area := Board{rows: lineCount, cols: len(inputArray[0]), board: stringsToPositions(inputArray)}
	antennaPositionDict := initializeAntennaPosDict(area)

	// I could make map with list all combinations (2 of X) as value, and antenna symbol as key
	// (number of combinations r of n: n!/(r!*(n-r)!)

	// Easier approach - make it all in one:
	// a) level 1 - iterate over each symbols position, generating combinations
	// b) level 2 - for each generated pair calculate distance (or even in place calculate rest)
	// c) level 2 - having distance and both Positions in pair get 2 antinode positions
	// by adding (to larger row/col)/subtracting (from smaller row/col) distance
	// and adding it right to map, or better to list and add to map later (might not matter, access by index anyway)

	combinationCountMap := make(map[string]int)
	combinationMap := make(map[string][]PositionPair)
	for k := range maps.Keys(antennaPositionDict) {
		combinationCount := getCombinationNumber(len(antennaPositionDict[k]), 2)
		combinationCountMap[k] = combinationCount
		combinationMap[k] = generateCombinations(antennaPositionDict[k], combinationCount)

	}
	generateAntinodes(combinationMap, &area)
	fmt.Println(countAntinodes(&area))
	isPart2 = true
	generateAntinodes(combinationMap, &area)
	fmt.Println(countAntinodes(&area))
}

func initializeAntennaPosDict(area Board) map[string][]SimplePosition {
	antennaPositionDict := make(map[string][]SimplePosition)
	for row := 0; row < area.rows; row++ {
		for col, pos := range area.board[row] {
			if pos.isAntenna {
				antennaPositionDict[pos.antennaType] = append(antennaPositionDict[pos.antennaType], SimplePosition{
					row: row,
					col: col,
				})
			}
		}
	}
	return antennaPositionDict
}

// only for combinations of length 2
func generateCombinations(positions []SimplePosition, n int) []PositionPair {
	pairs := make([]PositionPair, 0, n)
	for i := 0; i < len(positions)-1; i++ {
		for j := i + 1; j < len(positions); j++ {
			pairs = append(pairs, PositionPair{positions[i], positions[j]})
		}

	}
	return pairs
}

func generateAntinodes(antennaPairs map[string][]PositionPair, area *Board) {
	for _, v := range antennaPairs {
		setAntinodesOnBoard(createAntinodes(v, area.rows, area.cols), area)
	}
}

func createAntinodes(antennaPairs []PositionPair, rowSize, colSize int) []SimplePosition {
	antinodes := make([]SimplePosition, 0, len(antennaPairs)*2)
	for _, pair := range antennaPairs {
		rowDiff, colDiff := 0, 0
		rowDiff = int(math.Abs(float64(pair.pos1.row - pair.pos2.row)))
		colDiff = int(math.Abs(float64(pair.pos1.col - pair.pos2.col)))
		// should set global variable/constant in Solve function, but it would require resetting map state
		// or counting antinodes after placing them on board
		if !isPart2 {
			antinode1 := SimplePosition{}
			antinode2 := SimplePosition{}

			if pair.pos1.row > pair.pos2.row {
				antinode1.row = pair.pos1.row + rowDiff
				antinode2.row = pair.pos2.row - rowDiff
			} else {
				antinode1.row = pair.pos1.row - rowDiff
				antinode2.row = pair.pos2.row + rowDiff
			}
			if pair.pos1.col > pair.pos2.col {
				antinode1.col = pair.pos1.col + colDiff
				antinode2.col = pair.pos2.col - colDiff
			} else {
				antinode1.col = pair.pos1.col - colDiff
				antinode2.col = pair.pos2.col + colDiff
			}
			addNodeToList(&antinodes, antinode1, rowSize, colSize)
			addNodeToList(&antinodes, antinode2, rowSize, colSize)

		} else {
			//fmt.Printf("part2, rowDif, colDif: %d,%d\n", rowDiff, colDiff)
			curRow, curCol := pair.pos1.row, pair.pos1.col
			antinodes = append(antinodes, SimplePosition{
				row: curRow,
				col: curCol,
			})
			for curRow >= 0 && curCol >= 0 && curRow < rowSize && curCol < colSize {
				//fmt.Printf("Cur row: %d, Cur col: %d, Max row: %d, Max col: %d", curRow, curCol, rowSize, colSize)
				if pair.pos1.row < pair.pos2.row {
					curRow = curRow + rowDiff
				} else {
					curRow = curRow - rowDiff
				}
				if pair.pos1.col < pair.pos2.col {
					curCol = curCol + colDiff
				} else {
					curCol = curCol - colDiff
				}
				addNodeToList(&antinodes, SimplePosition{row: curRow, col: curCol}, rowSize, colSize)
			}
			// reset row and col to go other direction
			curRow, curCol = pair.pos1.row, pair.pos1.col
			for curRow >= 0 && curCol >= 0 && curRow < rowSize && curCol < colSize {

				if pair.pos1.row < pair.pos2.row {
					curRow = curRow - rowDiff
				} else {
					curRow = curRow + rowDiff
				}
				if pair.pos1.col < pair.pos2.col {
					curCol = curCol - colDiff
				} else {
					curCol = curCol + colDiff
				}
				addNodeToList(&antinodes, SimplePosition{row: curRow, col: curCol}, rowSize, colSize)
			}

		}
	}
	return antinodes
}

func addNodeToList(nodes *[]SimplePosition, node SimplePosition, rows, cols int) {
	if node.row >= 0 &&
		node.col >= 0 &&
		node.row < rows &&
		node.col < cols {
		*nodes = append(*nodes, node)
	}
}
func setAntinodesOnBoard(antinodes []SimplePosition, area *Board) {
	for _, node := range antinodes {
		if node.row >= 0 &&
			node.col >= 0 &&
			node.row < area.rows &&
			node.col < area.cols &&
			!area.board[node.row][node.col].isAntinode {
			area.board[node.row][node.col].isAntinode = true
		}
	}
}

func countAntinodes(area *Board) int {
	count := 0
	for _, row := range area.board {
		for _, pos := range row {
			if pos.isAntinode {
				count += 1
			}
		}
	}
	return count
}
