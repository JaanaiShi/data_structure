package main

import "fmt"

func climbStairs(n int) int {
	if n <= 1{
		return n
	}
	dp1, dp2 := 1, 2
	for i:=2; i < n; i++ {
		temp := dp1
		dp1 = dp2
		dp2 = temp + dp2
	}
	return dp2
}

func main() {
	fmt.Println(climbStairs(5))
}