package main

import "fmt"

func searchRange(nums []int, target int) []int {
	if nums == nil || len(nums) == 0 {
		return []int{-1, -1}
	}

	l := 0
	r := len(nums) - 1

	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			fmt.Printf("first set idx :%d is %d\n", target, mid)
			fmt.Printf("start point :%d\n", findStartPoint(nums, l, mid, r, target))
			fmt.Printf("end point :%d\n", findEndPoint(nums, l, mid, r, target))

			return []int{findStartPoint(nums, l, mid, r, target), findEndPoint(nums, l, mid, r, target)}
		} else if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		}
	}

	return []int{-1, -1}
}

func findStartPoint(nums []int, l, mid, r, target int) int {

	for mid >= 1 {
		if nums[mid-1] != target {
			break
		} else {
			r = mid - 1
			mid = bs(nums, &l, &r, target)
		}
	}

	return mid
}

func findEndPoint(nums []int, l, mid, r, target int) int {

	for mid < len(nums)-1 {
		if nums[mid+1] != target {
			break
		} else {
			l = mid + 1
			mid = bs(nums, &l, &r, target)
		}
	}

	return mid
}

func bs(nums []int, l *int, r *int, target int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}

	for *l <= *r {
		mid := (*l + *r) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			*l = mid + 1
		} else if nums[mid] > target {
			*r = mid - 1
		}
	}

	return -1
}

func main() {
	searchRange([]int{1, 1, 2, 4, 5, 5, 5, 6, 7, 7, 7, 7, 8, 9, 11, 13}, 5)
	searchRange([]int{2, 2}, 2)
}
