package main

import (
	"fmt"
)

var (
	nums = []int{-1,0,3,5,9,12} 
	target = 13
)

func search() int {
	var (
        medium,left int
		right = len(nums) - 1
	)
    if len(nums) == 0 {
        return 0
    }
	for (left <= right) {
        medium = left + (right - left) / 2
		if nums[medium] == target {
			return medium
		}else if nums[medium] > target {
			right = medium - 1
		}else {
			left = medium + 1
		}
	}

    return -1

}


func main() {
	fmt.Println(search())
}