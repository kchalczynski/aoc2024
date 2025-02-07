package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

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
					switch cell {
					case dirN:
						return dirN
					case dirE:
						return dirE
					case dirS:
						return dirS
					case dirW:
						return dirW
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
	case "N":
		g.direction = "E"
	case "E":
		g.direction = "S"
	case "S":
		g.direction = "W"
	case "W":
		g.direction = "N"
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
	case "N":
		g.posY--
	case "E":
		g.posX++
	case "S":
		g.posY++
	case "W":
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
	case "N":
		if !Room.board[g.posY-1][g.posX].isObstruction {
			return true
		}
	case "E":
		if !Room.board[g.posY][g.posX+1].isObstruction {
			return true
		}
	case "S":
		if !Room.board[g.posY+1][g.posX].isObstruction {
			return true
		}
	case "W":
		if !Room.board[g.posY][g.posX-1].isObstruction {
			return true
		}
	}
	return false
}

func checkRoomBounds(row, col int, direction string) error {
	newRow, newCol := row, col
	switch direction {
	case "N":
		newRow = row - 1
	case "E":
		newCol = col + 1
	case "S":
		newRow = row + 1
	case "W":
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
	for i, row := range areaMap {
		for j, val := range row {
			switch val.direction {
			case dirN:
				return Guard{i, j, "N", 1}, nil
			case dirE:
				return Guard{i, j, "E", 1}, nil
			case dirS:
				return Guard{i, j, "S", 1}, nil
			case dirW:
				return Guard{i, j, "W", 1}, nil
			}

		}
	}
	return Guard{}, errors.New("No guard on the map")
}

/*if slices.ContainsFunc(row, func(direction string) bool {
	switch direction {
	case
		dirN,
		dirE,
		dirS,
		dirW:
		return true
	}
return false
}) {*/

func readFile(fileName string) (string, int) {
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error opening file: ", err)
		return "", 0
	}
	var content string = ""

	defer file.Close()

	var lineCount int = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineCount += 1
		content += line + "\n"
	}

	return strings.TrimSuffix(content, "\n"), lineCount
}
