package main

import (
	"fmt"
	"math"
)

func maxArea(height []int) int {
	res := 0

	lh := height[0]
	rh := height[len(height)-1]
	i, j := 0, len(height)-1

	for i < j {
		area := int(math.Min(float64(lh), float64(rh))) * (j - i)
		if area > res {
			res = area
		}

		if lh < rh {
			for i < j && height[i] <= lh {
				i++
			}

			if i < j {
				lh = height[i]
			}
		} else {
			for i < j && height[j] <= rh {
				j--
			}

			if i < j {
				rh = height[j]
			}
		}
	}

	return res
}

func main() {
	height := []int{1, 2, 3}
	area := maxArea(height)
	fmt.Printf("area:%d", area)
}
