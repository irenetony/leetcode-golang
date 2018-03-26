package main

import (
	"fmt"
	"github.com/berryjam/leetcode-golang/util"
	"strings"
	"math/rand"
	"time"
)

type Solution struct {
	originalNums []int
	output       []int
}

func Constructor(nums []int) Solution {
	tmp := make([]int, len(nums))
	copy(tmp, nums)
	return Solution{originalNums: nums, output: tmp}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	tmp := make([]int, len(this.originalNums))
	copy(tmp, this.originalNums)
	this.output = tmp
	return this.output
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range this.output {
		id := r.Intn(len(this.output) - i)
		this.output[i], this.output[i+id] = this.output[i+id], this.output[i]
	}
	return this.output
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	obj := Constructor(nums)
	param_1 := obj.Reset()
	fmt.Printf("%s\n", strings.Join(util.IntArray2StringArray(param_1), ","))
	param_2 := obj.Shuffle()
	fmt.Printf("%s\n", strings.Join(util.IntArray2StringArray(param_2), ","))
}
