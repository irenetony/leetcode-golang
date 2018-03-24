package main

import "fmt"

func longestPalindrome(s string) string {
	array := []rune(s)
	res := ""
	maxLength := 0
	for i := 0; i < len(array); i++ {
		for j := i; j < len(array); j++ {
			if isPalindromeStr(string(array[i:j+1])) && maxLength < (j-i+1) {
				maxLength = j - i + 1
				res = string(array[i:j+1])
			}
		}
	}
	return res
}

func isPalindromeStr(s string) bool {
	array := []rune(s)
	for i := 0; i < len(array)/2; i++ {
		if array[i] != array[len(array)-1-i] {
			return false
		}
	}
	return true
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
}
