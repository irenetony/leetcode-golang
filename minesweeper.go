package main

func updateBoard(board [][]byte, click []int) [][]byte {
	cell := board[click[0]][click[1]]
	if cell == 'M' {
		board[click[0]][click[1]] = 'X'
		return board
	}

	if cell == 'E' {
		adjacentMines := countAdjacentMines(board, click)
		if adjacentMines == 0 {
			updateEmptySquareNoAdjacentMines(board, click)
		} else {
			updateEmptySquareWithAdjacentMines(board, click, adjacentMines)
		}
	}

	return board
}

func updateEmptySquareNoAdjacentMines(board [][]byte, click []int) {
	x := click[0]
	y := click[1]
	board[x][y] = 'B'

	// up
	if x-1 >= 0 && board[x-1][y] == 'E' {
		adjacentMines := countAdjacentMines(board, []int{x - 1, y})
		if adjacentMines == 0 {
			updateEmptySquareNoAdjacentMines(board, []int{x - 1, y})
		} else {
			updateEmptySquareWithAdjacentMines(board, []int{x - 1, y}, adjacentMines)
		}
	}

	// up left
	if x-1 >= 0 && y-1 >= 0 && board[x-1][y-1] == 'E' {
		adjacentMines := countAdjacentMines(board, []int{x - 1, y - 1})
		if adjacentMines == 0 {
			updateEmptySquareNoAdjacentMines(board, []int{x - 1, y - 1})
		} else {
			updateEmptySquareWithAdjacentMines(board, []int{x - 1, y - 1}, adjacentMines)
		}
	}

	// left
	if y-1 >= 0 && board[x][y-1] == 'E' {
		adjacentMines := countAdjacentMines(board, []int{x, y - 1})
		if adjacentMines == 0 {
			updateEmptySquareNoAdjacentMines(board, []int{x, y - 1})
		} else {
			updateEmptySquareWithAdjacentMines(board, []int{x, y - 1}, adjacentMines)
		}
	}

	// down left
	if x+1 < len(board) && y-1 >= 0 && board[x+1][y-1] == 'E' {
		adjacentMines := countAdjacentMines(board, []int{x + 1, y - 1})
		if adjacentMines == 0 {
			updateEmptySquareNoAdjacentMines(board, []int{x + 1, y - 1})
		} else {
			updateEmptySquareWithAdjacentMines(board, []int{x + 1, y - 1}, adjacentMines)
		}
	}

	// down
	if x+1 < len(board) && board[x+1][y] == 'E' {
		adjacentMines := countAdjacentMines(board, []int{x + 1, y})
		if adjacentMines == 0 {
			updateEmptySquareNoAdjacentMines(board, []int{x + 1, y})
		} else {
			updateEmptySquareWithAdjacentMines(board, []int{x + 1, y}, adjacentMines)
		}
	}

	// down right
	if x+1 < len(board) && y+1 < len(board[0]) && board[x+1][y+1] == 'E' {
		adjacentMines := countAdjacentMines(board, []int{x + 1, y + 1})
		if adjacentMines == 0 {
			updateEmptySquareNoAdjacentMines(board, []int{x + 1, y + 1})
		} else {
			updateEmptySquareWithAdjacentMines(board, []int{x + 1, y + 1}, adjacentMines)
		}
	}

	// right
	if y+1 < len(board[0]) && board[x][y+1] == 'E' {
		adjacentMines := countAdjacentMines(board, []int{x, y + 1})
		if adjacentMines == 0 {
			updateEmptySquareNoAdjacentMines(board, []int{x, y + 1})
		} else {
			updateEmptySquareWithAdjacentMines(board, []int{x, y + 1}, adjacentMines)
		}
	}

	// up right
	if x-1 >= 0 && y+1 < len(board[0]) && board[x-1][y+1] == 'E' {
		adjacentMines := countAdjacentMines(board, []int{x - 1, y + 1})
		if adjacentMines == 0 {
			updateEmptySquareNoAdjacentMines(board, []int{x - 1, y + 1})
		} else {
			updateEmptySquareWithAdjacentMines(board, []int{x - 1, y + 1}, adjacentMines)
		}
	}
}

func updateEmptySquareWithAdjacentMines(board [][]byte, click []int, mines byte) {
	x := click[0]
	y := click[1]
	board[x][y] = '0' + mines
}

func countAdjacentMines(board [][]byte, click []int) byte {
	var mines byte = 0

	x := click[0]
	y := click[1]

	// up
	if x-1 >= 0 && board[x-1][y] == 'M' {
		mines++
	}

	// up left
	if x-1 >= 0 && y-1 >= 0 && board[x-1][y-1] == 'M' {
		mines++
	}

	// left
	if y-1 >= 0 && board[x][y-1] == 'M' {
		mines++
	}

	// down left
	if x+1 < len(board) && y-1 >= 0 && board[x+1][y-1] == 'M' {
		mines++
	}

	// down
	if x+1 < len(board) && board[x+1][y] == 'M' {
		mines++
	}

	// down right
	if x+1 < len(board) && y+1 < len(board[0]) && board[x+1][y+1] == 'M' {
		mines++
	}

	// right
	if y+1 < len(board[0]) && board[x][y+1] == 'M' {
		mines++
	}

	// up right
	if x-1 >= 0 && y+1 < len(board[0]) && board[x-1][y+1] == 'M' {
		mines++
	}

	return mines
}

func main() {
}
