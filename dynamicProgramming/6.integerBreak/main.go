package main

import "fmt"

func integerBreak(n int) int {
	var (
		dp = make([]int, n+1, n+1)
	)
	if n == 0 || n == 1 {
		return n
	}
	for i:=2; i <= n; i++ {
		for j := 1; j < i; j++ {
			fmt.Println(dp)
			dp[i] = max(dp[i], max(j*(i-j), dp[i-j]*j))
		}
	}

	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}



func main() {
	fmt.Println(integerBreak(10))
}