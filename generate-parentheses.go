package main

import "fmt"

func generateParenthesis(n int) []string {
	res := make([]string, 0)

	if n == 0 {
		res = append(res, "")
	} else if n == 1 {
		res = append(res, "()")
	} else {
		nextParenthesis := generateParenthesis(n - 1)
		for _, parenthesis := range nextParenthesis {
			res = append(res, "("+parenthesis+")")
		}
		for _, parenthesis := range nextParenthesis {
			if isSingleParenthesis := isAllSingleParenthesis(parenthesis); isSingleParenthesis {
				res = append(res, "()"+parenthesis)
			} else {
				res = append(res, parenthesis+"()")
				res = append(res, "()"+parenthesis)
			}
		}
	}

	return res
}

func isAllSingleParenthesis(parenthesis string) bool {
	for idx, runeValues := range parenthesis {
		if idx%2 == 0 {
			if runeValues != '(' {
				return false
			}
		}
		if idx%2 == 1 {
			if runeValues != ')' {
				return false
			}
		}
	}
	return true
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
