package main

import (
	"fmt"
	"github.com/berryjam/leetcode-golang/util"
	"strings"
)

type NumCountPair struct {
	Num   int
	Count int
}

func topKFrequent(nums []int, k int) []int {
	hashMap := make(map[int]int)
	for _, v := range nums {
		if _, ok := hashMap[v]; ok {
			hashMap[v]++
		} else {
			hashMap[v] = 1
		}
	}

	pairs := make([]NumCountPair, 0)
	//countIntsMap := make(map[int][]int)
	for num, count := range hashMap {
		pairs = append(pairs, NumCountPair{Num: num, Count: count})
	}
	fmt.Printf("%+v\n", pairs)
	minHeap := make([]NumCountPair, 0)
	for _, pair := range pairs {
		if len(minHeap) < k {
			minHeap = append(minHeap, pair)
			if len(minHeap) == k {
				for idx := k/2 - 1; idx >= 0; idx-- {
					heapifyHeap(minHeap, idx)
				}
			}
		} else {
			if pair.Count > minHeap[0].Count {
				minHeap[0] = pair
				heapifyHeap(minHeap, 0)
			}
		}
	}
	res := make([]int, 0)
	for _, pair := range minHeap {
		res = append(res, pair.Num)
	}
	return res
}

func heapifyHeap(minHeap []NumCountPair, idx int) {
	curIdx := idx
	for curIdx <= len(minHeap)/2-1 {
		if 2*curIdx+2 <= len(minHeap)-1 { // has right child
			if minHeap[2*curIdx+1].Count < minHeap[2*curIdx+2].Count { // left is smaller
				if minHeap[curIdx].Count > minHeap[2*curIdx+1].Count {
					minHeap[2*curIdx+1], minHeap[curIdx] = minHeap[curIdx], minHeap[2*curIdx+1]
					curIdx = 2*curIdx + 1
				} else {
					break
				}
			} else { // right is smaller
				if minHeap[curIdx].Count > minHeap[2*curIdx+2].Count {
					minHeap[2*curIdx+2], minHeap[curIdx] = minHeap[curIdx], minHeap[2*curIdx+2]
					curIdx = 2*curIdx + 2
				} else {
					break
				}
			}
		} else { // has only left child
			if minHeap[2*curIdx+1].Count < minHeap[curIdx].Count { // left child is smaller
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
