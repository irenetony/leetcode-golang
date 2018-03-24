package main

import (
	"fmt"
	"github.com/berryjam/leetcode/util"
	"strings"
)

func LeastBricks(wall [][]int) int {
	res := len(wall)

	totalWidth := 0
	for _, val := range wall[0] {
		totalWidth += val
	}

	for pos := 1; pos <= totalWidth-1; pos++ {
		crossedBricks := 0
		for _, bricks := range wall {
			if isCrossBrick(bricks, pos) {
				crossedBricks++
			}
		}
		if crossedBricks < res {
			res = crossedBricks
		}
	}

	return res
}

func isCrossBrick(bricks []int, pos int) bool {
	tmp := 0
	for _, val := range bricks {
		tmp += val
		if tmp == pos {
			return false
		} else if tmp > pos {
			return true
		}
	}
	return false
}

func TestLeastBricks(wall [][]int) {
	res := LeastBricks(wall)
	fmt.Printf("Input:\n")
	fmt.Printf(twoDimensionArray2string(wall) + "\n")
	fmt.Printf("Output: %d", res)
}

func twoDimensionArray2string(array [][]int) string {
	res := "["
	content := make([]string, 0)
	for _, val := range array {
		content = append(content, oneDimensionArray2string(val))
	}
	for idx, val := range content {
		if idx == len(content)-1 {
			res += val
		} else {
			res += val + ",\n"
		}
	}
	res += "]"
	return res
}

func oneDimensionArray2string(array []int) string {
	res := "["
	tmp := util.IntArray2StringArray(array)
	res += strings.Join(tmp, ",")
	res += "]"
	return res
}

func main() {
	wall1 := [][]int{
		{1, 2, 1, 2},
		{3, 1, 2},
		{1, 3, 2},
		{2, 2, 2},
		{3, 1, 2},
		{1, 3, 1, 1},
	}
	TestLeastBricks(wall1)
}
