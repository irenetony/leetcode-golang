package main

import (
	"fmt"
	"github.com/berryjam/leetcode-golang/datastructure"
)

func dailyTemperatures(temperatures []int) []int {
	if temperatures == nil || len(temperatures) == 0 {
		return []int{}
	}
	res := make([]int, len(temperatures))
	stack := datastructure.NewStack()
	for i := len(temperatures) - 1; i >= 0; i-- {
		for stack.Len() > 0 {
			top := stack.Top()
			topVal := top.(int)
			if temperatures[i] >= temperatures[topVal] {
				stack.Pop()
			} else {
				res[i] = topVal - i
				break
			}
		}
		if stack.Len() == 0 {
			res[i] = 0
		}
		stack.Push(i)
	}
	return res
}

func main() {
	temps1 := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Printf("res = %d \n", dailyTemperatures(temps1))
	temps2 := []int{55, 38, 53, 81, 61, 93, 97, 32, 43, 78}
	fmt.Printf("res = %d \n", dailyTemperatures(temps2))
	temps3 := []int{34, 80, 80, 34, 34, 80, 80, 80, 80, 34}
	fmt.Printf("res = %d \n", dailyTemperatures(temps3))
}
