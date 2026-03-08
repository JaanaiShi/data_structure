package main

import (
	"fmt"
	"math"
)

// lengthOfLongestSubstringKDistinct 340. 至多包含 K 个不同字符的最长子串 (LeetCode 340)
// 给定字符串 s 和整数 k，找出 s 的最长子串，使其中不同字符的数量 ≤ k。
func lengthOfLongestSubstringKDistinct(s string, k int) int {
	var (
		window      = make(map[byte]int)
		left, right = 0, 0
		maxLen      = math.MinInt
	)

	for right < len(s) {
		temp := s[right]
		right++

		// 找到超过k的窗口
		window[temp]++
		if len(window) <= k {
			if count(window) > maxLen {
				maxLen = count(window)
			}
		} else {
			leftValue := s[left]
			left++
			if v := window[leftValue]; v > 1 {
				window[leftValue]--
			} else {
				delete(window, leftValue)
			}
		}
	}

	if maxLen == math.MinInt {
		return 0
	}

	return maxLen
}

func count(data map[byte]int) (res int) {
	for _, v := range data {
		res += v
	}
	return
}

func lengthOfLongestSubstringKDistinctV2(s string, k int) int {
	var (
		window      = make(map[byte]int)
		left, right = 0, 0
		maxLen      = 0
	)

	for right < len(s) {
		temp := s[right]
		right++
		window[temp]++

		for len(window) > k {
			leftValue := s[left]
			left++

			window[leftValue]--
			if window[leftValue] == 0 {
				delete(window, leftValue)
			}
		}

		if right-left > maxLen {
			maxLen = right - left
		}
	}
	return maxLen
}

func main() {
	tests := []struct {
		s    string
		k    int
		want int
	}{
		{"eceba", 2, 3},
		{"aa", 1, 2},
	}

	fmt.Println("开始测试 sliding_window/01_longest_at_most_k_distinct...")
	for _, tt := range tests {
		got := lengthOfLongestSubstringKDistinctV2(tt.s, tt.k)
		if got == tt.want {
			fmt.Printf("PASS: s=%q, k=%d, want=%d, got=%d\n", tt.s, tt.k, tt.want, got)
		} else {
			fmt.Printf("FAIL: s=%q, k=%d, want=%d, got=%d\n", tt.s, tt.k, tt.want, got)
		}
	}
	fmt.Println("测试结束。")
}
