package main

import (
	"fmt"
	myqueue "leetcode/stackAndQueue/6.maxSlidingWindow/myQueue"
)

func maxSlidingWindow(nums []int, k int) []int {
	var (
		result []int
		i   int
		myQueue = myqueue.NewMyQueue()
	)


	for i < k {
		myQueue.Push(nums[i])
		i++
	}

	result = append(result, myQueue.Front())

	for ; i < len(nums); i++ {
			myQueue.Pop(nums[i-k])

			myQueue.Push(nums[i])

			result = append(result, myQueue.Front())
	}

	

	return result
}

func main() {
	nums := []int{5,5,5,5}
	fmt.Println(maxSlidingWindow(nums, 3))
}
