package main

import (
	"fmt"
	"github.com/berryjam/leetcode/util"
	"strings"
)

func findOrder(numCourses int, prereuisites [][]int) []int {
	if numCourses == 0 {
		return nil
	}

	adjacentMatrix := make([][]int, numCourses)
	colIndegree := make(map[int]int)
	for i := 0; i < len(adjacentMatrix); i++ {
		adjacentMatrix[i] = make([]int, numCourses)
		colIndegree[i] = 0
	}

	for i := 0; i < len(prereuisites); i++ {
		adjacentMatrix[prereuisites[i][1]][prereuisites[i][0]] = 1
		colIndegree[prereuisites[i][0]]++
	}

	res := make([]int, 0)

	for len(colIndegree) > 0 {
		hasCircle := true
		for col, indegree := range colIndegree {
			if indegree == 0 {
				hasCircle = false
				res = append(res, col)
				for k := 0; k < len(adjacentMatrix[col]); k++ {
					if adjacentMatrix[col][k] == 1 {
						colIndegree[k]--
						adjacentMatrix[col][k] = 0
					}
				}
				delete(colIndegree, col)

			}
		}

		if hasCircle {
			return nil
		}
	}

	return res
}

func main() {
	//res := findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}})
	//fmt.Printf("%s", strings.Join(util.IntArray2StringArray(res), ","))

	res := findOrder(3, [][]int{{0, 1}, {0, 2}, {1, 2}})
	fmt.Printf("%s", strings.Join(util.IntArray2StringArray(res), ","))
}
