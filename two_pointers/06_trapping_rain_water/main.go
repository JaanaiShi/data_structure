package main

import "fmt"

// trap 42. 接雨水 (LeetCode 42)
// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
func trap(height []int) int {
	left, right, zero := 0, 0, 0
	res := 0
	leftVal := height[left]
	for right < len(height) {
		temp := height[right]
		right++
		if temp == 0 {
			zero = right
		}

		if temp >= leftVal || (zero != 0 && temp != 0) {
			// 计算当前的值
			minVal := min(leftVal, temp)
			max := minVal * minVal
			for i := left; i < right; i++ {
				max -= height[i]
			}
			res += max
			left = right
			leftVal = temp
			zero = 0
		}
	}
	if left != right {
		minVal := min(leftVal, height[right-1])
		max := minVal * minVal
		for i := left; i < right; i++ {
			max -= height[i]
		}
		res += max
	}
	return res
}

func main() {
	tests := []struct {
		height []int
		want   int
	}{
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		{[]int{4, 2, 0, 3, 2, 5}, 9},
	}

	fmt.Println("开始测试 two_pointers/06_trapping_rain_water...")
	for _, tt := range tests {
		got := trap(tt.height)
		if got == tt.want {
			fmt.Printf("PASS: height=%v, want=%d, got=%d\n", tt.height, tt.want, got)
		} else {
			fmt.Printf("FAIL: height=%v, want=%d, got=%d\n", tt.height, tt.want, got)
		}
	}
	fmt.Println("测试结束。")
}
