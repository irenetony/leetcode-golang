package main

import "fmt"

func minDistance(word1,word2 string) int {
	aArr := []rune(word1)
	bArr := []rune(word2)

	D := make([][]int, len(aArr)+1)
	for i := 0; i <= len(aArr); i++ {
		D[i] = make([]int, len(bArr)+1)
		D[i][0] = i
	}

	for j := 0; j <= len(bArr); j++ {
		D[0][j] = j
	}

	for i := 1; i < len(aArr)+1; i++ {
		for j := 1; j < len(bArr)+1; j++ {
			del := D[i-1][j] + 1
			add := D[i][j-1] + 1
			var flag int
			if aArr[i-1] == bArr[j-1] {
				flag = 0
			} else {
				flag = 1
			}
			update := D[i-1][j-1] + flag
			D[i][j] = min(min(del, add), update)
		}
	}

	return D[len(aArr)][len(bArr)]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Printf("res=%d", minDistance("qqqq", "w"))
}
