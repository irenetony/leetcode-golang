package main

import "fmt"

func hammingDistance(x int, y int) int {
	res := 0
	if x == y {
		return res
	} else {
		xbits := int2bits(x)
		ybits := int2bits(y)

		if x < y {
			res = countDiffBits(xbits, ybits)
		} else {
			res = countDiffBits(ybits, xbits)
		}
	}

	return res
}

func int2bits(i int) []int {
	res := make([]int, 0)
	for i > 0 {
		res = append(res, i%2)
		i /= 2
	}
	return res
}

func countDiffBits(lessBits, moreBits []int) int {
	res := 0
	i := 0
	for ; i < len(lessBits); i++ {
		if lessBits[i] != moreBits[i] {
			res++
		}
	}

	for ; i < len(moreBits); i++ {
		if moreBits[i] != 0 {
			res++
		}
	}
	return res
}

func TestHammingDistance(x int, y int) {
	fmt.Printf("Input: x = %d, y = %d\n", x, y)
	fmt.Printf("Output: %d\n", hammingDistance(x, y))
}

func main() {
	TestHammingDistance(1, 4)
}
