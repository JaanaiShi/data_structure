package main

import "fmt"

// longestOnes 1004. 最大连续1的个数 III (LeetCode 1004)
// 给定二进制数组 nums 和整数 k，最多翻转 k 个 0，求最长连续 1 的长度。
func longestOnes(nums []int, k int) int {
	return 0
}

func main() {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2, 6},
		{[]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3, 10},
	}

	fmt.Println("开始测试 sliding_window/07_max_consecutive_ones_iii...")
	for _, tt := range tests {
		got := longestOnes(tt.nums, tt.k)
		if got == tt.want {
			fmt.Printf("PASS: nums=%v, k=%d, want=%d, got=%d\n", tt.nums, tt.k, tt.want, got)
		} else {
			fmt.Printf("FAIL: nums=%v, k=%d, want=%d, got=%d\n", tt.nums, tt.k, tt.want, got)
		}
	}
	fmt.Println("测试结束。")
}
