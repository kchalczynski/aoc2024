package day6

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
	}
	return roomCopy
}

func (room *Area) markNewObstruction(position Position) {
	room.board[position.row][position.col].isObstruction = true
}

func (room *Area) markPosInRoom(g *Guard) {

	room.board[g.posY][g.posX].isObstruction = false
	room.board[g.posY][g.posX].isVisited = true
	room.board[g.posY][g.posX].direction = g.direction

}

func (room *Area) getBoardAsMultilineString() string {
	var roomMap = "  "
	for i := 0; i < len(room.board[0]); i++ {
		roomMap += fmt.Sprintf("%d ", i)
	}
	roomMap += "\n"
	for i, row := range room.board {
		roomMap += fmt.Sprintf("%d ", i)
		for j, col := range row {
			if col.isVisited {
				roomMap += room.board[i][j].direction
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
