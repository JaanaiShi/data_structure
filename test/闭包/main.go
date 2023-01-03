package main

import (
	"fmt"
	"time"
)


func CalTimeCost(f func(s string) string, s string) {
	start := time.Now()
	result := f(s)
	fmt.Println(result)
	end := time.Now()
	fmt.Println("cost time", end.Sub(start))
}


func MyFunc(s string) string {
	time.Sleep(2 * time.Second)

	return s + "_111"
}


func main() {
	CalTimeCost(MyFunc, "123")
}



