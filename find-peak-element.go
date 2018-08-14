package main

import "fmt"

func findPeakElement(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}

	if len(nums) == 1 {
		return 0
	}

	return findPeakElementRecur(nums, 0, len(nums)-1)
}

func findPeakElementRecur(nums []int, l, r int) int {
	for l <= r {
		mid := (l + r) / 2
		if mid-1 >= 0 && mid+1 <= len(nums)-1 && nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
			return mid
		}
		if (mid == 0 && nums[mid] > nums[mid+1]) || (mid == len(nums)-1 && nums[mid] > nums[mid-1]) {
			return mid
		}
		if nums[mid] <= nums[l] && mid != l {
			r = mid - 1
			continue
		}
		if nums[mid] <= nums[r] && mid != r {
			l = mid + 1
			continue
		}

		leftPeakPos := findPeakElementRecur(nums, l, mid-1)
		if leftPeakPos != -1 {
			return leftPeakPos
		}

		rightPeakPos := findPeakElementRecur(nums, mid+1, r)
		if rightPeakPos != - 1 {
			return rightPeakPos
		}

		return -1
	}

	return -1
}

func main() {
	arr1 := []int{1, 2, 3, 1}
	arr2 := []int{1, 2, 1, 3, 5, 6, 4}
	arr3 := []int{1}
	arr4 := []int{1, 2}
	fmt.Printf("res=%d\n", findPeakElement(arr1))
	fmt.Printf("res=%d\n", findPeakElement(arr2))
	fmt.Printf("res=%d\n", findPeakElement(arr3))
	fmt.Printf("res=%d\n", findPeakElement(arr4))
}
