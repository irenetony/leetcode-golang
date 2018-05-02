package main

import "fmt"

func letterCombinations(digits string) []string {
	arr := []rune(digits)
	if len(arr) == 0 {
		return nil
	} else if len(arr) == 1 {
		res := curLetterArr(arr[0])
		return res
	} else {
		var res []string
		curLetterCombinations := curLetterArr(arr[0])
		nextLetterCombinations := letterCombinations(digits[1:])
		for _, cur := range curLetterCombinations {
			for _, next := range nextLetterCombinations {
				res = append(res, string(cur)+next)
			}
		}
		return res
	}
}

func curLetterArr(cur rune) []string {
	var res []string
	if cur != '7' && cur != '9' {
		if cur != '8' {
			res = []string{string('a' + (cur-'2')*3), string('a' + (cur-'2')*3 + 1), string('a' + (cur-'2')*3 + 2)}
		} else {
			res = []string{string('a' + (cur-'2')*3 + 1), string('a' + (cur-'2')*3 + 2), string('a' + (cur-'2')*3 + 3)}
		}
	} else {
		if cur == '7' {
			res = []string{string('a' + (cur-'2')*3), string('a' + (cur-'2')*3 + 1), string('a' + (cur-'2')*3 + 2),
				string('a' + (cur-'2')*3 + 3)}
		} else {
			res = []string{string('a' + (cur-'2')*3+1), string('a' + (cur-'2')*3 + 2), string('a' + (cur-'2')*3 + 3),
				string('a' + (cur-'2')*3 + 4)}
		}
	}
	return res
}

func main() {
	fmt.Printf("%s=%+v\n", "23", letterCombinations("23"))
}
