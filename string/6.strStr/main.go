package main

import "fmt"

// 暴力破解
func strStr(haystack string, needle string) int {
	var (
		flag bool
	)
	h := []byte(haystack)
	n := []byte(needle)

	for i:=0; i < len(h); i++ {
		if h[i] == n[0] {
			num := i
			for j:=0; j < len(n); j++ {
				if num < len(h) && h[num] == n[j] {
					flag = true
				}else {
					flag = false
					break
				}
				num++
			}
			if flag {
				return i
			}
		}
	}

	return -1
}

// KMP算法实现（未完成）
// func strStr1(haystack string, needle string) int {

// }

/*
	abcdabx
	a
	ab
	abc
	abcd
	abcda
	abcdab

*/
// 构建next数组
func getNext(s string) []int {
	/*
		定义了两个指针i和j，j指向前缀末尾位置，i指向后缀末尾位置
	*/
	var (
		next = make([]int, len(s))
	)
	j := 0
	next[0] = j
	for i := 1; i < len(s); i++ {
		for j > 0 && s[i] != s[j] {  // 前后缀不相同; j要保证大于0，因为下面有j-1的操作
			j = next[j-1]   // 向前回退
		}
		if s[i] == s[j] { // 找到相等的前后缀
			j++
		}
		next[i] = j  // 将j（前缀的长度）赋给next[i]
	}
	return next
}

func strStr1(haystack string, needle string) int {
	// 1. 获取next数组
	next := getNext(needle)

	// 2. 获取子串的位置
	for i:=0; i < len(haystack); i++ {
		current := i
		j, result := 0, i
		if haystack[i] == needle[j] {
			for {
				if j > 0 && haystack[current] != needle[j] {
					j = next[j-1]
					result += j 
				}else if j == len(needle) - 1 {
					return result 
				}else {
					j++
					current++
				}
				if j == 0 || current >= len(haystack) {
					break
				}
			}
		}
		
	}

	return -1
}

func strStr2(haystack string, needle string) int {
	n := len(needle)

	if n == 0 {
		return 0
	}

	j := 0
	next := getNext(needle)

	for i:=0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j - 1]
		}

		if haystack[i] == needle[j] {
			j++
		}
		
		if j == n {
			return i - n + 1
		}
	}

	return -1
}



func main(){
	// fmt.Println(strStr("leetcode", "leeto"))
	fmt.Println(strStr2("mississippi","issip"))
}