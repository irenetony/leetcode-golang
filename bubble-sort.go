package main

import (
	"github.com/berryjam/leetcode/util"
	"fmt"
	"strings"
)

func BubbleSort(data []int) {
	if data == nil {
		return
	}

	for i := len(data) - 2; i >= 0; i-- {
		hasSawped := false
		for j := 0; j <= i; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
				hasSawped = true
			}
		}
		if !hasSawped {
			break
		}
	}
}

func main() {
	//data := []int{3, 7, 8, 5, 4, 10, 15, 1, 2, 9}
	//data := []int{}
	//data := []int{3}
	//data := []int{1, 2}
	data := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	BubbleSort(data)
	tmpData := util.IntArray2StringArray(data)
	fmt.Printf("Sorted data : %s", strings.Join(tmpData, ","))
}
