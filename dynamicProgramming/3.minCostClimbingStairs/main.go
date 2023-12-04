package main

import "fmt"

// 解法错误，理解上的错误
// func minCostClimbingStairs(cost []int) int {
// 	if len(cost) == 1{
// 		return cost[0]
// 	}else if len(cost) == 2 {
// 		return min(cost[0], cost[1])
// 	}
// 	dp1, dp2 := cost[0], cost[1]
// 	for i:=2; i < len(cost); i++ {
// 		dp1 += cost[i]
// 		dp2 = min(dp1, dp2)
// 		dp1, dp2 = dp2, dp1
// 		fmt.Printf("dp1: %d, dp2: %d\n", dp1, dp2)
// 	}
// 	return min(dp1, dp2)
// }


func minCostClimbingStairs(cost []int) int {
	var (
		dp  = make([]int, len(cost) + 1)
	)
	if len(cost) == 0 || len(cost) == 1 {
		return 0
	}
	// 初始化
	dp1, dp2 := 0, 0
	for i:=2; i < len(dp); i++ {
		dp[i] = min(dp1 + cost[i-2], dp2 + cost[i-1])
		dp1 = dp[i-1]
		dp2 = dp[i]
	}
	return dp[len(cost)]
}



func min(a,b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(minCostClimbingStairs([]int{1,100,1,1,1,100,1,1,100,1})) 
}