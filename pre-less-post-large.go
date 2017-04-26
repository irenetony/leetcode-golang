package main

import (
	"fmt"
	"leetcode/util"
	"strings"
)

func preLessPostLarge(array []int) []int {
	if array == nil || len(array) == 0 || len(array) == 1 {
		return array
	}

	res := make([]int, 0)
	flags := make([]bool, len(array))

	leftMax := array[0]
	for i := 0; i < len(array); i++ {
		if array[i] > leftMax {
			leftMax = array[i]
		} else {
			if i == 0 {
				continue
			}
			flags[i] = true
		}
	}
	if array[0] > leftMax {
		flags[0] = true
	}

	rightMin := array[len(array)-1]
	for i := len(array) - 1; i >= 0; i-- {
		if array[i] < rightMin {
			rightMin = array[i]
		} else {
			if i == len(array)-1 {
				continue
			}
			flags[i] = true
		}
	}
	if array[len(array)-1] < rightMin {
		flags[len(array)-1] = true
	}

	for idx, excluded := range flags {
		if !excluded {
			res = append(res, array[idx])
		}
	}

	return res
}

func main() {
	array := []int{1, 6, 4, 5, 9, 10 }
	res := preLessPostLarge(array)
	fmt.Printf("%s", strings.Join(util.IntArray2StringArray(res), ","))
}
