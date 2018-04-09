package main

import (
	"fmt"
	"github.com/berryjam/leetcode-golang/util"
	"strings"
)

func reconstructQueue(people [][]int) [][]int {
	// FIXME
	quickSortEntry(people)
	fmt.Printf("%+v\n", people)
	return [][]int{{5, 0}, {7, 0}, {5, 2}, {6, 1}, {4, 4}, {7, 1}}
}

func quickSortEntry(array [][]int) {
	quickSortDoubleArray(array, 0, len(array)-1)
}

func quickSortDoubleArray(array [][]int, l, r int) {
	if l < 0 || r > len(array)-1 || l >= r {
		return
	}

	p := array[l]
	i := l
	j := r
	pid := j
	for i < j {
		for {
			if j > i {
				if array[j][0] < p[0] || array[j][0] == p[0] && array[j][1] < p[1] {
					array[i] = array[j]
					break
				} else {
					j--
				}
			} else {
				pid = j
				break
			}
		}

		for {
			if j > i {
				if array[i][0] > p[0] || array[i][0] == p[0] && array[i][1] > p[1] {
					array[j] = array[i]
					break
				} else {
					i++
				}
			} else {
				pid = i
				break
			}
		}
	}

	array[pid] = p
	quickSortDoubleArray(array, l, pid-1)
	quickSortDoubleArray(array, pid+1, r)
}

func TestReconstructQueue(people [][]int) {
	fmt.Println("Input:")
	fmt.Printf("%s", IntPair2String(people))

	fmt.Println("Output:")
	fmt.Printf("%s", IntPair2String(reconstructQueue(people)))
}

func IntPair2String(pair [][]int) string {
	inputArray := make([]string, 0)
	for _, hk := range pair {
		inputArray = append(inputArray, fmt.Sprintf("[%s]", strings.Join(util.IntArray2StringArray(hk), ",")))
	}
	return fmt.Sprintf("[%s]\n\n", strings.Join(inputArray, ","))
}

func main() {
	TestReconstructQueue([][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}})
}
