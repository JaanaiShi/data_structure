package main

import (
	"fmt"
)

// firstUniqChar 找到字符串中第一个不重复的字符，并返回其索引。如果不存在，返回 -1。
func firstUniqChar(s string) int {
	sMap := make(map[byte]int, len(s))
	for i := 0; i < len(s); i++ {
		if v, ok := sMap[s[i]]; ok {
			v++
			sMap[s[i]] = v
		} else {
			sMap[s[i]] = 1
		}
	}
	for i := 0; i < len(s); i++ {
		if v := sMap[s[i]]; v == 1 {
			return i
		}
	}
	return -1
}

func main() {
	tests := []struct {
		s    string
		want int
	}{
		{"leetcode", 0},
		{"loveleetcode", 2},
		{"aabb", -1},
		{"z", 0},
		{"", -1},
	}

	fmt.Println("开始测试 string/01_first_unique_char...")
	for _, tt := range tests {
		got := firstUniqChar(tt.s)
		if got == tt.want {
			fmt.Printf("PASS: s=%q, want=%d, got=%d\n", tt.s, tt.want, got)
		} else {
			fmt.Printf("FAIL: s=%q, want=%d, got=%d\n", tt.s, tt.want, got)
		}
	}
	fmt.Println("测试结束。")
}
