package main

import (
	"fmt"
)

// isPalindrome 验证回文串
//
// 题目说明：
// 如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，短语正着读和反着读都一样，则可以认为该短语是一个 回文串 。
// 字母和数字都属于字母数字字符。
// 给你一个字符串 s，如果它是 回文串 ，返回 true ；否则，返回 false 。
//
// 示例 1：
// 输入：s = "A man, a plan, a canal: Panama"
// 输出：true
// 解释："amanaplanacanalpanama" 是回文串。
//
// 示例 2：
// 输入：s = "race a car"
// 输出：false
// 解释："raceacar" 不是回文串。
//
// 示例 3：
// 输入：s = " "
// 输出：true
// 解释：在移除非字母数字字符之后，s 是一个空字符串 "" 。
// 由于空字符串正着反着读都一样，所以是回文串。
func isPalindrome(s string) bool {
	// TODO: 请实现双指针算法
	// 提示：你可以使用 unicode.IsLetter 和 unicode.IsDigit 来判断字符
	// strings.ToLower 可以转换大小写
	return false
}

func main() {
	tests := []struct {
		s    string
		want bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{" ", true},
		{"0P", false},
	}

	fmt.Println("开始测试 05_valid_palindrome...")
	for _, tt := range tests {
		got := isPalindrome(tt.s)
		if got == tt.want {
			fmt.Printf("PASS: s=%q, want=%v, got=%v\n", tt.s, tt.want, got)
		} else {
			fmt.Printf("FAIL: s=%q, want=%v, got=%v\n", tt.s, tt.want, got)
		}
	}
	fmt.Println("测试结束。")
}
