package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

/*
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
const (
	visitedMark     = "X"
	notVisitedMark  = "."
	obstructionMark = "#"
	dirN            = "^"
	dirE            = ">"
	dirS            = "v"
	dirW            = "<"
)

var Room Area

func initializeArea(board [][]string) {
	Room.sizeX = len(board[0])
	Room.sizeY = len(board)
	Room.board = stringsToPositions(board)
}

func stringsToPositions(board [][]string) [][]Position {
	colN, rowN := len(board), len(board[0])
	positions := make([][]Position, colN)
	for i := 0; i < colN; i++ {
		positions[i] = make([]Position, rowN)
		for j := 0; j < rowN; j++ {
			positions[i][j] = Position{
				row:           i,
				col:           j,
				isVisited:     false,
				isObstruction: board[i][j] == obstructionMark,
				direction: func(cell string) string {
					if cell == dirN || cell == dirE || cell == dirS || cell == dirW {
						return cell
					}
					return ""
				}(board[i][j]),
			}
		}
	}
	return positions
}
func markVisited(y, x int) {
	Room.board[y][x].isVisited = true
}

func main() {

	var inputFile = "input.txt"
	var inputContent string
	var inputLinesCount int

	inputContent, inputLinesCount = readFile(inputFile)
	initializeArea(readContentIntoMatrix(inputContent, inputLinesCount))

	guard, err := initializeGuard(Room.board)
	if err != nil {
		fmt.Println(err)
	} else {
		markVisited(guard.posY, guard.posX)
		guard.Patrol()
		println(guard.posX, guard.posY, guard.visitedCount)

	}

}

type Area struct {
	sizeX, sizeY int
	board        [][]Position
}

type Position struct {
	row, col                 int
	isVisited, isObstruction bool
	direction                string
}

type Guard struct {
	posY, posX   int
	direction    string
	visitedCount int
	visitedList  []Position
}

func (g *Guard) Patrol() {
	for {
		err := g.Move()
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}

func (g *Guard) TurnR() {
	switch g.direction {
	case dirN:
		g.direction = dirE
	case dirE:
		g.direction = dirS
	case dirS:
		g.direction = dirW
	case dirW:
		g.direction = dirN
	}
}

func (g *Guard) Move() error {

	if err := checkRoomBounds(g.posY, g.posX, g.direction); err != nil {
		return err
	}

	if !g.CheckAhead() {
		g.TurnR()
	}

	switch g.direction {
	case dirN:
		g.posY--
	case dirE:
		g.posX++
	case dirS:
		g.posY++
	case dirW:
		g.posX--
	}

	if !Room.board[g.posY][g.posX].isVisited {
		markVisited(g.posY, g.posX)
		g.visitedCount++
	}

	return nil
}

func (g *Guard) CheckAhead() bool {

	switch g.direction {
	case dirN:
		if !Room.board[g.posY-1][g.posX].isObstruction {
			return true
		}
	case dirE:
		if !Room.board[g.posY][g.posX+1].isObstruction {
			return true
		}
	case dirS:
		if !Room.board[g.posY+1][g.posX].isObstruction {
			return true
		}
	case dirW:
		if !Room.board[g.posY][g.posX-1].isObstruction {
			return true
		}
	}
	return false
}

func checkRoomBounds(row, col int, direction string) error {
	newRow, newCol := row, col
	switch direction {
	case dirN:
		newRow = row - 1
	case dirE:
		newCol = col + 1
	case dirS:
		newRow = row + 1
	case dirW:
		newCol = col - 1
	}
	if newRow < 0 || newCol < 0 || newRow >= Room.sizeY || newCol >= Room.sizeX {
		return errors.New("Out of room bounds")
	}
	return nil
}

func readContentIntoMatrix(input string, linesCount int) [][]string {

	letterMatrix := make([][]string, linesCount)

	for i, val := range strings.Split(input, "\n") {
		letterMatrix[i] = strings.Split(val, "")
	}
	return letterMatrix
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

func readFile(fileName string) (string, int) {
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error opening file: ", err)
		return "", 0
	}
	var content = ""
	var lineCount = 0

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineCount += 1
		content += line + "\n"
	}

	return strings.TrimSuffix(content, "\n"), lineCount
}
