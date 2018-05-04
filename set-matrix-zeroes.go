package main

import (
	"fmt"
	"github.com/berryjam/leetcode-golang/util"
	"strings"
)

func setZeroes(matrix [][]int) {
	rowMap := make(map[int]bool)
	colMap := make(map[int]bool)
	for m := range matrix {
		for n := range matrix[m] {
			if matrix[m][n] == 0 {
				rowMap[m]=true
				colMap[n] = true
			}
		}
	}

	for m := range matrix {
		if _,inRow:= rowMap[m];inRow{
			for col := 0; col < len(matrix[m]); col++ {
				matrix[m][col] = 0
			}
		}
	}

	for n := range matrix[0]{
		if _,inCol := colMap[n];inCol{
			for row := 0; row < len(matrix); row++ {
				matrix[row][n] = 0
			}
		}
	}
}


func setMatrixZeroes(matrix [][]int, m, n int) {
	for col := 0; col < len(matrix[m]); col++ {
		matrix[m][col] = 0
	}

	for row := m + 1; row < len(matrix); row++ {
		if matrix[row][n] != 0 {
			matrix[row][n] = 0
		} else {
			break
		}
	}
}

func printMatrix(matrix [][]int) {
	fmt.Println("[")
	for i, row := range matrix {
		rowArr := util.IntArray2StringArray(row)
		if i != len(matrix)-1 {
			fmt.Printf(" [%s],\n", strings.Join(rowArr, ","))
		} else {
			fmt.Printf(" [%s]\n", strings.Join(rowArr, ","))
		}
	}
	fmt.Println("]")
}

func main() {
	//matrix := [][]int{
	//	{1, 1, 1},
	//	{1, 0, 1},
	//	{1, 1, 1,},
	//}
	//setZeroes(matrix)
	//printMatrix(matrix)
	//
	//matrix = [][]int{
	//	{0, 1, 2, 0},
	//	{3, 4, 5, 2},
	//	{1, 3, 1, 5},
	//}
	//setZeroes(matrix)
	//printMatrix(matrix)

	matrix := [][]int{
		{0, 0, 0, 5},
		{4, 3, 1, 4},
		{0, 1, 1, 4},
		{1, 2, 1, 3},
		{0, 0, 1, 1},
	}
	setZeroes(matrix)
	printMatrix(matrix)
}
