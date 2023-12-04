package main

import "fmt"


func numTrees(n int) int {
	var (
		dp = make([]int, n+1, n+1)
	)	
	// 空树也算是一颗树，所以为1
	dp[0] = 1; dp[1] = 1
	for i:=2; i <= n; i++ {
		for j:=0; j < i; j++ {
			dp[i] += dp[j] * dp[i - j - 1]
		}
	}
	return dp[n]
}

func main() {
	fmt.Println(numTrees(3))
}