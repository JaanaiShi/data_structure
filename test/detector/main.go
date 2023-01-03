package main

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

// 装饰器
type util func() error

func TimeSince(f util) util {

	defer func(t time.Time) {
		fmt.Println("函数执行时间: ", runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), 
		time.Since(t))
	}(time.Now())

	return f
} 

func Println() error {
	fmt.Println("打印函数")
	return nil
}

func Sum() error {
	fmt.Println("加和函数")
	return nil
}

func main() {
	// 这样写会更好一点
	println := TimeSince(Println)
	println()
	sum := TimeSince(Sum)
	sum()
}
