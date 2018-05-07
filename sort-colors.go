package main

import (
	"fmt"
	"github.com/berryjam/leetcode-golang/util"
	"strings"
)

func sortColors(nums []int) {
	zeroCount := 0
	oneCount := 0
	for _, v := range nums {
		if v == 0 {
			zeroCount++
		}
		if v == 1 {
			oneCount++
		}
	}

	for i := range nums {
		if zeroCount > 0 {
			nums[i] = 0
			zeroCount--
		} else if oneCount > 0 {
			nums[i] = 1
			oneCount--
		} else {
			nums[i] = 2
		}

	}
}

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Printf("res=[%s]\n", strings.Join(util.IntArray2StringArray(nums), ","))
}
