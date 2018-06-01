package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	res := 0

	runeMap := make(map[rune]int)

	runeArr := []rune(s)

	lastIdx := 0

	for i := 0; i < len(runeArr); i++ {

		if rIdx, ok := runeMap[runeArr[i]]; !ok {
			runeMap[runeArr[i]] = i
			if i-lastIdx+1 > res {
				res = i - lastIdx + 1
			}
		} else {
			if rIdx > lastIdx {
				lastIdx = rIdx
			}
			if runeArr[i] != runeArr[lastIdx] {
				if i-lastIdx+1 > res {
					res = i - lastIdx + 1
				}
			} else {
				if i-lastIdx > res {
					res = i - lastIdx
				}
			}

			lastIdx = rIdx + 1
			runeMap[runeArr[i]] = i
		}
	}

	return res
}

func main() {
	fmt.Printf("res=%d\n", lengthOfLongestSubstring("abcabcbb"))
	fmt.Printf("res=%d\n", lengthOfLongestSubstring("bbbbb"))
	fmt.Printf("res=%d\n", lengthOfLongestSubstring("pwwkew"))
	fmt.Printf("res=%d\n", lengthOfLongestSubstring(""))
	fmt.Printf("res=%d\n", lengthOfLongestSubstring("c"))
	fmt.Printf("res=%d\n", lengthOfLongestSubstring("dvdf"))
	fmt.Printf("res=%d\n", lengthOfLongestSubstring("abba"))
	fmt.Printf("res=%d\n", lengthOfLongestSubstring("tmmzuxt"))
}
