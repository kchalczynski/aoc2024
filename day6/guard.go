package main

import (
	"fmt"
	"slices"
)

type Guard struct {
	posY, posX   int
	direction    string
	visitedCount int
	visitedList  []Position
}

func (g *Guard) Patrol(room Area) {
	for {
		err := g.move(room)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}

func (g *Guard) turnR() {
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

func (g *Guard) checkAhead(room Area) (bool, error) {

	newRow, newCol := g.posY, g.posX
	switch g.direction {
	case dirN:
		newRow = g.posY - 1
	case dirE:
		newCol = g.posX + 1
	case dirS:
		newRow = g.posY + 1
	case dirW:
		newCol = g.posX - 1
	}

	if newRow < 0 || newCol < 0 || newRow >= BaseRoom.sizeY || newCol >= BaseRoom.sizeX {
		return false, fmt.Errorf("Out of room bounds; newRow: %d, newCol: %d", newRow, newCol)
	}

	if !room.board[newRow][newCol].isObstruction {
		return true, nil
	} else {
		return false, nil
	}
}

func (g *Guard) move(room Area) error {

	canGoAhead, err := g.checkAhead(room)
	if err != nil {
		return err
	}
	if !canGoAhead {
		g.turnR()
	}

	// setting next position based on direction could be in separate method
	// out of bounds check probably could be implemented here as well:
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

		// TODO: direction might have to be changed earlier: after the turn, before the move, so Guard is not pointing towards obstruction

		room.board[g.posY][g.posX].direction = g.direction
		g.visitedList = append(g.visitedList, room.board[g.posY][g.posX])
		g.visitedCount++
	}

	return nil
}

func (g *Guard) CopyGuard(moveIndex int) *Guard {
	if moveIndex == 0 || moveIndex > len(g.visitedList) {
		return nil
	}

	tempGuard := &Guard{visitedList: make([]Position, moveIndex)}

	copy(tempGuard.visitedList, g.visitedList[:moveIndex])

	lastPosition := tempGuard.visitedList[moveIndex-1]
	tempGuard.posY, tempGuard.posX = lastPosition.row, lastPosition.col
	tempGuard.direction = lastPosition.direction
	tempGuard.visitedCount = len(tempGuard.visitedList)

	return tempGuard
}

func (g *Guard) isLoop() bool {

	if len(g.visitedList) > 1 {
		if slices.ContainsFunc(g.visitedList[:len(g.visitedList)-1], func(position Position) bool {
			return position.col == g.posX && position.row == g.posY && position.direction == g.direction
		}) {
			return true
		}
	}
	return false
}

// Adding to visitedList inside original method was problematic, maybe ill fix it later, I'd have to refactor part 1 sol
func (g *Guard) move2(room Area) error {

	canGoAhead, err := g.checkAhead(room)
	if err != nil {
		return err
	}
	if !canGoAhead {
		g.turnR()
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
		g.visitedCount++
	}

	return nil
}

func (g *Guard) countPossibleLoops() int {
	loopCount := 0
	for i := 1; i < g.visitedCount; i++ {
		fmt.Println("iter", i)
		// create clean copy of Room with default obstructions
		tempRoom := BaseRoom.CopyRoom()
		//tempRoom.PrintRoom()

		// create clean copy of Guard with moves up to that iteration
		// so his visited list is limited to "i"th move

		// first deep copy visited list, slice of struct cannot be copied directly
		tempVisitedList := make([]Position, len(g.visitedList))
		copy(tempVisitedList, g.visitedList)

		// then copy guard itself and set his current position
		tempGuard := g.CopyGuard(i)
		//tempGuard.posY, tempGuard.posX = tempVisitedList[i-1].row, tempVisitedList[i-1].col

		// mark guards route in room to this iteration
		tempRoom.markRouteInRoom(tempGuard)
		//tempRoom.PrintRoom()

		// mark new Obstruction on Guards next position
		tempRoom.markNewObstruction(g.visitedList[i])

		// simulate guards route with that new Obstruction
		if tempGuard.checkLoopInNewRoute(tempRoom) {
			loopCount++
			//fmt.Println("iter", i)
			println("OOPS")
			//tempRoom.PrintRoom()
		}

	}
	return loopCount
}

func (g *Guard) checkLoopInNewRoute(room Area) bool {

	// execute until guard is out of room bounds or there is a loop
	for {
		err := g.move2(room)
		if err != nil {
			fmt.Println(err)
			return false
		} else {

			// TODO: Najpierw sprawdzam aktualną pozycję z resztą na liście, potem ruch+dodanie do listy
			if g.isLoop() {
				return true
			} else {
				// if new obstruction does not result in a loop, add position as visited and place obstruction at the next position
				g.visitedList = append(g.visitedList, Position{g.posY, g.posX, true, false, g.direction})
			}
		}
	}
}
