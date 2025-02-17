package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
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

var BaseRoom Area

func initializeArea(board [][]string) {
	BaseRoom.sizeX = len(board[0])
	BaseRoom.sizeY = len(board)
	BaseRoom.board = stringsToPositions(board)
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
func (room *Area) markVisited(y, x int) {
	room.board[y][x].isVisited = true
}

func main() {

	var inputFile = "input2.txt"
	var inputContent string
	var inputLinesCount int

	inputContent, inputLinesCount = readFile(inputFile)
	initializeArea(readContentIntoMatrix(inputContent, inputLinesCount))
	guard, err := initializeGuard(BaseRoom.board)

	// Changed to new room variable from global BaseRoom because I need to use modified room to check for loops,
	// but don't want to change the global one
	room := BaseRoom
	if err != nil {
		fmt.Println(err)
	} else {

		room.markVisited(guard.posY, guard.posX)
		guard.Patrol(room)
		println(guard.posX, guard.posY, guard.visitedCount)

	}
	room.printRoom()
	loopCount := guard.countPossibleLoops()
	println(loopCount)

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

func (g *Guard) Patrol(room Area) {
	for {
		err := g.Move(room)
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

func (g *Guard) Move(room Area) error {
	//fmt.Println("test1")
	if err := checkRoomBounds(g.posY, g.posX, g.direction); err != nil {
		return err
	}
	//fmt.Println("test2")
	if !g.CheckAhead(room) {
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

	if !room.board[g.posY][g.posX].isVisited {
		room.markVisited(g.posY, g.posX)

		// TODO: direction might have to be changed before, after the turn, before the move, so Guard is not pointing towards obstruction
		room.board[g.posY][g.posX].direction = g.direction
		g.visitedList = append(g.visitedList, room.board[g.posY][g.posX])
		g.visitedCount++
	}

	return nil
}

func (g *Guard) CheckAhead(room Area) bool {

	switch g.direction {
	case dirN:
		if !room.board[g.posY-1][g.posX].isObstruction {
			return true
		}
	case dirE:
		if !room.board[g.posY][g.posX+1].isObstruction {
			return true
		}
	case dirS:
		if !room.board[g.posY+1][g.posX].isObstruction {
			return true
		}
	case dirW:
		if !room.board[g.posY][g.posX-1].isObstruction {
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
	// can use global BaseRoom below, as dimensions don't change
	if newRow < 0 || newCol < 0 || newRow >= BaseRoom.sizeY || newCol >= BaseRoom.sizeX {
		return fmt.Errorf("Out of room bounds; newRow: %d, newCol: %d", newRow, newCol)
	}
	return nil
}

func (g *Guard) countPossibleLoops() int {
	loopCount := 0
	for i := 1; i < len(g.visitedList); i++ {
		fmt.Println("iter", i)
		// create clean copy of Room with default obstructions
		tempRoom := CopyRoom(BaseRoom)
		tempRoom.printRoom()

		// create clean copy of Guard with moves up to that iteration
		// so his visited list is limited to "i"th move

		// first deep copy visited list, slice of struct cannot be copied directly
		tempVisitedList := make([]Position, len(g.visitedList))
		copy(tempVisitedList, g.visitedList)

		// then copy guard itself and set his current position
		tempGuard := g.CopyGuard(tempVisitedList[0:i])
		tempGuard.posY, tempGuard.posX = tempVisitedList[i-1].row, tempVisitedList[i-1].col

		// mark guards route in room to this iteration
		tempRoom.markRouteInRoom(tempGuard)
		tempRoom.printRoom()

		// mark new Obstruction on Guards next position
		tempRoom.markNewObstruction(g.visitedList[i])

		// simulate guards route with that new Obstruction
		if tempGuard.checkLoopInNewRoute(tempRoom) {
			loopCount++
			//fmt.Println("iter", i)
			println("OOPS")
			tempRoom.printRoom()
		}

	}
	return loopCount
}

func CopyRoom(room Area) Area {
	roomCopy := room
	roomCopy.board = make([][]Position, len(room.board))
	for i := range room.board {
		roomCopy.board[i] = make([]Position, len(room.board[i]))
		copy(roomCopy.board[i], room.board[i])
	}

	for i := range roomCopy.board {
		for j := range roomCopy.board[i] {
			roomCopy.board[i][j].isVisited = false
			roomCopy.board[i][j].direction = ""
		}
	}

	return roomCopy
}

func (room *Area) markNewObstruction(position Position) {
	room.board[position.row][position.col].isObstruction = true
}

func (g *Guard) IsLoop(p Position) bool {
	if slices.Contains(g.visitedList, p) {
		return true
	}
	return false
}

func (g *Guard) checkLoopInNewRoute(room Area) bool {

	// dopóki nie wywali się error albo nie będzie pętli
	for {
		fmt.Println("move")
		err := g.Move(room)
		if err != nil {
			fmt.Println(err)
			return false
		} else {

			// g.posY i g.posX coś nie spełniają swojego zadania

			// chcę sprawdzić, czy ta pozycja była już odwiedzona
			// ale skoro sprawdzam aktualną pozycję - zawsze będzie na liście
			// TODO: Najpierw ruch?
			if g.IsLoop(room.board[g.posY][g.posX]) {
				room.printRoom()
				fmt.Println("is loop")
				return true
			} else {
				// co tu robię
				fmt.Println("else nie loop")
				g.visitedList = append(g.visitedList, Position{g.posY, g.posX, true, false, g.direction})
			}
		}
	}
}
func (g *Guard) CopyGuard(visitedPositions []Position) Guard {
	// dereference

	tempGuard := *g
	tempGuard.visitedList = make([]Position, len(visitedPositions))
	copy(tempGuard.visitedList, visitedPositions)
	return tempGuard
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

func (room *Area) printRoom() {
	var roomMap = ""
	for _, row := range room.board {
		for _, col := range row {
			if col.isVisited {
				roomMap += col.direction
			} else {
				if col.isObstruction {
					roomMap += obstructionMark
				} else {
					roomMap += notVisitedMark
				}
			}
			roomMap += " "
		}
		roomMap += "\n"
	}

	fmt.Println(roomMap)
}

func (room *Area) markRouteInRoom(g Guard) {
	for _, pos := range g.visitedList {
		room.board[pos.row][pos.col].isObstruction = false
		room.board[pos.row][pos.col].isVisited = true
		room.board[pos.row][pos.col].direction = pos.direction
	}
}
