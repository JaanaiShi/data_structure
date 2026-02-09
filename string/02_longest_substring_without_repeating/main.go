package main

import (
	"fmt"
)

// lengthOfLongestSubstring 给定一个字符串 s，请你找出其中不含有重复字符的最长子串的长度。
// 示例 1: 输入: s = "abcabcbb" 输出: 3 (因为无重复字符的最长子串是 "abc")
// 示例 2: 输入: s = "bbbbb" 输出: 1 (因为无重复字符的最长子串是 "b")
// 示例 3: 输入: s = "pwwkew" 输出: 3 (因为无重复字符的最长子串是 "wke")
// 提示: 0 <= s.length <= 5 * 10^4, s 由英文字母、数字、符号和空格组成
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	slow, fast := 0, 1
	notDuplication := make(map[byte]struct{}, len(s))
	notDuplication[s[slow]] = struct{}{}
	maxLength := len(notDuplication)
	for fast < len(s) {
		if _, ok := notDuplication[s[fast]]; ok {
			delete(notDuplication, s[slow])
			slow++
			continue
		} else {
			notDuplication[s[fast]] = struct{}{}
			if len(notDuplication) > maxLength {
				maxLength = len(notDuplication)
			}
		}
		fast++
	}
	return maxLength
}

func main() {
	tests := []struct {
		s    string
		want int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
		{" ", 1},
		{"au", 2},
		{"dvdf", 3},
	}

	fmt.Println("开始测试 string/02_longest_substring_without_repeating...")
	for _, tt := range tests {
		got := lengthOfLongestSubstring(tt.s)
		if got == tt.want {
			fmt.Printf("PASS: s=%q, want=%d, got=%d\n", tt.s, tt.want, got)
		} else {
			fmt.Printf("FAIL: s=%q, want=%d, got=%d\n", tt.s, tt.want, got)
		}
	}
	fmt.Println("测试结束。")
}
