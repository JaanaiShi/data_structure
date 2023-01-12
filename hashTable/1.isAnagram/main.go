package main

import "fmt"

// 根据map去实现
func isAnagram(s string, t string) bool {
	var (
		sMap = make(map[byte]int, 0)
		tMap = make(map[byte]int, 0)
		
	)

	for i:=0; i < len(s); i++ {
		fmt.Println(s[i])
		if v, ok := sMap[s[i]]; ok {
			sMap[s[i]] = v + 1
		}else {
			sMap[s[i]] = 1
		}
	}

	for i:=0; i < len(t); i++ {
		if v, ok := tMap[t[i]]; ok {
			tMap[t[i]] = v + 1
		}else {
			tMap[t[i]] = 1
		}
	}

	for k, v := range tMap {
		if sv, ok := sMap[k]; ok {
			if sv != v {
				return false
			}
		}else {
			return false
		}
	}

	for k, v := range sMap {
		if tv, ok := tMap[k]; ok {
			if tv != v {
				return false
			}
		}else {
			return false
		}
	}

	return true
}

// 因为题目中的字符串都是英文字母，因此可以采用一个record表，枚举s的过程负责添加，枚举t负责删除，如果record表
// 最终为空的话，则证明互为异位词

func isAnagram1(s string, t string) bool {
	var (
		record = [26]int{}
	)

	// 添加record的值
	for _, v := range s {
		i := v - rune('a')
		record[i] +=  1
	}

	for _, v := range t {
		i := v - rune('a')
		record[i] -= 1
	}

	return record == [26]int{}
} 





func main() {
	var (
		example1 = [2]string{"anagram", "nagaram"}
		example2 = [2]string{"rat", "cat"}
		example3 = [2]string{"ab", "b"}

	)

	fmt.Println(isAnagram1(example1[0], example1[1]))
	fmt.Println(isAnagram1(example2[0], example2[1]))
	fmt.Println(isAnagram1(example3[0], example3[1]))


}