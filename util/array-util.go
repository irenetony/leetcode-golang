package util

import (
	"strconv"
)

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

func TwoDimensionToString(array [][]int) string {
	res := "["

	for idx := range array {
		if idx >= 1 {
			res += " "
		}
		if idx != len(array)-1 {
			if len(array) > 1 {
				res += OneDimensionToString(array[idx]) + ",\n"
			} else {
				res += OneDimensionToString(array[idx])
			}
		} else {
			res += OneDimensionToString(array[idx])
		}
	}

	res += "]"
	return res
}

func OneDimensionToString(array []int) string {
	res := "["

	for idx := range array {
		if idx != len(array)-1 {
			if len(array) > 1 {
				res += strconv.Itoa(array[idx]) + ","
			} else {
				res += strconv.Itoa(array[idx])
			}
		} else {
			res += strconv.Itoa(array[idx])
		}
	}

	res += "]"

	return res
}
