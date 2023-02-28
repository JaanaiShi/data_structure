package main

import "fmt"

func threeSum(nums []int) [][]int {
	var (
		result [][]int
	)

	for index, num := range nums {
		if len(nums) - index > 0 {
			towSumResult := twoSum(nums[index+1:], 0-num)
			for i:=0; i < len(towSumResult); i++ {
				result = append(result, append(towSumResult[i], num))
			}
		}
	}
	return result
}

func twoSum(nums []int, target int) [][]int {
	var (
		result [][]int
		numsMap = make(map[int]struct{})
	)

	for i:=0; i < len(nums); i++ {
		diff := target - nums[i]
		if _, ok := numsMap[diff]; ok {
			result = append(result, []int{
				diff,
				nums[i],
			})
		}else {
			numsMap[nums[i]] = struct{}{}
		}
	}

	return result
}

func main() {
	fmt.Println(threeSum([]int{-1,0,1,2,-1,-4}))
}