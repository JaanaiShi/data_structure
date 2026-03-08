package main

import "fmt"

// subarraysWithKDistinct 992. K 个不同整数的子数组 (LeetCode 992)
// 给定整数数组 nums 和整数 k，返回子数组个数，使每个子数组恰好包含 k 个不同整数。
func subarraysWithKDistinct(nums []int, k int) int {
	return subarraysWithAtMostKDistinct(nums, k) - subarraysWithAtMostKDistinct(nums, k-1)
}

func subarraysWithAtMostKDistinct(nums []int, k int) int {
	var (
		window      = make(map[int]int)
		left, right = 0, 0
		res         = 0
	)
	for right < len(nums) {
		temp := nums[right]
		right++

		window[temp]++

		for len(window) > k {
			leftValue := nums[left]
			window[leftValue]--
			if window[leftValue] <= 0 {
				delete(window, leftValue)
			}
			left++
		}
		res += right - left
	}

	return res
}

func main() {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{1, 2, 1, 2, 3}, 2, 7},
		{[]int{1, 2, 1, 3, 4}, 3, 3},
	}

	fmt.Println("开始测试 sliding_window/02_subarrays_with_exactly_k_distinct...")
	for _, tt := range tests {
		got := subarraysWithKDistinct(tt.nums, tt.k)
		if got == tt.want {
			fmt.Printf("PASS: nums=%v, k=%d, want=%d, got=%d\n", tt.nums, tt.k, tt.want, got)
		} else {
			fmt.Printf("FAIL: nums=%v, k=%d, want=%d, got=%d\n", tt.nums, tt.k, tt.want, got)
		}
	}
	fmt.Println("测试结束。")
}
