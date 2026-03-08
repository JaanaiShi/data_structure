package main

import (
	"fmt"
)

// minWindow 给你一个字符串 s 和一个字符串 t，返回 s 中涵盖 t 所有字符的最小子串。
// 如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 ""。
// 注意：t 中的重复字符，在子串中也必须包含相同数量。
// 示例 1: 输入: s = "ADOBECODEBANC", t = "ABC" 输出: "BANC" (最小覆盖子串)
// 示例 2: 输入: s = "a", t = "a" 输出: "a"
// 示例 3: 输入: s = "a", t = "aa" 输出: "" (t 中有两个 'a'，s 中只有一个)
// 提示: 1 <= s.length, t.length <= 10^5, s 和 t 由英文字母组成
func minWindow(s string, t string) string {
	if len(t) == 0 || len(t) > len(s) {
		return ""
	}
	var (
		need   = [128]int{}
		window = [128]int{}
	)

	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	var (
		start, minLen = 0, 0
		having        = 0
		left, right   = 0, 0
	)
	for right < len(s) {
		temp := s[right]
		right++

		if need[temp] > 0 {
			window[temp]++
			if window[temp] <= need[temp] {
				having++
			}
		}

		// 左指针走，找最小的范围
		for having == len(t) {
			temp := s[left]
			// 找到一样的值则退出
			if need[temp] > 0 {
				// 和之前的比较
				if right-left < minLen || minLen == 0 {
					start = left
					minLen = right - left
				}
				window[temp]--
				if window[temp] < need[temp] {
					having--
				}
			}
			left++
		}
	}
	if minLen == 0 {
		return ""
	}
	return s[start : start+minLen]
}

func main() {
	tests := []struct {
		s    string
		t    string
		want string
	}{
		{"ADOBECODEBANC", "ABC", "BANC"},
		{"a", "a", "a"},
		{"a", "aa", ""},
		{"a", "b", ""},
		{"ab", "a", "a"},
		{"ab", "b", "b"},
		{"cabwefgewcwaefgcf", "cae", "cwae"},
	}

	fmt.Println("开始测试 string/03_minimum_window_substring...")
	for _, tt := range tests {
		got := minWindow(tt.s, tt.t)
		if got == tt.want {
			fmt.Printf("PASS: s=%q, t=%q, want=%q, got=%q\n", tt.s, tt.t, tt.want, got)
		} else {
			fmt.Printf("FAIL: s=%q, t=%q, want=%q, got=%q\n", tt.s, tt.t, tt.want, got)
		}
	}
	fmt.Println("测试结束。")
}
