package main

import "fmt"

// characterReplacement 424. 替换后的最长重复字符 (LeetCode 424)
// 给定字符串 s 和整数 k，最多替换 k 个字符后，求能得到的最长重复字符子串的长度。
func characterReplacement(s string, k int) int {
	return 0
}

func main() {
	tests := []struct {
		s    string
		k    int
		want int
	}{
		{"ABAB", 2, 4},
		{"AABABBA", 1, 4},
	}

	fmt.Println("开始测试 sliding_window/06_longest_repeating_character_replacement...")
	for _, tt := range tests {
		got := characterReplacement(tt.s, tt.k)
		if got == tt.want {
			fmt.Printf("PASS: s=%q, k=%d, want=%d, got=%d\n", tt.s, tt.k, tt.want, got)
		} else {
			fmt.Printf("FAIL: s=%q, k=%d, want=%d, got=%d\n", tt.s, tt.k, tt.want, got)
		}
	}
	fmt.Println("测试结束。")
}
