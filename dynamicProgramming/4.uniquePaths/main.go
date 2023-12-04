package main

import "fmt"

func uniquePaths(m int, n int) int {
	
	dp := make([][]int, m)
	for i:=0; i < len(dp); i++ {
		dp[i] = make([]int, n)
	}
	// 初始化
	for i:=0; i < m; i++ {
		dp[i][0] = 1
	}
	for i:=0; i < n; i++ {
		dp[0][i] = 1
	}
	// 确定遍历顺序
	for j:=1; j < m; j++ {
		for i:=1; i < n; i++ {
			dp[j][i] = dp[j-1][i] + dp[j][i-1]
		}
	}
	for _, v := range dp {
		fmt.Println(v)
	}

	return dp[m-1][n-1]
}

func main() {
	uniquePaths(3, 3)
}