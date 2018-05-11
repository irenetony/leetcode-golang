package main

import "fmt"

func unorderArrayMedian(nums []int) float64 {
	if len(nums)%2 == 1 {
		m := (len(nums) + 1) / 2
		minHeap := make([]int, m)
		for i, v := range nums {
			if i < m-1 {
				minHeap[i] = v
			} else if i == m-1 {
				minHeap[i] = v
				initMinHeap(minHeap)
			} else {
				if v > minHeap[0] {
					minHeap[0] = v
					heapifyMinHeap(minHeap)
				}
			}
		}
		return float64(minHeap[0])
	} else {
		m := len(nums) / 2
		minHeap := make([]int, m)
		maxHeap := make([]int, m)

		for i, v := range nums {
			if i < m-1 {
				minHeap[i] = v
			} else if i == m-1 {
				minHeap[i] = v
				initMinHeap(minHeap)
			} else {
				if v > minHeap[0] {
					minHeap[0] = v
					heapifyMinHeap(minHeap)
				}
			}
		}

		for i, v := range nums {
			if i < m-1 {
				maxHeap[i] = v
			} else if i == m-1 {
				maxHeap[i] = v
				initMaxHeap(maxHeap)
			} else {
				if v < maxHeap[0] {
					maxHeap[0] = v
					heapifyMaxHeap(maxHeap)
				}
			}
		}

		return (float64(minHeap[0]) + float64(maxHeap[0])) / 2
	}
}

func initMinHeap(nums []int) {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		cur := i
		for cur <= len(nums)/2 {
			if 2*cur+2 < len(nums) { // has right child
				if nums[2*cur+1] < nums[2*cur+2] && nums[cur] > nums[2*cur+1] {
					nums[2*cur+1], nums[cur] = nums[cur], nums[2*cur+1]
					cur = 2*cur + 1
				} else if nums[2*cur+1] > nums[2*cur+2] && nums[cur] > nums[2*cur+2] {
					nums[2*cur+2], nums[cur] = nums[cur], nums[2*cur+2]
					cur = 2*cur + 2
				} else {
					break
				}
			} else { // has only left child
				if nums[2*cur+1] < nums[cur] {
					nums[2*cur+1], nums[cur] = nums[cur], nums[2*cur+1]
					cur = 2*cur + 1
				} else {
					break
				}
			}
		}
	}
}

func initMaxHeap(nums []int) {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		cur := i
		for cur <= len(nums)/2-1 {
			if 2*cur+2 < len(nums) { // has right child
				if nums[2*cur+1] > nums[2*cur+2] && nums[cur] < nums[2*cur+1] {
					nums[2*cur+1], nums[cur] = nums[cur], nums[2*cur+1]
					cur = 2*cur + 1
				} else if nums[2*cur+1] < nums[2*cur+2] && nums[cur] < nums[2*cur+2] {
					nums[2*cur+2], nums[cur] = nums[cur], nums[2*cur+2]
					cur = 2*cur + 2
				} else {
					break
				}
			} else { // has only left child
				if nums[2*cur+1] > nums[cur] {
					nums[2*cur+1], nums[cur] = nums[cur], nums[2*cur+1]
					cur = 2*cur + 1
				} else {
					break
				}
			}
		}
	}
}

func heapifyMinHeap(nums []int) {
	for cur := 0; cur <= len(nums)/2-1; {
		if 2*cur+2 < len(nums) { // has right child
			if nums[2*cur+1] < nums[2*cur+2] && nums[cur] > nums[2*cur+1] {
				nums[2*cur+1], nums[cur] = nums[cur], nums[2*cur+1]
				cur = 2*cur + 1
			} else if nums[2*cur+1] > nums[2*cur+2] && nums[cur] > nums[2*cur+2] {
				nums[2*cur+2], nums[cur] = nums[cur], nums[2*cur+2]
				cur = 2*cur + 2
			} else {
				break
			}
		} else { // has only left child
			if nums[2*cur+1] < nums[cur] {
				nums[2*cur+1], nums[cur] = nums[cur], nums[2*cur+1]
				cur = 2*cur + 1
			} else {
				break
			}
		}
	}
}

func heapifyMaxHeap(nums []int) {
	for cur := 0; cur <= len(nums)/2-1; {
		if 2*cur+2 < len(nums) { // has right child
			if nums[2*cur+1] > nums[2*cur+2] && nums[cur] < nums[2*cur+1] {
				nums[2*cur+1], nums[cur] = nums[cur], nums[2*cur+1]
				cur = 2*cur + 1
			} else if nums[2*cur+1] < nums[2*cur+2] && nums[cur] < nums[2*cur+2] {
				nums[2*cur+2], nums[cur] = nums[cur], nums[2*cur+2]
				cur = 2*cur + 2
			} else {
				break
			}
		} else { // has only left child
			if nums[2*cur+1] > nums[cur] {
				nums[2*cur+1], nums[cur] = nums[cur], nums[2*cur+1]
				cur = 2*cur + 1
			} else {
				break
			}
		}
	}
}

func main() {
	fmt.Printf("res=%+v\n", unorderArrayMedian([]int{1, 4, 9, 8, 5, 10, 11, 13, 19}))
	fmt.Printf("res=%+v\n", unorderArrayMedian([]int{1, 4, 2, 3, 6, 5, 8, 7}))
}
