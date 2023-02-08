package main

import "fmt"
// 此种解法是以快指针为出发点的，而要以慢指针为出发点的话，则要复杂的多
func removeElement(nums []int, val int) int {
	var (
		length = len(nums)
		slowIndex = 0
	)

	for i:=0; i < length; i++ {
		if nums[i] != val {
			nums[slowIndex] = nums[i]  // 当slowIndex没有指到val的时候，则slowIndex与i相同
			slowIndex++
		}
	}

	return len(nums[:slowIndex])
}

func main() {
	var (
		nums = []int{0,1,2,2,3,0,4,2}
		val  = 2
	)

	fmt.Println(removeElement(nums, val))

	fmt.Println(nums)
}	