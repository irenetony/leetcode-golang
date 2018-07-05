package main

import "fmt"

func numDecodings(s string) int {
	if s == "" {
		return 0
	}
	runeArr := []rune(s)
	count := 0
	if len(runeArr) == 0 {
		return 0
	}
	if len(runeArr) == 1 && runeArr[0] != '0' {
		return 1
	} else if len(runeArr) == 2 {
		if runeArr[0] == '1' {
			if runeArr[1] != '0' {
				count += 2
			} else {
				count += 1
			}
		} else if runeArr[0] == '2' {
			if runeArr[1] == '0' {
				count += 1
			} else if runeArr[1] >= '1' && runeArr[1] <= '6' {
				count += 2
			} else if runeArr[1] >= '7' && runeArr[1] <= '9' {
				count += 1
			}
		} else if runeArr[0] >= '3' && runeArr[0] <= '9' {
			if runeArr[1] != '0' {
				count += 1
			}
		}
	} else {
		if runeArr[0] == '1' {
			count += numDecodings(string(runeArr[1:]))
			count += numDecodings(string(runeArr[2:]))
		} else if runeArr[0] == '2' {
			if runeArr[1] == '0' {
				count += numDecodings(string(runeArr[2:]))
			} else if runeArr[1] >= '1' && runeArr[1] <= '6' {
				count += numDecodings(string(runeArr[1:]))
				count += numDecodings(string(runeArr[2:]))
			} else if runeArr[1] >= '7' && runeArr[1] <= '9' {
				count += numDecodings(string(runeArr[2:]))
			}
		} else if runeArr[0] >= '3' && runeArr[0] <= '9' {
			if runeArr[1] != '0' {
				count += numDecodings(string(runeArr[1:]))
			}
		}
	}
	return count
}

func main() {
	fmt.Printf("res=%d\n", numDecodings("12"))
	fmt.Printf("res=%d\n", numDecodings("226"))
	fmt.Printf("res=%d\n", numDecodings("206"))
	fmt.Printf("res=%d\n", numDecodings("611"))
	fmt.Printf("res=%d\n", numDecodings("9355141842519423434975558424574331479338913773268525913972842463395324957127519863135646135786662832"))
}
