package main

import "fmt"

func findDuplicate(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return -1
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			cur := nums[i]
			next := nums[cur-1]
			if nums[next-1] == next {
				return next
			}
			for cur != i+1 {
				nums[cur-1] = cur
				cur = next
				next = nums[next-1]
				if cur == next {
					break
				}
			}
			nums[i] = cur
		}
	}

	return -1
}

func main() {
	fmt.Printf("res=%d\n", findDuplicate([]int{1, 3, 4, 2, 2}))
	fmt.Printf("res=%d\n", findDuplicate([]int{3, 4, 4, 4, 2}))
	fmt.Printf("res=%d\n", findDuplicate([]int{2, 5, 9, 6, 9, 3, 8, 9, 7, 1}))
	fmt.Printf("res=%d\n", findDuplicate([]int{9, 4, 9, 5, 7, 2, 8, 9, 3, 9}))
	fmt.Printf("res=%d\n", findDuplicate([]int{1, 1, 1, 1, 1}))
}
