package day6

type Position struct {
	row, col                 int
	isVisited, isObstruction bool
	direction                string
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
