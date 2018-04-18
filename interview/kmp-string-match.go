package main

import "fmt"

func matchSubString(s string, p string) int {
	sArr := []rune(s)
	pArr := []rune(p)
	slen := len(sArr)
	plen := len(pArr)

	i, j := 0, 0

	for i < slen && j < plen {
		if sArr[i] == pArr[j] {
			i++
			j++
		} else {
			i = i - (j - 1)
			j = 0
		}
	}

	if j == plen {
		return i - j
	} else {
		return -1
	}
}

func main() {
	res := matchSubString("BBCABCDABABCDABCDABDE", "ABDE")
	fmt.Printf("res=%d\n", res)
}
