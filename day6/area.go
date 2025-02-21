package main

import "fmt"

type Area struct {
	sizeX, sizeY int
	board        [][]Position
}

func (room *Area) markVisited(y, x int) {
	room.board[y][x].isVisited = true
}

func (room *Area) CopyRoom() Area {
	roomCopy := Area{
		sizeX: room.sizeX,
		sizeY: room.sizeY,
		board: make([][]Position, len(room.board)),
	}

	for i := range room.board {
		roomCopy.board[i] = make([]Position, len(room.board[i]))
		copy(roomCopy.board[i], room.board[i])
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

func (room *Area) markRouteInRoom(g *Guard) {
	for _, pos := range g.visitedList {
		room.board[pos.row][pos.col].isObstruction = false
		room.board[pos.row][pos.col].isVisited = true
		room.board[pos.row][pos.col].direction = pos.direction
	}
}

func (room *Area) getBoardAsMultilineString() string {
	var roomMap = ""
	for i := 0; i < len(room.board[0]); i++ {
		roomMap += fmt.Sprintf("%d ", i)
	}
	roomMap += "\n"
	for i, row := range room.board {
		roomMap += fmt.Sprintf("%d ", i)
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
	return roomMap
}

func (room *Area) PrintRoom() {
	fmt.Println(room.getBoardAsMultilineString())
}
