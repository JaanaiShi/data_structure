package main

import "fmt"

/*
给定一个含有 n 个正整数的数组和一个正整数 target 。

找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。

示例 1：
输入：target = 7, nums = [2,3,1,2,4,3]
输出：2
解释：子数组 [4,3] 是该条件下的长度最小的子数组。

示例 2：
输入：target = 4, nums = [1,4,4]
输出：1
*/

// minSubArrayLen 做法为暴力解法
func minSubArrayLen(target int, nums []int) int {
	var (
		result = -1
	)
	for i := 0; i < len(nums); i++ {
		sum := 0
		count := 0
		resultArray := []int{}
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			count++
			resultArray = append(resultArray, nums[j])
			if sum >= target {
				if result == -1 {
					result = count
				}
				if count < result {
					result = count
					break
				}
			}
		}
		fmt.Println(resultArray)

	}
	if result == -1 {
		return 0
	}
	return result
}

// minSubArrayLen1 做法为滑动窗口
/*
	使用一个for循环，此种写法在leetcode上执行还是会超时
*/
func minSubArrayLen1(target int, nums []int) int {
	var (
		result,  left int
	)
	// i为滑动窗口的终止位置
	for i:=0; i< len(nums); i++ {
		// 此循环是让窗口的和再次小于target
		for {
			var sum  = 0
			// 取窗口的和
			temp := left
			for temp <= i {
				sum += nums[temp]
				temp++
			}
			if sum < target {
				break
			}else {
				fmt.Printf("sum:%d,left:%d, right:%d\n", sum, left, i)
				if result > (i - left + 1) || result == 0 {
					result = i - left + 1
				}
				left++
			}
		}
	}

	return result
	
}

func MinSubArrayLen2(target int, nums []int) int {
	var (
		result, sum, left int
	)
	// i为滑动窗口的终止位置
	for i:=0; i< len(nums); i++ {
		// 此循环是让窗口的和再次小于target
		sum += nums[i]
		tag := false

		for sum >= target {
			tag = true
			sum -= nums[left]
			left++
		}
		if tag {
			if result == 0 {
				result = i - left + 2
			} else if result > (i - left + 2) {
				result = i - left + 2
			}
		}
	}

	return result
	
}

func main() {
	var (
		target = 1000
		nums   = []int{5,1,3,5,10,7,4,9,2,8}
	)
	fmt.Println(MinSubArrayLen2(target, nums))

}
