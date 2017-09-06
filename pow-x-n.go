package main

import (
	"fmt"
)

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	} else if n > 0 {
		binaryArray := dex2binary(n)
		res := binaryArray2myPow(binaryArray, x)
		//fmt.Printf("%s\n", strings.Join(util.IntArray2StringArray(binaryArray), ","))
		return res
	} else {
		binaryArray := dex2binary(-n)
		res := binaryArray2myPow(binaryArray, x)
		return 1 / res
	}
}

func dex2binary(n int) []int {
	res := make([]int, 0)

	for n > 0 {
		bit := n % 2
		res = append(res, bit)
		n /= 2
	}

	return res
}

func binaryArray2myPow(binaryArray []int, x float64) float64 {
	curPow := x
	res := 1.0
	for i := 0; i < len(binaryArray); i++ {
		if binaryArray[i] == 1 {
			res *= curPow
		}
		curPow *= curPow
	}

	return res
}

func TestMyPow(x float64, n int) {
	fmt.Printf("pow(%f,%d)=%f\n", x, n, myPow(x, n))
}

func main() {
	TestMyPow(3.0, 3)
	TestMyPow(2, -1)
}
