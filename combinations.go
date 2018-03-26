package main

import (
	"fmt"
	"github.com/berryjam/leetcode-golang/util"
	"strings"
)

func Combine(n, k int) [][]int {
	var array []int
	for i := 1; i <= n; i++ {
		array = append(array, i)
	}
	return CombineRecur(array, k)
}

func CombineRecur(array []int, count int) [][]int {
	switch count {
	case 0:
		return nil
	case 1:
		var res [][]int
		for _, val := range array {
			res = append(res, []int{val})
		}
		return res
	default:
		var res [][]int
		for i, val := range array {
			nextArray := DeleteElement(array, i)
			nextCombination := CombineRecur(nextArray, count-1)
			if nextCombination != nil {
				for _, v := range nextCombination {
					res = append(res, append(CopyArray(v), val))
				}
			}
		}
		return res
	}
}

func DeleteElement(array []int, idx int) []int {
	var res []int
	for i := idx + 1; i < len(array); i++ {
		res = append(res, array[i])
	}
	return res
}

func CopyArray(array []int) []int {
	var res []int
	for _, val := range array {
		res = append(res, val)
	}
	return res
}

func TestCombine(n, k int) {
	array := Combine(n, k)
	fmt.Printf("[\n")
	for _, val := range array {
		fmt.Printf("\t[%s]\n", strings.Join(util.IntArray2StringArray(RevertArray(val)), ","))
	}
	fmt.Printf("\n]\n")
}

func RevertArray(array []int) []int {
	var res []int
	for i := len(array) - 1; i >= 0; i-- {
		res = append(res, array[i])
	}
	return res
}

func main() {
	TestCombine(4, 2)
	TestCombine(4, 3)
}
