package main

import (
	"fmt"
	"strings"
	"leetcode/util"
	ds "leetcode/datastructure"
)

type Pos struct {
	i int
	j int
}

func Solve(board [][]rune) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'O' {
				paintBoard(board, i, j)
			}
		}
	}
}

func paintBoard(board [][]rune, i, j int) {
	queue := ds.NewQueue()
	queue.Enqueue(Pos{i, j})
	regions := make([]Pos, 0)
	toPaint := true
	visitBoard := make([][]bool, len(board))
	for i := range visitBoard {
		visitBoard[i] = make([]bool, len(board[0]))
	}
	visitBoard[i][j] = true
	for !queue.IsEmpty() {
		tmp, err := queue.Dequeue()
		checkErr(err)
		curPos := tmp.(Pos)
		if curPos.i != 0 && curPos.i != len(board)-1 && curPos.j != 0 && curPos.j != len(board[0])-1 {
			regions = append(regions, curPos)
			// up
			if curPos.i-1 >= 0 && board[curPos.i-1][j] == 'O' && !visitBoard[curPos.i-1][curPos.j] {
				queue.Enqueue(Pos{curPos.i - 1, curPos.j})
				visitBoard[curPos.i-1][curPos.j] = true
			}
			// left
			if curPos.j-1 >= 0 && board[curPos.i][curPos.j-1] == 'O' && !visitBoard[curPos.i][curPos.j-1] {
				queue.Enqueue(Pos{curPos.i, curPos.j - 1})
				visitBoard[curPos.i][curPos.j-1] = true
			}
			// right
			if curPos.j+1 < len(board[0]) && board[curPos.i][curPos.j+1] == 'O' && !visitBoard[curPos.i][curPos.j+1] {
				queue.Enqueue(Pos{curPos.i, curPos.j + 1})
				visitBoard[curPos.i][curPos.j+1] = true
			}
			// down
			if i+1 < len(board) && board[curPos.i+1][curPos.j] == 'O' && !visitBoard[curPos.i+1][curPos.j] {
				queue.Enqueue(Pos{curPos.i + 1, curPos.j})
				visitBoard[curPos.i+1][curPos.j] = true
			}
		} else {
			toPaint = false
			break
		}
	}

	if toPaint {
		for _, pos := range regions {
			board[pos.i][pos.j] = 'X'
		}
	}
}

func TestSolve(board [][]rune) {
	Solve(board)

	for _, outter := range board {
		fmt.Printf("%s\n", strings.Join(util.RuneArray2StringArray(outter), "  "))
	}
}

func main() {
	//board1 := [][]rune{
	//	{'X', 'X', 'X', 'X', },
	//	{'X', 'O', 'O', 'X', },
	//	{'X', 'X', 'O', 'X', },
	//	{'X', 'O', 'X', 'X', },
	//}
	//TestSolve(board1)
	//
	//fmt.Println("")

	board2 := [][]rune{
		{'X', 'X', 'X', 'X', 'X', 'X', 'X'},
		{'X', 'O', 'X', 'X', 'O', 'O', 'X'},
		{'O', 'X', 'X', 'O', 'O', 'O', 'X'},
		{'X', 'O', 'O', 'O', 'O', 'X', 'X'},
		{'O', 'X', 'X', 'O', 'X', 'O', 'X'},
		{'O', 'X', 'O', 'O', 'O', 'O', 'X'},
		{'O', 'O', 'X', 'X', 'X', 'X', 'X'},
	}
	TestSolve(board2)

	fmt.Println("")

	board3 := [][]rune{
		{'X', 'O', 'X', 'X'},
		{'X', 'X', 'O', 'X'},
		{'O', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
		{'O', 'X', 'X', 'O'},
	}
	TestSolve(board3)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
