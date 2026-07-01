import (
	"slices"
)

func isValidSudoku(board [][]byte) bool {
	flag := true 

	for _, row := range board {
		for j, elem := range row {
			if unicode.IsDigit(rune(elem)) {
				if slices.Contains(row[:j], elem) {
					flag = false
					break
				}
			}
		}

		if !flag {
			return flag
		}
	}

	for j := range 9 {
		curCol := make([]byte, 0, 9)
		for i := range 9 {
			if unicode.IsDigit(rune(board[i][j])) {
				if slices.Contains(curCol, board[i][j]) {
					flag = false
				}
			}
			curCol = append(curCol, board[i][j])
		}

		if !flag {
			return flag
		}
	}

	j := 0 
	for j < 9 {
		i := 0
		for i < 9 {
			curSquare := make([]byte, 0, 9)

			//square 3x3
			for b := j; b < j+3; b++ {
				for a := i; a < i+3; a++ {
					if unicode.IsDigit(rune(board[a][b])) {
						if slices.Contains(curSquare, board[a][b]) {
							flag = false 
							break
						}
					}

					curSquare = append(curSquare, board[a][b])
				}

				if !flag {
					return flag
				}
			}
			i += 3
		}
		j += 3
	}

	return flag
}
