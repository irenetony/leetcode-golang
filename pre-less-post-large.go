package main

import (
	"fmt"
	"github.com/berryjam/leetcode-golang/util"
	"strings"
)

func preLessPostLarge(array []int) []int {
	if array == nil || len(array) == 0 || len(array) == 1 {
		return array
	}

	res := make([]int, 0)
	flags := make([]bool, len(array))

	max := array[0]
	min := array[0]

	rightMin := array[len(array)-1]
	for i := len(array) - 1; i >= 0; i-- {
		if array[i] < rightMin {
			rightMin = array[i]
			min = array[i]
		} else {
			max = array[i]
			if i == len(array)-1 {
				continue
			}
			flags[i] = true
		}
	}
	if array[0] > min {
		flags[0] = true
	}
	if array[len(array)-1] < max {
		flags[len(array)-1] = true
	}

	leftMax := array[0]
	for i := 0; i < len(array); i++ {
		if array[i] > leftMax {
			leftMax = array[i]
		} else {
			if i != 0 {
				flags[i] = true
			}
		}
		if !flags[i] {
			res = append(res, array[i])
		}
	}

	return res
}

func main() {
	array := []int{1, 6, 4, 5, 7, 9, 10 }
	res := preLessPostLarge(array)
	fmt.Printf("%s", strings.Join(util.IntArray2StringArray(res), ","))
}
