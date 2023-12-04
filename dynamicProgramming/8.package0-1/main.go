package main

import "fmt"


func MaxValue(weight, value []int, sum int) int {
	var (
		dp = make([][]int, len(weight))
	)
	// 构建dp数组
	for  i:=0; i < len(dp); i++ {
		dp[i] = make([]int, sum + 1, sum + 1)
	}
	// 进行初始化
	
	for i:=1; i < len(weight); i++ {
		for j:=0; j < sum; j++ {
			dp[i][j] = max(dp[i-1][j], dp[i-1][j - weight[i]] + value[i])
		}
	}

	fmt.Println(dp)
	return dp[len(weight)-1][sum]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	MaxValue([]int{1,3,4}, []int{15,20,30}, 4)	
}