package main

import (
	"fmt"
)

// maxArea 使用双指针算法计算容器能盛的最多水量。
//
// 题目说明：
// 给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
// 找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
// 返回容器可以储存的最大水量。
// 说明：你不能倾斜容器。
//
// 示例 1：
// 输入：[1,8,6,2,5,4,8,3,7]
// 输出：49
// 解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
//
// 示例 2：
// 输入：height = [1,1]
// 输出：1
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	res := 0
	for left < right {
		temp := (right - left) * min(height[left], height[right])
		if temp > res {
			res = temp
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return res
}

func main() {
	tests := []struct {
		height []int
		want   int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
		{[]int{1, 9, 10}, 20},
	}

	fmt.Println("开始测试 02_container_most_water...")
	for _, tt := range tests {
		got := maxArea(tt.height)
		if got == tt.want {
			fmt.Printf("PASS: height=%v, want=%d, got=%d\n", tt.height, tt.want, got)
		} else {
			fmt.Printf("FAIL: height=%v, want=%d, got=%d\n", tt.height, tt.want, got)
		}
	}
	fmt.Println("测试结束。")
}
