package main

import (
	"fmt"
	"github.com/berryjam/leetcode-golang/util"
	"strings"
)

func splitArrayBinarySearch(array []int, num int) int {
	l := 0
	r := len(array) - 1

	for l <= r {
		m := (l + r) / 2
		if array[m] == num {
			return m
		}
		if array[l] == num {
			return l
		}
		if array[r] == num {
			return r
		}
		if array[l] <= array[m] && array[r] >= array[m] {
			if num > array[m] {
				l = m + 1
			} else {
				r = m - 1
			}
		} else {
			if array[l] >= array[m] {
				if num >= array[l] {
					r = m - 1
				} else {
					if num < array[m] {
						r = m - 1
					} else {
						l = m + 1
					}
				}
			} else {
				if num > array[m] {
					l = m + 1
				} else {
					if num > array[r] {
						r = m - 1
					} else {
						l = m + 1
					}
				}
			}
		}
	}

	return -1
}

func TestSplitArrayBinarySearch(array []int, num int) {
	fmt.Printf("Input: [%s]\n", strings.Join(util.IntArray2StringArray(array), ","))
	fmt.Printf("Output: %d\n", splitArrayBinarySearch(array, num))
}

func main() {
	TestSplitArrayBinarySearch([]int{9, 11, 12, 13, 14, 15, 0, 1, 3, 4}, 4)
}
