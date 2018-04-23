package main

import (
	"fmt"
	ds "github.com/berryjam/leetcode-golang/datastructure"
)

func largestRectangleArea(heights []int) int {
	stack := ds.NewStack()
	maxArea := 0
	for i := 0; i <= len(heights); i++ {
		var h int
		if i == len(heights) {
			h = 0
		} else {
			h = heights[i]
		}

		if stack.IsEmpty() {
			stack.Push(i)
		} else if peekHeight := heights[stack.Top().(int)]; h >= peekHeight {
			stack.Push(i)
		} else {
			t := stack.Pop().(int)
			var width int
			if stack.IsEmpty() {
				width = i
			} else {
				width = i - 1 - stack.Top().(int)
			}
			maxArea = maxVal(maxArea, heights[t]*width)
			i--
		}
	}

	return maxArea
}

func maxVal(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Printf("res=%d\n", largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
}
