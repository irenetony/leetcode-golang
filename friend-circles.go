package main

import (
	"fmt"
	"github.com/berryjam/leetcode/util"
	ds "github.com/berryjam/leetcode/datastructure"
)

func findCircleNum(M [][]int) int {
	res := 0
	visitedNums := make([]bool, len(M))
	for i := 0; i < len(M); i++ {
		if !visitedNums[i] {
			res++
		}
		queue := ds.NewQueue()
		queue.Enqueue(i)
		for !queue.IsEmpty() {
			val, err := queue.Dequeue()
			if err != nil {
				panic(err)
			}
			num := val.(int)
			visitedNums[num] = true
			for j := 0; j < len(M); j++ {
				if !visitedNums[j] && M[num][j] == 1 {
					queue.Enqueue(j)
				}
			}
		}
	}

	return res
}

func TestFindCircleNum(M [][]int) {
	fmt.Println("Input:")
	fmt.Println(util.TwoDimensionToString(M))
	fmt.Printf("Output: %d", findCircleNum(M))
	fmt.Println()
}

func main() {
	array1 := [][]int{
		{1, 1, 0},
		{1, 1, 0},
		{0, 0, 1},
	}
	TestFindCircleNum(array1)

	array2 := [][]int{
		{1, 1, 0},
		{1, 1, 1},
		{0, 1, 1},
	}
	TestFindCircleNum(array2)
}
