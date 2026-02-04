package main

import (
	"fmt"
	"leetcode/common"
	"slices"
)

// isValid 判断括号字符串是否有效。
// 有效字符串需满足：
// 1. 左括号必须用相同类型的右括号闭合。
// 2. 左括号必须以正确的顺序闭合。
//
// 使用栈来解决：
// 1. 遇到左括号，压入栈中。
// 2. 遇到右括号，检查栈顶元素是否是对应的左括号。
func isValid(s string) bool {
	stack := common.Stack{}
	for _, v := range s {
		if slices.Contains([]rune{'(', '[', '{'}, v) {
			stack.Push(v)
		} else {
			item := stack.Pop()
			if (item == '(' && v != ')') ||
				(item == '[' && v != ']') ||
				(item == '{' && v != '}') || item == 0 {
				return false
			}
		}
	}
	if stack.Count() == 0 {
		return true
	}
	return false
}

func main() {
	tests := []struct {
		s    string
		want bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
		{"]", false},
		{"", true},
	}

	fmt.Println("开始测试 stack/01_valid_parentheses...")
	for _, tt := range tests {
		got := isValid(tt.s)
		if got == tt.want {
			fmt.Printf("PASS: s=%q, want=%v, got=%v\n", tt.s, tt.want, got)
		} else {
			fmt.Printf("FAIL: s=%q, want=%v, got=%v\n", tt.s, tt.want, got)
		}
	}
	fmt.Println("测试结束。")
}
