package main

import "fmt"


func rotate(nums []int, k int) []int  {
	var newNums []int = make([]int, len(nums), len(nums))
	for i:=0; i < len(nums) - k; i++ {
		newNums[k + 1 + i] = nums[i]
	}
	return newNums
}


func main() {
	var (
		nums = []int{1,2,3,4,5,6,7}
		k = 2
	)
	fmt.Println(rotate(nums, k))
}