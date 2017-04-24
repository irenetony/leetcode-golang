package main

func maxProfitII(prices []int) int {
	if prices == nil || len(prices) < 2 {
		return 0
	}
	profit := 0
	for i := 0; i < len(prices)-1; i++ {
		diff := prices[i+1] - prices[i]
		if diff > 0 {
			profit += diff
		}
	}
	return profit
}

func main() {
}
