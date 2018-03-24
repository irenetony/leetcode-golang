package main

import (
	"fmt"
	"strings"
	"github.com/berryjam/leetcode/util"
)

func jump(nums []int) int {
	res := 0

	if nums == nil || len(nums) == 0 {
		return res
	} else if len(nums) == 1 {
		return 0
	}

	curIdx := 0
	res = 0
	for curIdx < len(nums) {
		curNextJumpLength := nums[curIdx]
		nextMaxJumpLength := 0
		nextIdx := 0
		for i := 1; i <= curNextJumpLength; i++ {
			if curNextJumpLength+curIdx >= len(nums)-1 {
				res++
				return res
			} else {
				nextJumpLength := nums[i+curIdx]
				if curIdx+i+nextJumpLength > nextMaxJumpLength {
					nextMaxJumpLength = curIdx+i+nextJumpLength
					nextIdx = i
				}
			}
		}
		res++
		curIdx += nextIdx
		if curIdx >= len(nums)-1 {
			return res + 1
		}
	}

	return res
}

func TestJump(nums []int) {
	fmt.Printf("Given array A = [%s]. ", strings.Join(util.IntArray2StringArray(nums), ","))
	fmt.Printf("The result is: %d\n", jump(nums))
}

func main() {
	//A := []int{2, 3, 1, 1, 4}
	//A := []int{0}
	//A := []int{2, 1}
	//A := []int{1, 2, 3}
	//A := []int{1, 3, 2}
	//A := []int{1, 2, 0, 1}
	//A := []int{1, 1, 1, 1}
	A := []int{4, 1, 1, 3, 1, 1, 1}
	TestJump(A)
}
