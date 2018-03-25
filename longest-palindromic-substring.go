package main

import (
	"fmt"
)

func longestPalindrome(s string) string {
	sArray := []rune(s)
	newArray := make([]rune, 2*len(sArray)+1)
	for i := 0; i < 2*len(sArray)+1; i++ {
		if i%2 == 0 {
			newArray[i] = 0x00
		} else {
			newArray[i] = sArray[i/2]
		}
	}

	RL := make([]int, len(newArray))
	for i := 0; i < len(RL); i++ {
		RL[i] = 0
	}
	maxRight := 0
	pos := 0
	maxLen := 0
	index := 0
	for i := 0; i < len(newArray)-1; i++ {
		if i < maxRight {
			if RL[2*pos-i] < maxRight-i {
				RL[i] = RL[2*pos-i]
			} else {
				RL[i] = maxRight - i
			}
		} else {
			RL[i] = 1
		}

		for i-RL[i] >= 0 && i+RL[i] < len(newArray) && newArray[i+RL[i]] == newArray[i-RL[i]] {
			RL[i]++
		}

		if i+RL[i] > maxRight {
			maxRight = i + RL[i]
			pos = i
		}
	}

	for i := 1; i < len(newArray)-1; i++ {
		if RL[i] > maxLen {
			maxLen = RL[i]
			index = i
		}
	}

	return string(sArray[(index-maxLen)/2:(index-maxLen)/2+maxLen-1])
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
