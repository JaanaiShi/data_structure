package main

func reverseString(s []byte)  {
	middleware := len(s) / 2
	left, right := 0, len(s) - 1

	for left < middleware {
		s[left], s[right] = s[right], s[left]

		left++
		right--
	}
}


func main() {
	
}