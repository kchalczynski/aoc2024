package day8

import (
	"aoc2024/internal/utils"
	"github.com/kr/pretty"
)

func Solve() {
	var inputFile = "puzzles/day8/test2.txt"

	inputContent, lineCount := utils.ReadFile(inputFile)

	inputArray := utils.ReadIntoMatrixByCharacter(inputContent, lineCount)
	//pretty.Println(inputArray)
	area := Board{rows: lineCount, cols: len(inputArray[0]), board: stringsToPositions(inputArray)}
	//pretty.Println(area)
	antennaPositionDict := initializeAntennaPosDict(area)
	pretty.Println(antennaPositionDict)

	// I could make map with list all combinations (2 of X) as value, and antenna symbol as key
	// Easier approach - make it all in one:
	// a) level 1 - iterate over each symbols position, generating combinations
	// b) level 2 - for each generated pair calculate distance (or even in place calculate rest)
	// c) level 2 - having distance and both Positions in pair get 2 antinode positions
	// by adding (to larger row/col)/subtracting (from smaller row/col) distance
	// and adding it right to map, or better to list and add to map later (might not matter, access by index anyway)
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
