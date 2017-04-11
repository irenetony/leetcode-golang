package main

import (
	"strconv"
	"strings"
	"fmt"
)

var cacheFlagMatrix [][]bool

func pacificAtlantic(matrix [][]int) [][]int {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}

	res := make([][]int, 0)
	cacheFlagMatrix = make([][]bool, len(matrix))
	for i := 0; i < len(cacheFlagMatrix); i++ {
		cacheFlagMatrix[i] = make([]bool, len(matrix[0]))
	}

	width := len(matrix)
	height := len(matrix[0])
	for i := 0; i < len(matrix); i++ {
		cacheFlagMatrix[i] = make([]bool, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			if i == 0 || j == 0 {
				if canFlow2Atlantic(matrix, i, j, newVisitedMatrix(width, height)) {
					res = append(res, []int{i, j})
					cacheFlagMatrix[i][j] = true
				}
			} else if i == len(matrix)-1 || j == len(matrix[0])-1 {
				if canFlow2Pacific(matrix, i, j, newVisitedMatrix(width, height)) {
					res = append(res, []int{i, j})
					cacheFlagMatrix[i][j] = true
				}
			} else if canFlow2Atlantic(matrix, i, j, newVisitedMatrix(width, height)) && canFlow2Pacific(matrix,
				i, j, newVisitedMatrix(width, height)) {
				res = append(res, []int{i, j})
				cacheFlagMatrix[i][j] = true
			}
		}
	}

	return res
}

func newVisitedMatrix(i, j int) [][] bool {
	res := make([][]bool, i)

	for i := 0; i < len(res); i++ {
		res[i] = make([]bool, j)
	}

	return res
}

func canFlow2Pacific(matrix [][]int, i, j int, visitedMatrix [][]bool) bool {
	if i == 0 || j == 0 {
		return true
	}

	upRes := false
	if i-1 >= 0 {
		if matrix[i][j] >= matrix[i-1][j] {
			if cacheFlagMatrix[i-1][j] {
				cacheFlagMatrix[i][j] = true
				return true
			} else {
				if !visitedMatrix[i-1][j] {
					visitedMatrix[i-1][j] = true
					upRes = canFlow2Pacific(matrix, i-1, j, visitedMatrix)
				}
			}
		}
	}

	downRes := false
	if i+1 < len(matrix) {
		if matrix[i][j] >= matrix[i+1][j] {
			if cacheFlagMatrix[i+1][j] {
				cacheFlagMatrix[i+1][j] = true
				return true
			} else {
				if !visitedMatrix[i+1][j] {
					visitedMatrix[i+1][j] = true
					downRes = canFlow2Pacific(matrix, i+1, j, visitedMatrix)
				}
			}
		}
	}

	leftRes := false
	if j-1 >= 0 {
		if matrix[i][j] >= matrix[i][j-1] {
			if cacheFlagMatrix[i][j-1] {
				cacheFlagMatrix[i][j-1] = true
				return true
			} else {
				if !visitedMatrix[i][j-1] {
					visitedMatrix[i][j-1] = true
					leftRes = canFlow2Pacific(matrix, i, j-1, visitedMatrix)
				}
			}
		}
	}

	rightRes := false
	if j+1 < len(matrix[0]) {
		if matrix[i][j] >= matrix[i][j+1] {
			if cacheFlagMatrix[i][j+1] {
				cacheFlagMatrix[i][j+1] = true
				return true
			} else {
				if !visitedMatrix[i][j+1] {
					visitedMatrix[i][j+1] = true
					rightRes = canFlow2Pacific(matrix, i, j+1, visitedMatrix)
				}
			}
		}
	}

	return upRes || downRes || leftRes || rightRes
}

func canFlow2Atlantic(matrix [][]int, i, j int, visitedMatrix [][]bool) bool {
	if i == len(matrix)-1 || j == len(matrix[0])-1 {
		return true
	}

	upRes := false
	if i-1 >= 0 {
		if matrix[i][j] >= matrix[i-1][j] {
			if cacheFlagMatrix[i-1][j] {
				cacheFlagMatrix[i][j] = true
				return true
			} else {
				if !visitedMatrix[i-1][j] {
					visitedMatrix[i-1][j] = true
					upRes = canFlow2Atlantic(matrix, i-1, j, visitedMatrix)
				}
			}
		}
	}

	downRes := false
	if i+1 < len(matrix) {
		if matrix[i][j] >= matrix[i+1][j] {
			if cacheFlagMatrix[i+1][j] {
				cacheFlagMatrix[i+1][j] = true
				return true
			} else {
				if !visitedMatrix[i+1][j] {
					visitedMatrix[i+1][j] = true
					downRes = canFlow2Atlantic(matrix, i+1, j, visitedMatrix)
				}
			}
		}
	}

	leftRes := false
	if j-1 >= 0 {
		if matrix[i][j] >= matrix[i][j-1] {
			if cacheFlagMatrix[i][j-1] {
				cacheFlagMatrix[i][j-1] = true
				return true
			} else {
				if !visitedMatrix[i][j-1] {
					visitedMatrix[i][j-1] = true
					leftRes = canFlow2Atlantic(matrix, i, j-1, visitedMatrix)
				}
			}
		}
	}

	rightRes := false
	if j+1 < len(matrix[0]) {
		if matrix[i][j] >= matrix[i][j+1] {
			if cacheFlagMatrix[i][j+1] {
				cacheFlagMatrix[i][j+1] = true
				return true
			} else {
				if !visitedMatrix[i][j+1] {
					visitedMatrix[i][j+1] = true
					rightRes = canFlow2Atlantic(matrix, i, j+1, visitedMatrix)
				}
			}
		}
	}

	return upRes || downRes || leftRes || rightRes
}

func main() {
	matrix := [][]int{
		{1, 2, 2, 3, 5},
		{3, 2, 3, 4, 4},
		{2, 4, 5, 3, 1},
		{6, 7, 1, 4, 5},
		{5, 1, 1, 2, 4},
	}

	res := pacificAtlantic(matrix)

	tmpArray := make([]string, 0)
	for _, val := range res {
		tmp := "[" + strconv.Itoa(val[0]) + ", " + strconv.Itoa(val[1]) + "]"
		tmpArray = append(tmpArray, tmp)
	}

	output := "[" + strings.Join(tmpArray, ",") + "]"
	fmt.Printf("Return:\n%s", output)
}
