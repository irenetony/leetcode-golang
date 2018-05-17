package main

func findMin(nums []int) int {
	l := 0
	r := len(nums)-1

	min := nums[(l+r)/2]

	for l <= r {
		mid := (l+r)/2
		if min > nums[mid] {
			min = nums[mid]
		}
		if nums[mid] < nums[len(nums)-1] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return min
}

func main() {
}
