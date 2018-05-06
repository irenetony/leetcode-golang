package main

import "fmt"

func canJump(nums []int) bool {
	curIdx := 0
	for curIdx < len(nums) {
		nextMaxIdx := curIdx + nums[curIdx]
		nextIdx := nextMaxIdx
		for offset := 1; offset <= nums[curIdx]; offset++ {
			if curIdx+offset >= len(nums)-1 {
				return true
			}
			if curIdx+offset+nums[curIdx+offset] >= len(nums)-1 {
				return true
			}
			if curIdx+offset+nums[curIdx+offset] > nextMaxIdx {
				nextMaxIdx = curIdx + offset + nums[curIdx+offset]
				nextIdx = curIdx + offset
				if nextMaxIdx >= len(nums)-1 {
					return true
				}
			}
		}
		curIdx = nextIdx
		if curIdx >= len(nums)-1 {
			return true
		}
		if nums[curIdx] == 0 {
			return false
		}
	}

	return false
}

func main() {
	fmt.Printf("res=%+v\n", canJump([]int{2, 3, 1, 1, 4}))
	fmt.Printf("res=%+v\n", canJump([]int{3, 2, 1, 0, 4}))
	fmt.Printf("res=%+v\n", canJump([]int{0}))
	fmt.Printf("res=%+v\n", canJump([]int{5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0}))
}
