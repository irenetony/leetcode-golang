package main

import "fmt"

func largestRectangleArea(heights []int) int {
	res := 0
	for i := 0; i < len(heights); i++ {
		minHeight := heights[i]
		for j := i; j < len(heights); j++ {
			if minHeight > heights[j] {
				minHeight = heights[j]
			}
			if res < (j-i+1)*minHeight {
				res = (j - i + 1) * minHeight
			}
		}
	}

	return res
}

func main() {
	fmt.Printf("res=%d\n", largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
}
