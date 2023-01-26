package main

import "fmt"

func replaceSpace(s string) string {
	sByte := []byte(s)
	newByte := []byte{}
	replace := []byte{37,50,48}

	for _, v := range sByte {
		if v == 32 {
			newByte = append(newByte, replace...)
		}else {
			newByte = append(newByte, v)
		}
	}

	return string(newByte)
}

// 原地修改
func replaceSpace1(s string) string {
    b := []byte(s)
    length := len(b)
    spaceCount := 0
    // 计算空格数量
    for _, v := range b {
        if v == ' ' {
            spaceCount++
        }
    }
    // 扩展原有切片
    resizeCount := spaceCount * 2
    tmp := make([]byte, resizeCount)
    b = append(b, tmp...)
    i := length - 1
    j := len(b) - 1
    for i >= 0 {
        if b[i] != ' ' {
            b[j] = b[i]
            i--
            j--
        } else {
            b[j] = '0'
            b[j-1] = '2'
            b[j-2] = '%'
            i--
            j = j - 3
        }
    }
    return string(b)
}

func main() {
	fmt.Println(replaceSpace("We are happy."))
}