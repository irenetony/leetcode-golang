package util

import "strconv"

func IntArray2StringArray(array []int) (res []string) {
	for _, val := range array {
		res = append(res, strconv.Itoa(val))
	}
	return
}

func RuneArray2StringArray(array []rune) (res []string) {
	for _, val := range array {
		res = append(res, strconv.QuoteRune(val))
	}
	return
}
