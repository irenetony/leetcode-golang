package main

import (
	"fmt"
	"github.com/berryjam/leetcode-golang/util"
	"strings"
)

func MaxCoins(nums []int) int {
	if nums == nil {
		panic(nums)
	} else if len(nums) == 1 {
		return nums[0]
	} else {
		var res int = -1
		for idx, val := range nums {
			var cur int
			nextNums := make([]int, len(nums)-1)
			if idx == 0 {
				copy(nextNums, nums[1:])
				cur = val*nums[1] + MaxCoins(nextNums)
			} else if idx == len(nums)-1 {
				copy(nextNums, nums[0:len(nums)-1])
				cur = val*nums[len(nums)-2] + MaxCoins(nextNums)
			} else {
				copy(nextNums[0:idx], nums[0:idx])
				copy(nextNums[idx:], nums[idx+1:])
				cur = nums[idx-1]*val*nums[idx+1] + MaxCoins(nextNums)
			}
			if res < cur {
				res = cur
			}
		}
		return res
	}
}

func TestMaxCoins(nums []int) {
	fmt.Printf("Given [%s] Return %d\n", strings.Join(util.IntArray2StringArray(nums), ","), MaxCoins(nums))
}

func main() {
	nums1 := []int{3, 1, 5, 8}
	TestMaxCoins(nums1)

	nums2 := []int{3, 1, 5}
	TestMaxCoins(nums2)
}
