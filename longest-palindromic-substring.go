package main

import (
	"fmt"
)

func longestPalindrome(s string) string {
	res := ""
	sArr := []rune(s)
	for i := 0; i < len(sArr); i++ {
		if tmp := longestPalindromeHelper(s, i); len(tmp) > len(res) {
			res = tmp
		}
	}

	return res
}

func longestPalindromeHelper(s string, mid int) string {
	l := mid - 1
	r := mid + 1
	sArr := []rune(s)
	for r < len(sArr) && sArr[r] == sArr[mid] {
		r++
	}
	for l >= 0 && sArr[l] == sArr[mid] {
		l--
	}
	for r < len(sArr) && l >= 0 && sArr[l] == sArr[r] {
		r++
		l--
	}
	return string(sArr[l+1:r])
}

func TestLongestPalindrome(s string) {
	fmt.Printf("Input: \"%s\"\n", s)
	fmt.Println()
	fmt.Printf("Output: \"%s\"\n", longestPalindrome(s))
	fmt.Println()
}

func main() {
	TestLongestPalindrome("babad")
	TestLongestPalindrome("我是中是我的")
	TestLongestPalindrome("cbbd")
}
