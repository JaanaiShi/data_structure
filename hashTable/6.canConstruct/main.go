package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	var (
		magazineMap = make(map[rune]int)
	)

	for _, v := range magazine {
		magazineMap[rune(v) -97]++
	}

	for _, v := range ransomNote {
		ran := rune(v) - 97
		if value, ok := magazineMap[ran]; ok {
			magazineMap[ran] = value - 1
			if magazineMap[ran] < 0 {
				return false
			}
		}else {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(canConstruct("aab", "aassdsadfasdfsafb"))

}