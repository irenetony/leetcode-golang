package main

import (
	"fmt"
)

func convert(s string, numRows int) string {
	if s == "" || len(s) == 0 {
		return ""
	}

	if numRows == 1 {
		return s
	}

	totalColumn, _ := calZigColumn(len(s), numRows)

	grid := make([][]rune, numRows)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]rune, totalColumn)
	}

	for idx, char := range s {
		var curRow int
		curColumn, lastZigLength := calZigColumn(idx+1, numRows)
		if lastZigLength == 0 {
			curRow = 1
		} else if lastZigLength >= 1 && lastZigLength <= numRows {
			curRow = lastZigLength - 1
		} else {
			curRow = numRows - 1 - (lastZigLength - numRows)
		}
		grid[curRow][curColumn-1] = char
	}

	res := ""
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] != 0 {
				res += string(grid[i][j])
			}
		}
	}

	return res
}

func calZigColumn(len, numRows int) (int, int) {
	zigLength := 2*numRows - 2

	zigNums := len / zigLength

	lastZigLength := len - zigNums*zigLength

	totalColumn := (numRows - 1) * zigNums

	if lastZigLength >= 1 && lastZigLength <= numRows {
		totalColumn += 1
	} else if lastZigLength > numRows {
		totalColumn += (lastZigLength - numRows) + 1
	}

	return totalColumn, lastZigLength
}

func TestConvert(s string, numRows int) {
	res := convert(s, numRows)
	if s == "PAYPALISHIRING" {
		if res != "PAHNAPLSIIGYIR" {
			fmt.Println("Wrong!")
		} else {
			fmt.Println("Correct!")
		}
	}
}

func main() {
	s1 := "PAYPALISHIRING"
	TestConvert(s1, 3)

	s2 := "ABCDE"
	TestConvert(s2, 4)
}
