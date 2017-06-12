package main

import (
	"fmt"
)

func countBattleships(board [][]byte) int {
	res := 0
	inCountingState := false

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if inCountingState {
				if board[i][j] =='.'{
					inCountingState = false
				}else{
					if j-1>=0&&board[i][j-1]=='x'{
						continue
					}
				}
			}else{
				if board[i][j]=='x'{
					if i-1>=0&&board[i-1][j]=='x'{
						continue
					}else{
						res++
						inCountingState = true
					}
				}
			}
		}
	}

	return res
}

func testCountBattleships(board [][]byte) {
	for _, outter := range board {
		res := ""
		for _, inner := range outter {
			res += string(inner)
		}
		fmt.Printf("%s\n", res)
	}
	fmt.Printf("There are %d battleships.\n", countBattleships(board))
}

func main() {
	board1 := [][]byte{
		{'x', '.', '.', 'x'},
		{'.', 'x', '.', 'x'},
		{'.', '.', '.', 'x'},
		{'x', 'x', '.', 'x'},
		{'.', '.', 'x', '.'},
	}
	testCountBattleships(board1)

	board2 := [][]byte{
		{'x', '.', '.', 'x'},
		{'.', '.', '.', 'x'},
		{'.', '.', '.', 'x'},
	}
	testCountBattleships(board2)
}
