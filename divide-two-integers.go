package main

import "fmt"

func divide(dividend int, divisor int) int {
	if dividend == -1<<31 && divisor == -1 {
		return 1<<31 - 1
	}
	if dividend == 0 {
		return 0
	} else if divisor == 0 {
		panic("divisor can not be zero!")
	} else {
		res := 0
		tmpDividend := dividend
		tmpDivisor := divisor
		if tmpDividend < 0 {
			tmpDividend = -tmpDividend
		}
		if tmpDivisor < 0 {
			tmpDivisor = -tmpDivisor
		}

		for tmpDividend >= tmpDivisor {
			if tmpDividend == tmpDivisor {
				res++
				break
			}
			tmp := tmpDivisor
			var i uint32 = 1
			for ; tmp < tmpDividend; i++ {
				tmp = tmp << 1
				if tmp > tmpDividend {
					shiftCount := 1 << (i - 1)
					res += shiftCount
					tmpDividend -= tmp >> 1
					break
				}
			}
		}

		if dividend > 0 && divisor < 0 || dividend < 0 && divisor > 0 {
			res = -res
		}

		return res
	}
}

func main() {
	fmt.Printf("res=%d\n", divide(-(1 << 31), -1))
}
