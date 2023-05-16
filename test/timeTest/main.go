package main

import (
	"fmt"
	"time"
)



func main() {
	now := "2023-03-22T12:18:05+08:00"
	layout := time.DateOnly
	time, err := time.Parse(layout, now, )
	if err != nil {
		fmt.Println("err: ", err)
	}

	fmt.Println(time)
}