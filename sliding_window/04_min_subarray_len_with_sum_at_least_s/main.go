package main

import "fmt"

// minSubArrayLen 209. 长度最小的子数组 (LeetCode 209)
// 给定正整数数组 nums 和目标 S，找出最短子数组长度，使子数组和 ≥ S，若不存在则返回 0。
func minSubArrayLen(target int, nums []int) int {
	return 0
}

func main() {
	tests := []struct {
		target int
		nums   []int
		want   int
	}{
		{7, []int{2, 3, 1, 2, 4, 3}, 2},
		{4, []int{1, 4, 4}, 1},
		{11, []int{1, 1, 1, 1, 1, 1, 1, 1}, 0},
	}

	fmt.Println("开始测试 sliding_window/04_min_subarray_len_with_sum_at_least_s...")
	for _, tt := range tests {
		got := minSubArrayLen(tt.target, tt.nums)
		if got == tt.want {
			fmt.Printf("PASS: target=%d, nums=%v, want=%d, got=%d\n", tt.target, tt.nums, tt.want, got)
		} else {
			fmt.Printf("FAIL: target=%d, nums=%v, want=%d, got=%d\n", tt.target, tt.nums, tt.want, got)
		}
	}
	fmt.Println("测试结束。")
}
