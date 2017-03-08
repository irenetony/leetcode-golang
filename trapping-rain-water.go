package main

import (
	"fmt"
	"leetcode/util"
	"strings"
)

func Trap(height []int) int {
	left := 0
	right := len(height) - 1
	res := 0
	if left >= right {
		return res
	}
	leftHeight := height[left]
	rightHeight := height[right]
	for left < right {
		if leftHeight < rightHeight {
			left++
			if leftHeight > height[left] {
				res += leftHeight - height[left]
			} else {
				leftHeight = height[left]
			}
		} else {
			right--
			if rightHeight > height[right] {
				res += rightHeight - height[right]
			} else {
				rightHeight = height[right]
			}
		}
	}
	return res
}

func TestTrap(height []int) {
	fmt.Printf("Given [%s] ,return %d.\n", strings.Join(util.IntArray2StringArray(height), ","), Trap(height))
}

func main() {
	height1 := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	TestTrap(height1)
}
