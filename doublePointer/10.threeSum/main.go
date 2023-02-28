package main

import "fmt"

func threeSum(nums []int) [][]int {
	var (
		leftIndex  int = 0
		result                [][]int
	)
	// 1. 先进行排序
	sort(nums)

	// 面对这样的题，有了思路之后，在写代码的时候，要考虑各种条件，比如下面这道题就是需要考虑退出条件
	for ; leftIndex < len(nums); leftIndex++ {

		if leftIndex > 0 && nums[leftIndex - 1] == nums[leftIndex] {
			continue
		}

		i, rightIndex := leftIndex + 1, len(nums) - 1
		for i < rightIndex {
			sum := nums[leftIndex] + nums[i] + nums[rightIndex]
			if sum == 0 {
				result = append(result, []int{
					nums[leftIndex], nums[i], nums[rightIndex],
				})

				// 去重逻辑需要放在找到第一个三元组之后，对b，c去重
				for i < rightIndex && nums[i] == nums[i+1] {
					i++
				}

				for i < rightIndex && nums[rightIndex] == nums[rightIndex - 1] {
					rightIndex--
				}

				// 找到答案时，双指针同时收缩
				i++
				rightIndex--
			}else if sum > 0 {
				rightIndex--
			}else {
				i++
			}
		}
	}

	return result

}

func sort(nums []int) {

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i; j++ {
			if j+1 < len(nums)-i {
				if nums[j] > nums[j+1] {
					nums[j], nums[j+1] = nums[j+1], nums[j]
				}
			}
		}
	}
}


func main() {
	nums := []int{-1,0,1,2,-1,-4,-2,-3,3,0,4}

	fmt.Println(threeSum(nums))
}
