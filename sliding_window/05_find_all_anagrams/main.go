package main

import (
	"fmt"
	"reflect"
)

// findAnagrams 438. 找到字符串中所有字母异位词 (LeetCode 438)
// 给定字符串 s 和 p，返回 s 中所有 p 的异位词的起始索引数组。
func findAnagrams(s string, p string) []int {
	return nil
}

func main() {
	tests := []struct {
		s    string
		p    string
		want []int
	}{
		{"cbaebabacd", "abc", []int{0, 6}},
		{"abab", "ab", []int{0, 1, 2}},
	}

	fmt.Println("开始测试 sliding_window/05_find_all_anagrams...")
	for _, tt := range tests {
		got := findAnagrams(tt.s, tt.p)
		if reflect.DeepEqual(got, tt.want) {
			fmt.Printf("PASS: s=%q, p=%q, want=%v, got=%v\n", tt.s, tt.p, tt.want, got)
		} else {
			fmt.Printf("FAIL: s=%q, p=%q, want=%v, got=%v\n", tt.s, tt.p, tt.want, got)
		}
	}
	fmt.Println("测试结束。")
}
