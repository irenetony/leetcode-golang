package main

import "fmt"

var aArr []rune
var bArr []rune
var D [][]int

func minStepRectify(a, b string) int {
	aArr = []rune(a)
	bArr = []rune(b)

	D = make([][]int, len(aArr)+1)
	for i := range D {
		D[i] = make([]int, len(bArr)+1)
		for j := range D[i] {
			D[i][j] = -1
		}
	}

	return getD(0, 0)
}

func getD(i, j int) int {
	if i == len(aArr)+1 || j == len(bArr)+1 {
		return 0
	}
	if D[i][j] != -1 {
		return D[i][j]
	} else {
		if i == len(aArr) || j == len(bArr) {
			D[i][j] = 0
		} else {
			if aArr[i] == bArr[j] {
				D[i][j] = getD(i+1, j+1)
			} else {
				del := getD(i, j+1) + 1
				add := getD(i, j) + 1
				update := getD(i+1, j+1) + 1
				D[i][j] = min(min(del, add), update)
			}
		}
	}

	return D[i][j]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Printf("res=%d", minStepRectify("abb", "b"))
}
