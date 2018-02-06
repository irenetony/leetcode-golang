package main

import (
	"fmt"
	"strings"
	"leetcode/util"
)

func arrayNesting(nums []int) int {
	// FIXME
	return 6
}

func TestArrayNesting(nums []int) {
	fmt.Printf("Input: A = [%s]\n", strings.Join(util.IntArray2StringArray(nums), ","))
	fmt.Printf("Output: %+v\n", arrayNesting(nums))
}

func main() {
	nums := []int{5, 4, 0, 3, 1, 6, 2}
	TestArrayNesting(nums)
}
