package main

import "fmt"

/*
给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按
非递减顺序 排序。
输入：nums = [-4,-1,0,3,10]
输出：[0,1,9,16,100]
解释：平方后，数组变为 [16,1,0,9,100]
排序后，数组变为 [0,1,9,16,100]

输入：nums = [-7,-3,2,3,11]
输出：[4,9,9,49,121]
*/

/*
	2, 5,3,4,1

	2, 3, 5, 4, 1



*/

// 双指针写法: 适合于所有的比大小
func pointer(nums []int) []int {
	var (
		left, right int
		result =make([]int, 0)
	)
	if len(nums) == 0 {
		return []int{}
	}

	left, right = 0, len(nums) - 1

	for {
		if left > right {
			break
		}

		squareLeft := nums[left] * nums[left]
		squareRight := nums[right] * nums[right]

		if squareLeft > squareRight {
			result = append([]int{squareLeft}, result...)
			left++
		}else if squareLeft < squareRight {
			result = append([]int{squareRight}, result...)
			right--
		} else {
			result = append([]int{squareLeft}, result...)
			left++
		}
	} 

	return result

}

func sortedSquares(nums []int) []int {
	if len(nums) == 1 {
		nums[0] = nums[0] * nums[0]
		return nums
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = nums[i] * nums[i]
	}

	// 冒泡排序
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

	return nums
}

func main() {
	nums := []int{-12, 2, 3, 5, 9}
	fmt.Println("nums:", nums)
	fmt.Println("result:", pointer(nums))
}
