package main

import "fmt"

const MinInt32 = -1 << 31

func maxRotateFunction(A []int) int {
	if A == nil || len(A) == 0 {
		return 0
	}

	res := MinInt32

	for i := 0; i < len(A); i++ {
		tmp := rotateFunction(A, i)
		if tmp > res {
			res = tmp
		}
	}

	return res
}

func rotateFunction(A []int, k int) int {
	res := 0
	for i := 0; i < len(A); i++ {
		res += i * A[((len(A) - k + i) % len(A))]
	}
	return res
}

func main() {
	fmt.Printf("%d", maxRotateFunction([]int{4,3,2,6}))
}
