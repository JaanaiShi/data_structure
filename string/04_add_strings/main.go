package main

import (
	"fmt"
	"strconv"
)

func addStrings(num1 string, num2 string) string {
	n1, n2 := len(num1), len(num2)
	isAdd := int64(0)
	mut := int64(1)
	res := ""
	for n1 > 0 || n2 > 0 || isAdd > 0 {
		res1, res2 := int64(0), int64(0)
		if n1-1 >= 0 {
			res1, _ = strconv.ParseInt(string(num1[n1-1]), 10, 64)
		}
		if n2-1 >= 0 {
			res2, _ = strconv.ParseInt(string(num2[n2-1]), 10, 64)
		}
		sum := res1 + res2 + isAdd
		if sum >= 10 {
			isAdd = 1
			res = strconv.FormatInt(sum-10, 10) + res
		} else {
			isAdd = 0
			res = strconv.FormatInt(sum, 10) + res
		}
		n1--
		n2--
		mut *= 10
	}
	return res
}

func main() {
	// 测试用例
	tests := []struct {
		num1 string
		num2 string
		want string
	}{
		// 题目示例
		{"21543655", "4332656442", "4354200097"},
		// 基础用例
		{"11", "123", "134"},
		{"456", "77", "533"},
		// 进位传播
		{"9", "1", "10"},
		{"999", "1", "1000"},
		// 长度不等
		{"1", "9999999999999999999", "10000000000000000000"},
		// 含 0
		{"0", "0", "0"},
		{"0", "12345", "12345"},
	}

	fmt.Println("开始测试 string/04_add_strings...")
	for _, tt := range tests {
		got := addStrings(tt.num1, tt.num2)
		if got == tt.want {
			fmt.Printf("PASS: num1=%q + num2=%q = %q\n", tt.num1, tt.num2, got)
		} else {
			fmt.Printf("FAIL: num1=%q + num2=%q, want=%q, got=%q\n", tt.num1, tt.num2, tt.want, got)
		}
	}
	fmt.Println("测试结束。")
}
