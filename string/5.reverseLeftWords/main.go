package main

import (
	"fmt"
)

func reverseLeftWords(s string, n int) string {
	
	b := []byte(s)
	n = n - 1
	length := len(b)
	for i:=0; i < length; i++ {
		if i <= n {
			b = append(b, b[i])
		} 
	}

	b = b[n+1:]

	return string(b)

}

func reverseLeftWords1(s string, n int) string {
	
	b := []byte(s)

	reverse(&b, 0, n-1)
	reverse(&b, n, len(b) -1)
	reverse(&b, 0, len(b) - 1)


	return string(b)

}

func reverse(b *[]byte, left, right int) {

	for left <= right {
		(*b)[left], (*b)[right] = (*b)[right], (*b)[left]
		left++
		right--
	}
}

func main() {
	s  := "abcdefg"

	fmt.Println(reverseLeftWords1(s, 2))

	// b := []byte("abcdefg")
	// reverse(&b, 0, 1)
	// fmt.Println(string(b))
}