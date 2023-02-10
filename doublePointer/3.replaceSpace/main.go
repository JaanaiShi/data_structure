package main

import "fmt"


func replaceSpace(s string) string {
	var (
		sByte = []byte(s)
		slowIndex = 0
	)

	length := len(sByte)

	for i:=0; i < length; i++ {
		if sByte[i] == ' ' {
			temp := make([]byte, 2)
			sByte = append(sByte, temp...)
		}
	}

	length = len(sByte)
	slowIndex = length - 1
	for i:=length-1; i >= 0; i-- {
		if sByte[i] != 0 && sByte[i] != ' ' {
			sByte[i], sByte[slowIndex] = sByte[slowIndex], sByte[i]
			slowIndex--
		}else if sByte[i] == ' ' {
			temp := []byte("02%")
			for _, v := range temp {
				sByte[slowIndex] = v
				slowIndex--
			}
		}
	}

	return string(sByte)
}

func main() {
	fmt.Println(replaceSpace("  "))
}