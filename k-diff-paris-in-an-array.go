package main

import (
	"fmt"
	"github.com/berryjam/leetcode-golang/util"
	"strings"
	"sort"
)

func findPairs(nums []int, k int) int {
	res := 0
	samePairs := make(map[int]bool)
	pairs := make(map[int]bool)
	sort.Ints(nums)
	if k == 0 {
		for i := 0; i < len(nums); i++ {
			if _, ok1 := samePairs[nums[i]]; !ok1 {
				if _, ok2 := pairs[nums[i]]; ok2 {
					res++
					samePairs[nums[i]] = true
				}
			}
			pairs[nums[i]] = true
		}
	} else {
		for i := 0; i < len(nums); i++ {
			if _, ok1 := pairs[nums[i]]; !ok1 {
				if _, ok2 := pairs[nums[i]-k]; ok2 {
					res++
				}
			}
			pairs[nums[i]] = true
		}
	}
	return res
}

func TestFindPairs(nums []int, k int) {
	res := findPairs(nums, k)
	fmt.Printf("Input:[%s], k = %d\n", strings.Join(util.IntArray2StringArray(nums), ","), k)
	fmt.Printf("Output: %d\n", res)
}

func main() {
	//nums := []int{3, 1, 4, 1, 5}
	//k := 2

	//nums := []int{1, 2, 3, 4, 5}
	//k := 1

	nums := []int{2, 9, 0, 8, 9, 6, 5, 9, 8, 1, 9, 6, 9, 2, 8, 8, 7, 5, 7, 8, 8, 3, 7, 4, 1, 1, 6, 2, 9, 9, 3, 9, 2, 4, 6, 5, 6, 5, 1, 5, 9, 9, 8, 1, 4, 3, 2, 8, 5, 3, 5, 4, 5, 7, 0, 0, 7, 6, 4, 7, 2, 4, 9, 3, 6, 6, 5, 0, 0, 0, 8, 9, 9, 6, 5, 4, 6, 2, 1, 3, 2, 5, 0, 1, 4, 2, 6, 9, 5, 4, 9, 6, 0, 8, 3, 8, 0, 0, 2, 1, 1}
	k := 1
	TestFindPairs(nums, k)
}
