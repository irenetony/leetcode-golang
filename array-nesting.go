package main

import (
	"fmt"
	"strings"
	"github.com/berryjam/leetcode-golang/util"
)

func arrayNesting(nums []int) int {
	result := 0
	flags := make(map[int]bool)

	for idx := range nums {
		curIdx := idx
		curResult := 0
		for {
			if _, ok := flags[curIdx]; !ok {
				flags[curIdx] = true
				curResult++
				curIdx = nums[curIdx]
			} else {
				break
			}
		}
		if result < curResult {
			result = curResult
		}
	}

	return result
}

func TestArrayNesting(nums []int) {
	fmt.Printf("Input: A = [%s]\n", strings.Join(util.IntArray2StringArray(nums), ","))
	fmt.Printf("Output: %+v\n", arrayNesting(nums))
}

func main() {
	nums := []int{5, 4, 0, 3, 1, 6, 2}
	TestArrayNesting(nums)
}
