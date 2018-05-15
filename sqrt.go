package main

import "fmt"

func mySqrt(x int) int {
	l := 0
	r := x
	for l <= r {
		mid := (l + r) / 2
		if mid*mid == x {
			return mid
		} else if mid*mid > x {
			r = mid - 1
		} else if mid*mid < x {
			if (mid+1)*(mid+1) == x {
				return mid + 1
			}
			if (mid+1)*(mid+1) > x {
				return mid
			}
			l = mid + 1
		}
	}
	return 0
}

func main() {
	fmt.Printf("res=%d\n", mySqrt(4))
	fmt.Printf("res=%d\n", mySqrt(8))
}
