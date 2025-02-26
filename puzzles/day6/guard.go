package day6

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
	} else {
		// TODO FIX: TURN AND MOVE CANNOT HAPPEN IN THE SAME ITERATION/METHOD CALL

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

			g.visitedCount++
		}
	}
	visitedRoom := Position{g.posY, g.posX, true, false, g.direction}
	g.visitedList = append(g.visitedList, visitedRoom)

	return nil
}

func (g *Guard) CopyGuard(moveIndex int) *Guard {
	if moveIndex == 0 || moveIndex >= len(g.visitedList) {
		return nil
	}

	tempGuard := &Guard{visitedList: make([]Position, moveIndex)}

	copy(tempGuard.visitedList, g.visitedList[:moveIndex])

	lastPosition := tempGuard.visitedList[len(tempGuard.visitedList)-1]
	tempGuard.posY, tempGuard.posX = lastPosition.row, lastPosition.col
	tempGuard.direction = lastPosition.direction

	return tempGuard
}

/*
// Adding to visitedList inside original method was problematic, maybe ill fix it later, I'd have to refactor part 1 sol
func (g *Guard) move2(room Area) error {

	canGoAhead, err := g.checkAhead(room)
	if err != nil {
		return err
	}
	if !canGoAhead {
		g.turnR()
		return nil
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
}*/

func (g *Guard) isLoop() bool {

	// At least 4 moves/turns required for loop
	if len(g.visitedList) > 1 {
		if slices.ContainsFunc(g.visitedList[:len(g.visitedList)-1], func(position Position) bool {
			return position.col == g.posX &&
				position.row == g.posY &&
				position.direction == g.direction
		}) {
			return true
		}
	}
	return false
}

func (g *Guard) countPossibleLoops() int {
	loopCount := 0
	// create clean copy of Room with default obstructions
	baseRoomForMove := BaseRoom.CopyRoom()

	for i := 1; i < len(g.visitedList); i++ {
		// cannot put obstruction at guards starting position, no matter what iteration
		if (g.visitedList[i].row == g.visitedList[0].row && g.visitedList[i].col == g.visitedList[0].col) ||
			g.visitedList[i].isObstruction {
			continue
		}

		if i%250 == 0 {
			fmt.Println(i, "loop count:", loopCount)
		}

		//TODO: Maybe I could optimize by making BaseRoom always updated up to nth iteration
		// then tempRoom doesnt have to be updated every time from the beginning
		// same with Guards position

		// copy guard itself and set his current position
		tempGuard := g.CopyGuard(i)

		baseRoomForMove.markPosInRoom(tempGuard)
		tempRoom := baseRoomForMove.CopyRoom()

		// check if new obstruction position hasn't been visited by guard before
		// if it did, position is invalid, as obstruction must be placed before guard starts patrol
		// move to the next iteration
		if slices.ContainsFunc(tempGuard.visitedList, func(position Position) bool {
			return position.row == g.visitedList[i].row && position.col == g.visitedList[i].col
		}) {
			continue
		}

		// mark new Obstruction on Guards next position
		tempRoom.markNewObstruction(g.visitedList[i])

		// simulate guards route with that new Obstruction
		if tempGuard.checkLoopInNewRoute(tempRoom) {
			loopCount++
		}
	}

	return loopCount
}

func (g *Guard) checkLoopInNewRoute(room Area) bool {

	// execute until guard is out of room bounds or there is a loop
	for {
		err := g.move(room)
		if err != nil {
			return false
		} else {
			if g.isLoop() {
				return true
			}
		}
	}
}
