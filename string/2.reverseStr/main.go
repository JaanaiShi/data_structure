package main

import "fmt"


func reverseStr(s string, k int) string {
	return string(reverse([]byte(s), k))
}

func reverse(sByte []byte, k int) []byte {
	var (
		sByte1, sByte2 []byte
		newByte []byte
	)

	if len(sByte) > 2 * k {
		sByte1 = sByte[:k]
		reverseString(sByte1)   // 没有扩容，所以切片的地址没有改变
		sByte1 = append(sByte1, sByte[k:2*k]...)
		sByte2 = reverse(sByte[2*k:], k)
		newByte = append(sByte1, sByte2...)
	}else if len(sByte) <= 2 * k && len(sByte) > k {
		sByte1 = sByte[:k]
		reverseString(sByte1)
		newByte = append(sByte1, sByte[k:]...)
	}else {
		reverseString(sByte)
		newByte = sByte
	}

	return newByte
}

func reverseString(s []byte)  {
	middleware := len(s) / 2
	left, right := 0, len(s) - 1

	for left < middleware {
		s[left], s[right] = s[right], s[left]

		left++
		right--
	}
}

// 循环
func reverseStr2(s string, k int) string {
	var(
		sByte = []byte(s)
		length = len(sByte)
	)

	for i:=0; i < length; i += 2 * k {
		if length - i > 2 * k {
			reverseString(sByte[i:k])
		}else {
			
		}
	}

	return string(sByte)
}

func main() {	
	/*
		abdd sadf dzcv x
		badd asdf zdcv x
		badd asdf zdcv cxbf egrw fegc
	*/ 
	fmt.Println(reverseStr("abddsadfdzcvx", 2))
}