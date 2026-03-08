package main

import "fmt"

// checkInclusion 567. 字符串的排列 (LeetCode 567)
// 给定两个字符串 s1 和 s2，判断 s2 是否包含 s1 的任意排列作为子串。
func checkInclusion(s1 string, s2 string) bool {
	return false
}

func main() {
	tests := []struct {
		s1   string
		s2   string
		want bool
	}{
		{"ab", "eidbaooo", true},
		{"ab", "eidboaoo", false},
	}

	fmt.Println("开始测试 sliding_window/03_permutation_in_string...")
	for _, tt := range tests {
		got := checkInclusion(tt.s1, tt.s2)
		if got == tt.want {
			fmt.Printf("PASS: s1=%q, s2=%q, want=%v, got=%v\n", tt.s1, tt.s2, tt.want, got)
		} else {
			fmt.Printf("FAIL: s1=%q, s2=%q, want=%v, got=%v\n", tt.s1, tt.s2, tt.want, got)
		}
	}
	fmt.Println("测试结束。")
}
