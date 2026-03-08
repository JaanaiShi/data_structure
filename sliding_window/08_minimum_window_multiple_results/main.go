package main

import (
	"fmt"
	"reflect"
)

// minWindowAll 76. 最小覆盖子串 (LeetCode 76 变体，原题返回子串)
// 给定字符串 s 和字符集合 C（可能包含重复），返回 s 中包含 C 的所有最短子串的起始索引。
func minWindowAll(s string, t string) []int {
	return nil
}

func main() {
	tests := []struct {
		s    string
		t    string
		want []int
	}{
		{"abcbcba", "abc", []int{1, 3}},
	}

	fmt.Println("开始测试 sliding_window/08_minimum_window_multiple_results...")
	for _, tt := range tests {
		got := minWindowAll(tt.s, tt.t)
		if reflect.DeepEqual(got, tt.want) {
			fmt.Printf("PASS: s=%q, t=%q, want=%v, got=%v\n", tt.s, tt.t, tt.want, got)
		} else {
			fmt.Printf("FAIL: s=%q, t=%q, want=%v, got=%v\n", tt.s, tt.t, tt.want, got)
		}
	}
	fmt.Println("测试结束。")
}
