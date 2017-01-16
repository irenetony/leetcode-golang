package main

import "fmt"

func reverse(x int) int {
	result := 0
	if x == 0 {
		return 0
	} else if x > 0 {
		for i := x / 10; i != 0; i = i / 10 {
			result = result * 10 + x % 10;
			x = x / 10
		}
		result = result * 10 + x % 10
	} else {
		x = -x
		for i := x / 10; i != 0; i = i / 10 {
			result = result * 10 + x % 10;
			x = x / 10
		}
		result = result * 10 + x % 10
		result = - result
	}
	return result
}

func main() {
	x := 123
	fmt.Printf("reverse %d:%d\n", x, reverse(x))

	y := -123
	fmt.Printf("reverse %d:%d\n", y, reverse(y))

	z := 0
	fmt.Printf("reverse %d:%d\n", z, reverse(z))

	i := -9812371232
	fmt.Printf("reverse %d:%d\n", i, reverse(i))
}