package main

import (
	"fmt"
	"github.com/berryjam/leetcode-golang/util"
	"strings"
)

func topKFrequent(nums []int, k int) []int {
	hashMap := make(map[int]int)
	for _, v := range nums {
		if _, ok := hashMap[v]; ok {
			hashMap[v]++
		} else {
			hashMap[v] = 1
		}
	}

	countIntsMap := make(map[int][]int)
	for key, val := range hashMap {
		if _, ok := countIntsMap[val]; ok {
			countIntsMap[val] = append(countIntsMap[val], key)
		} else {
			countIntsMap[val] = []int{key}
		}
	}
	fmt.Printf("%+v\n", countIntsMap)
	minHeap := make([]int, 0)
	for key := range countIntsMap {
		if len(minHeap) < k {
			minHeap = append(minHeap, key)
			if len(minHeap) == k {
				for idx := k/2 - 1; idx >= 0; idx-- {
					heapifyHeap(minHeap, idx)
				}
			}
		} else {
			if key > minHeap[0] {
				minHeap[0] = key
				heapifyHeap(minHeap, 0)
			}
		}
	}
	res := make([]int, 0)
	for _, count := range minHeap {
		res = append(res, countIntsMap[count]...)
	}
	return res
}

func heapifyHeap(minHeap []int, idx int) {
	curIdx := idx
	for curIdx <= len(minHeap)/2-1 {
		if 2*curIdx+2 <= len(minHeap)-1 { // has right child
			if minHeap[2*curIdx+1] < minHeap[2*curIdx+2] { // left is smaller
				if minHeap[curIdx] > minHeap[2*curIdx+1] {
					minHeap[2*curIdx+1], minHeap[curIdx] = minHeap[curIdx], minHeap[2*curIdx+1]
					curIdx = 2*curIdx + 1
				} else {
					break
				}
			} else { // right is smaller
				if minHeap[curIdx] > minHeap[2*curIdx+2] {
					minHeap[2*curIdx+2], minHeap[curIdx] = minHeap[curIdx], minHeap[2*curIdx+2]
					curIdx = 2*curIdx + 2
				} else {
					break
				}
			}
		} else { // has only left child
			if minHeap[2*curIdx+1] < minHeap[curIdx] { // left child is smaller
				minHeap[2*curIdx+1], minHeap[curIdx] = minHeap[curIdx], minHeap[2*curIdx+1]
				curIdx = 2*curIdx + 1
			} else {
				break
			}
		}
	}
}

func TestTopKFrequent(nums []int, k int) []int {
	return topKFrequent(nums, k)
}

func main() {
	res := TestTopKFrequent([]int{1, 1, 1, 2, 2, 3, 4, 4, 4, 5, 5, 5, 5, 5, 6, 7, 7, 7, 7, 7, 7, 7}, 3)
	fmt.Printf("res=[%s]", strings.Join(util.IntArray2StringArray(res), ","))
}
