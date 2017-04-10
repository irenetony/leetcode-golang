package main

import (
	"fmt"
	"leetcode/util"
	"strings"
)

func QuickSort(array []int) {
	if array == nil {
		return
	}

	quickSort(array, 0, len(array)-1)
}

func quickSort(array []int, left, right int) {
	if left >= right {
		return
	}

	pivot := array[left]

	i, j := left, right

	for j > i {

		for j >= 0 && j > i {
			if array[j] < pivot {
				array[i] = array[j]
				i++
				break
			} else {
				j--
			}
		}

		for i >= 0 && i < j {
			if array[i] > pivot {
				array[j] = array[i]
				j--
				break
			} else {
				i++
			}
		}
	}

	array[j] = pivot
	quickSort(array, left, j-1)
	quickSort(array, j+1, right)
}

func main() {
	//data := []int{3, 7, 8, 5, 4, 10, 15, 1, 2, 9}
	//data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//data := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	//data := []int{}
	var data []int = nil
	QuickSort(data)
	tmpData := util.IntArray2StringArray(data)
	fmt.Printf("Sorted data : %s", strings.Join(tmpData, ","))
}
