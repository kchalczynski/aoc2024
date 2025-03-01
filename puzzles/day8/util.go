package day8

func stringsToPositions(board [][]string) [][]Position {
	rows, cols := len(board), len(board[0])
	positions := make([][]Position, rows)
	for i := 0; i < rows; i++ {
		positions[i] = make([]Position, cols)
		for j := 0; j < cols; j++ {
			positions[i][j] = Position{
				row:        i,
				col:        j,
				isAntinode: false,
				isAntenna:  board[i][j] != ".",
				antennaType: func(cell string) string {
					if cell != "." {
						return cell
					}
					return ""
				}(board[i][j]),
			}
		}
	}
	return positions
}
