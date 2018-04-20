package main

import "fmt"

func uniquePaths(m int, n int) int {
	if m < 1 || n < 1 {
		return -1
	}
	D := make([][]int, m+1)
	for i := range D {
		D[i] = make([]int, n+1)
		D[i][1] = 1
	}

	for j := 0; j < n+1; j++ {
		D[1][j] = 1
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if i == 1 || j == 1 {
				D[i][j] = 1
			} else {
				D[i][j] = D[i-1][j] + D[i][j-1]
			}
		}
	}

	return D[m][n]
}

func main() {
	fmt.Printf("res=%d\n", uniquePaths(3, 2))
	fmt.Printf("res=%d\n", uniquePaths(7, 3))
}
