package util

import "strconv"

func IntArray2StringArray(array []int) (res []string) {
	for _, val := range array {
		res = append(res, strconv.Itoa(val))
	}
	return
}
