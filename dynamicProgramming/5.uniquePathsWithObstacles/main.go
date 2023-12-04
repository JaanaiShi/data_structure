package main

import "fmt"



func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	// 初始化dp数组
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	
	dp := make([][]int, m)
	for i:=0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	// 初始化dp数组
	// 初始化第一列
	for i:=0; i < m && obstacleGrid[i][0] == 0; i++ {
		dp[i][0] = 1
	}
	// 初始化第一行
	for i:=0; i < n && obstacleGrid[0][i] == 0; i++ {
		dp[0][i] = 1
	}

	// 遍历
	for i:=1; i < m; i++ {
		for j := 1; j < m; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			}else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	fmt.Println(dp)
	return dp[m-1][n-1]
}

func main() {
	uniquePathsWithObstacles([][]int{
		{0, 0},
	})
}
