package main

import (
	"fmt"
	"leetcode/util"
	"strings"
)

func IsSelfCrossing(x []int) bool {
	for idx := 3; idx < len(x); idx++ {
		if x[idx] >= x[idx-2] && x[idx-3] >= x[idx-1] {
			return true
		}
		if idx >= 4 && x[idx-1] == x[idx-3] && x[idx] >= x[idx-2]-x[idx-4] {
			return true
		}
		if idx >= 5 && x[idx-2] >= x[idx-4] && x[idx-3] >= x[idx-1] && x[idx-1] >= x[idx-3]-x[idx-5] &&
			x[idx] >= x[idx-2]-x[idx-4] {
			return true
		}

	}
	return false
}

func TestIsSelfCrossing1() {
	x := []int{2, 1, 1, 2}
	fmt.Printf("Given x = [%s]\nReturn %t (self crossing)\n",
		strings.Join(util.IntArray2StringArray(x), ","), IsSelfCrossing(x))
}

func TestIsSelfCrossing2() {
	x := []int{1, 2, 3, 4}
	fmt.Printf("Given x = [%s]\nReturn %t (self crossing)\n",
		strings.Join(util.IntArray2StringArray(x), ","), IsSelfCrossing(x))
}

func TestIsSelfCrossing3() {
	x := []int{1, 1, 1, 1}
	fmt.Printf("Given x = [%s]\nReturn %t (self crossing)\n",
		strings.Join(util.IntArray2StringArray(x), ","), IsSelfCrossing(x))
}

func main() {
	TestIsSelfCrossing1()
	TestIsSelfCrossing2()
	TestIsSelfCrossing3()
}
