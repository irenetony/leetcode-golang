package main

import (
	"fmt"
	"github.com/berryjam/leetcode-golang/util"
	"strings"
)

type Direction int

const (
	UP   Direction = iota + 1
	DOWN
)

func findDiagonalOrder(matrix [][]int) []int {
	if matrix == nil || len(matrix) == 0 {
		return []int{}
	}

	M := len(matrix)
	N := len(matrix[0])

	if M == 0 || N == 0 {
		return []int{}
	}

	var res []int
	direction := UP
	for i := M - 1; i >= 1-N; i-- {
		minX := 0
		if -i > minX {
			minX = -i
		}
		maxX := N - 1
		if maxX > M-1-i {
			maxX = M - 1 - i
		}
		switch direction {
		case UP:
			for j := minX; j <= maxX; j++ {
				res = append(res, matrix[M-1-(j+i)][j])
			}
			direction = DOWN
		case DOWN:
			for j := maxX; j >= minX; j-- {
				res = append(res, matrix[M-1-(j+i)][j])
			}
			direction = UP
		}
	}

	return res
}

func TestFindDiagonalOrder(matrix [][]int) {
	res := findDiagonalOrder(matrix)
	fmt.Println("Input:")
	fmt.Println("[")
	if matrix != nil && len(matrix) > 0 {
		for idx, arr := range matrix {
			if arr != nil && len(arr) > 0 {
				strArr := util.IntArray2StringArray(arr)
				if idx != len(matrix)-1 {
					fmt.Printf(" [ %s],\n", strings.Join(strArr, ", "))
				} else {
					fmt.Printf(" [ %s]\n", strings.Join(strArr, ", "))
				}
			}
		}
	}
	fmt.Println("]")
	resStrArr := util.IntArray2StringArray(res)
	fmt.Printf("Output:  [%s]\n", strings.Join(resStrArr, ","))
}

func main() {
	matrix1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	TestFindDiagonalOrder(matrix1)
}
