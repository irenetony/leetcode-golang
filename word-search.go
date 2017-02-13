package main

import (
	"fmt"
	"unicode/utf8"
)

func Exist(board [][]rune, word string) bool {
	flagBoard := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		flagBoard[i] = make([]bool, len(board[i]))
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if existRecur(board, word[0:], i, j, flagBoard) {
				return true
			}
		}
	}

	return false
}

func existRecur(board [][]rune, word string, i, j int, flagBoard [][]bool) bool {
	if len(word) == 1 {
		runeValue, _ := utf8.DecodeRune([]byte(word))
		if runeValue == board[i][j] {
			return true
		} else {
			return false
		}
	}

	runeValue, _ := utf8.DecodeRune([]byte{word[0]})

	if board[i][j] != runeValue {
		return false
	}

	// right
	if j+1 < len(board[0]) {
		if !flagBoard[i][j+1] {
			r := copyArray(flagBoard)
			r[i][j+1] = true
			right := existRecur(board, word[1:], i, j+1, r)
			if right {
				return true
			}
		}
	}

	// left
	if j-1 >= 0 {
		if !flagBoard[i][j-1] {
			l := copyArray(flagBoard)
			l[i][j-1] = true
			left := existRecur(board, word[1:], i, j-1, l)
			if left {
				return true
			}
		}
	}

	// up
	if i-1 >= 0 {
		if !flagBoard[i-1][j] {
			u := copyArray(flagBoard)
			u[i-1][j] = true
			up := existRecur(board, word[1:], i-1, j, u)
			if up {
				return true
			}
		}
	}

	// down
	if i+1 < len(board) {
		if !flagBoard[i+1][j] {
			d := copyArray(flagBoard)
			d[i+1][j] = true
			down := existRecur(board, word[1:], i+1, j, d)
			if down {
				return true
			}
		}
	}

	return false
}

func copyArray(array [][]bool) (res [][]bool) {
	res = make([][]bool, len(array))
	for i := 0; i < len(array); i++ {
		res[i] = make([]bool, len(array[i]))
		for j := 0; j < len(array[i]); j++ {
			res[i][j] = array[i][j]
		}
	}
	return
}

func TestExist() {
	board := [][]rune{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word1 := "ABCCED"
	word2 := "SEE"
	word3 := "ABCB"
	word4 := "SFCS"
	word5 := "SFCSA"

	fmt.Printf("word = \"%s\", -> returns %t\n", word1, Exist(board, word1))
	fmt.Printf("word = \"%s\", -> returns %t\n", word2, Exist(board, word2))
	fmt.Printf("word = \"%s\", -> returns %t\n", word3, Exist(board, word3))
	fmt.Printf("word = \"%s\", -> returns %t\n", word4, Exist(board, word4))
	fmt.Printf("word = \"%s\", -> returns %t\n", word5, Exist(board, word5))
}

func main() {
	TestExist()
}
