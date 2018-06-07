package main

import "fmt"

func numFriendRequests(ages []int) int {
	res := 0
	qs(ages)

	lastAge := -1
	lastAgeId := -1
	for i, age := range ages {
		minAge := age/2 + 8
		minAgeId := bss(ages, minAge)
		if minAgeId != -1 && i >= minAgeId {
			if i == len(ages)-1 || ages[i+1] != age {
				if age != lastAge {
					res += i - minAgeId
				} else {
					res += (lastAgeId - minAgeId) * (i - lastAgeId + 1)
					res += (i - lastAgeId) * (i - lastAgeId + 1)
				}
			}
		}
		if age != lastAge {
			lastAge = age
			lastAgeId = i
		}
	}

	return res
}

func bss(nums []int, num int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		mid := (l + r) / 2
		if num == nums[mid] {
			if mid == 0 || mid > 0 && nums[mid-1] != num {
				return mid
			} else {
				r = mid - 1
			}
		} else if num > nums[mid] {
			if mid == len(nums)-1 {
				return -1
			}
			if mid < len(nums) && num < nums[mid+1] {
				return mid + 1
			}
			l = mid + 1
			if l == r && nums[l] != num {
				return l
			}
		} else if num < nums[mid] {
			if mid > 0 && num > nums[mid-1] || mid == 0 {
				return mid
			}
			r = mid - 1
			if l == r && nums[l] != num {
				return l
			}
		}
	}

	return -1
}

func qs(nums []int) {
	if nums == nil || len(nums) == 0 || len(nums) == 1 {
		return
	}

	pivot := getPivot(nums)
	qs(nums[:pivot])
	qs(nums[pivot+1:])
}

func getPivot(nums []int) int {
	num := nums[0]

	i := 0
	j := len(nums) - 1
	for i < j {
		for i < j {
			if nums[j] < num {
				nums[i] = nums[j]
				break
			} else {
				j--
			}
		}

		for i < j {
			if nums[i] > num {
				nums[j] = nums[i]
				break
			} else {
				i++
			}
		}

	}
	nums[i] = num
	return i
}

func main() {
	//nums := []int{9, 8, 11, 3, 6, 17, 15}
	//qs(nums)
	//fmt.Printf("%+v\n", nums)
	fmt.Printf("res=%d\n", numFriendRequests([]int{16, 16}))
	fmt.Printf("res=%d\n", numFriendRequests([]int{16, 17, 18}))
	fmt.Printf("res=%d\n", numFriendRequests([]int{20, 30, 100, 110, 120}))
	fmt.Printf("res=%d\n", numFriendRequests([]int{30, 29, 79, 119, 70}))
	fmt.Printf("res=%d\n", numFriendRequests([]int{8, 85, 24, 85, 69}))
}
