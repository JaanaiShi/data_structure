package main

import (
	"fmt"
	"strconv"
)

func isHappy(n int) bool {
	var (
		// result *bool   // 此时只是声名了一个指针，但是没有初始化
		result = new(bool)   // 使用new创建指针，并进行分配内存
	)

	recursion(n, result)

	return *result
}

func recursion(n int, result *bool) {

	var (
		sum int
	)

	str := strconv.Itoa(n)

	for _, v := range str {
		p := int(v - 48)
		sum += p * p
	}

	if sum  >= n {
		return
	}

	if sum == 1 {
		*result = true
		return
	}

	recursion(sum, result)
}

func isHappy2(n int) bool {
	var (
		nMap = make(map[int]bool, 0)
	)

	for n != 1 && !nMap[n] {
		n, nMap[n] = getSum(n), true
	}

	return n == 1
}

func getSum(n int) int {
	var (
		sum int
	)
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}

	return sum
}

func main()  {
	fmt.Println(isHappy(2))
}