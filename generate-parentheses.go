package main

import (
	"fmt"
	ds "github.com/berryjam/leetcode-golang/datastructure"
)

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{""}
	}

	res := make([]string, 0)

	stack := ds.NewStack()
	cur := ""
	stack.Push(-1)
	for !stack.IsEmpty() {
		if isValid(cur, "(", n) {
			cur += "("
			stack.Push(len(cur) - 1)
		} else {
			if isValid(cur, ")", n) {
				cur += ")"
				if len(cur) == 2*n {
					if isValid(cur, "", n) {
						res = append(res, cur)
					} else {
						break
					}
					val := stack.Pop()
					if val.(int) == -1 {
						break
					}
					cur = cur[0:val.(int)] + ")"
				}
			} else {
				val := stack.Pop()
				if val.(int) == -1 {
					break
				}
				cur = cur[0:val.(int)] + ")"
			}
		}
	}

	return res
}

func isValid(cur, p string, n int) bool {
	leftCount := 0
	rightCount := 0
	for _, r := range cur {
		if string(r) == "(" {
			leftCount++
		} else if string(r) == ")" {
			rightCount++
		}
		if rightCount > leftCount {
			return false
		}
	}

	if p == "(" {
		leftCount++
	} else if p == ")" {
		rightCount++
	}

	if leftCount >= rightCount && leftCount <= n {
		return true
	}

	return false
}

func TestGenerateParenthesis(n int) {
	parenthesis := generateParenthesis(n)
	for _, val := range parenthesis {
		fmt.Printf("%s\n", val)
	}
}

func main() {
	TestGenerateParenthesis(3)
}
