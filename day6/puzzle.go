package main

import (
	"errors"
	"fmt"
)

/*
https://adventofcode.com/2024/day/6

Part 2
Copy map for each try
Obstructions can only be placed on visitedMarks, except the first one (starting point)
If guards path AFTER new obstruction contains both visited position and position with same direction
then it should lead to loop in path, as no new obstruction can be placed and nothing can really change his way

Task list:
 1. Get list of visited positions;
 2. Copy map3
 3. May need a starting position, place obstruction at the next position of the guard
    Q: Can I place obstruction at the second position of the guard, just after starting?
 4. Simulate path and compare
*/

var BaseRoom Area

func main() {

	var inputFile = "input.txt"
	var inputContent string
	var inputLinesCount int

	inputContent, inputLinesCount = readFile(inputFile)
	initializeArea(readContentIntoMatrix(inputContent, inputLinesCount))
	guard, err := initializeGuard(BaseRoom.board)

	// Global variable for part 1, copies for part 2
	room := BaseRoom.CopyRoom()
	if err != nil {
		fmt.Println(err)
	} else {
		room.markVisited(guard.posY, guard.posX)
		guard.Patrol(room)
		fmt.Println(guard.posX, guard.posY, guard.visitedCount)
	}

	loopCount := guard.countPossibleLoops()
	fmt.Println(loopCount)

}

func initializeArea(board [][]string) {
	BaseRoom.sizeX = len(board[0])
	BaseRoom.sizeY = len(board)
	BaseRoom.board = stringsToPositions(board)
}

func initializeGuard(areaMap [][]Position) (Guard, error) {
	visitedList := make([]Position, 0)

	for i, row := range areaMap {
		for j, val := range row {
			if val.direction == dirN || val.direction == dirE || val.direction == dirS || val.direction == dirW {
				areaMap[i][j].isVisited = true
				visitedList = append(visitedList, areaMap[i][j])
				return Guard{i, j, val.direction, 1, visitedList}, nil
			}
		}
	}
	return Guard{}, errors.New("No guard on the map")
}
