package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	res := 0

	runeMap := make(map[rune]int)

	runeArr := []rune(s)

	cur := 0

	for i := 0; i < len(runeArr); i++ {

		if rIdx, ok := runeMap[runeArr[i]]; !ok {
			runeMap[runeArr[i]] = i
			cur++
			if len(runeMap) > res {
				res = len(runeMap)
			}
		} else {
			for k, v := range runeMap {
				if v <= rIdx {
					delete(runeMap, k)
				}
			}
			runeMap[runeArr[i]] = i
			if cur > res {
				res = cur
			}
			if len(runeMap) > res {
				res = len(runeMap)
			}
			cur = 1
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
}
