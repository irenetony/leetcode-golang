package main

import "fmt"

func majorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	res := nums[0]
	curCount := 1
	for i := 1; i < len(nums); i++ {
		if res != nums[i] {
			curCount--
			if curCount == 0 {
				res = nums[i]
				curCount = 1
			}
		} else {
			curCount++
		}
	}

	return res
}

func main() {
	fmt.Printf("%d", majorityElement([]int{8, 8, 7, 7, 7}))
}
