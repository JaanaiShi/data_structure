package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	var (
		slowIndex int
	)
	// 去掉头部和尾部的空格
	s = strings.TrimSpace(s)
	sByte := []byte(s)
	// 去掉中间的空格
	//删除单词间冗余空格
	for i:=0; i < len(sByte); i++ {
		if i-1 > 0 && sByte[i-1] == sByte[i] && sByte[i] == ' ' {
			continue
		}
		sByte[slowIndex] = sByte[i]
		slowIndex++
	}
	sByte = sByte[:slowIndex]
	slowIndex = 0
	sByte = append(sByte, ' ') // 为了可以反转最后一个单词
	for i:=0; i < len(sByte); i++ {
		if sByte[i] == ' ' {
			reverse(sByte, slowIndex, i - 1)
			slowIndex = i + 1
		}
	}

	

	reverse(sByte, 0, len(sByte) - 1)


	return string(sByte[1:])
}

func reverse(s []byte, left, right int) {
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
	
}


func main() {
	fmt.Println(reverseWords(" hello    world     "))
}