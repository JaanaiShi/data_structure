package main

import "fmt"

func reverseString(s []byte)  {
	var (
		left, right  int
	)

	left, right = 0, len(s) - 1

	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

func main() {
	s  := []byte{'h','e','l','l','o'}
	fmt.Println(string(s))

	reverseString(s)

	fmt.Println(string(s))
}