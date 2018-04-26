package main

import "fmt"

func findDuplicates(nums []int) []int {
	res := make([]int, 0)

	for i := 0; i < len(nums); {
		if nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		} else {
			i++
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			res = append(res, nums[i])
		}
	}

	return res
}

func main() {
	//fmt.Printf("%s\n", int2Bits(2))
	fmt.Printf("res=%+v", findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}
