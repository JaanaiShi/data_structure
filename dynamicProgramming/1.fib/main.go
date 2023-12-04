package main

import "fmt"

/*
1. 确定dp数组（dp table）以及下标的含义
2. 确定递推公式
3. dp数组如何初始化
4. 确定遍历顺序
5. 举例推导dp数组
*/

func fib(n int ) int {
	if n <= 1{
		return n
	}
	f1, f2 := 0, 1
	for i:=2; i <= n; i++ {
		fmt.Printf("f1:%d f2:%d\n", f1, f2)
		temp := f1
		f1 = f2
		f2 = temp + f2
		
	}

	return f2
}


func main() {
	fmt.Println(fib(10))
}