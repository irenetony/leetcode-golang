package main

import (
	"fmt"
	"leetcode/util"
	"strings"
)

func heapSort(array []int) {
	initHeap(array)
	length := len(array) - 1
	for i := len(array) - 1; i >= 1; i-- {
		array[0], array[i] = array[i], array[0]
		heapify(array, 0, length)
		length--
	}
}

func initHeap(array []int) {
	lastNonLeafIdx := len(array)/2 - 1
	length := len(array)
	for i := lastNonLeafIdx; i >= 0; i-- {
		heapify(array, i, length)
	}
}

func heapify(array []int, i, limit int) {
	curIdx := i
	for curIdx*2+2 < len(array) {
		maxIdx := 2*curIdx + 1
		if maxIdx >= limit {
			break
		}
		maxChildVal := array[maxIdx]
		if 2*curIdx+2 < len(array) && array[2*curIdx+2] > maxChildVal && 2*curIdx+2 < limit {
			maxIdx = 2*curIdx + 2
			maxChildVal = array[maxIdx]
		}
		if array[curIdx] < maxChildVal {
			array[curIdx], array[maxIdx] = array[maxIdx], array[curIdx]
			curIdx = maxIdx
		} else {
			break
		}
	}
}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7}
	heapSort(array)
	fmt.Printf("%s", strings.Join(util.IntArray2StringArray(array), ","))
}
