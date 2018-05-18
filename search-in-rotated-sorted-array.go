package main

import "fmt"

func search(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}

	l := 0
	r := len(nums) - 1

	min := nums[(l+r)/2]
	midIdx := (l + r) / 2

	for l <= r {
		mid := (l + r) / 2
		if min > nums[mid] {
			min = nums[mid]
			midIdx = mid
		}
		if nums[mid] < nums[len(nums)-1] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	if midIdx == 0 {
		l = 0
		r = len(nums) - 1

	} else {
		if target >= nums[0] && target <= nums[midIdx-1] {
			l = 0
			r = midIdx - 1
		}
		if target >= nums[midIdx] && target <= nums[len(nums)-1] {
			l = midIdx
			r = len(nums) - 1
		}
	}

	return binarySearch(nums, l, r, target)
}

func binarySearch(nums []int, l, r, target int) int {
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}

func main() {
	fmt.Printf("res=%d\n", search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	fmt.Printf("res=%d\n", search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
}
