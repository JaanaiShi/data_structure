package main

import (
	"fmt"
	"testing"
)

/*
给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。

输入: nums = [-1,0,3,5,9,12], target = 9
输出: 4
解释: 9 出现在 nums 中并且下标为 4

输入: nums = [-1,0,3,5,9,12], target = 2
输出: -1
解释: 2 不存在 nums 中因此返回 -1
*/

/*
	解题思路：
	首先理解题意：该数组是一个有序的，且是整型 数组，考虑到升序我们首先想到的二分法（用于查找的算法）

	这时我们需要想到二分法的实现过程。
*/
func search(nums []int, target int) int {
	var (
		left, right = 0, len(nums) - 1
		result = -1
	)

	for left <= right {
		middle := (right - left) / 2 + left
		if nums[middle] == target {
			result = middle
			break
		} else if nums[middle] > target {
			right--
		} else {
			left++
		}
	}

	return result
}


func TestSearch(* testing.T) {
	var (
		nums = []int{-1,0,3,5,9,12}
		target = 2
	)

	fmt.Println(search(nums, target))
}

/*
	给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。
元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

说明:

为什么返回数值是整数，但输出的答案是数组呢?
请注意，输入数组是以「引用」方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。
你可以想象内部操作如下:

// nums 是以“引用”方式传递的。也就是说，不对实参作任何拷贝
int len = removeElement(nums, val);

// 在函数里修改输入数组对于调用者是可见的。
// 根据你的函数返回的长度, 它会打印出数组中 该长度范围内 的所有元素。
for (int i = 0; i < len; i++) {
    print(nums[i]);
}

输入：nums = [3,2,2,3], val = 3
输出：2, nums = [2,2]
解释：函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。你不需要考虑数组中超出新长度后面的元素。例如，函数返回的新长度为 2 ，而 nums = [2,2,3,3] 或 nums = [2,2,0,0]，也会被视作正确答案。

输入：nums = [0,1,2,2,3,0,4,2], val = 2
输出：5, nums = [0,1,4,0,3]
解释：函数应该返回新的长度 5, 并且 nums 中的前五个元素为 0, 1, 3, 0, 4。注意这五个元素可为任意顺序。你不需要考虑数组中超出新长度后面的元素。
*/

func removeElement(nums []int, val int) int {
	// 使用双指针的做法
	var (
		result int
	)

	for i:=0; i < len(nums); i++ {
		if nums[i] == val {
			for j:=i; j < len(nums); j++ {
				if nums[j] != val {
					nums[i], nums[j] = nums[j], nums[i]
					break
				}
			}
		}
	}

	for i:=0; i < len(nums); i++ {
		if nums[i] == val {
			result = i
			break
		}
	}

	if result == 0 && len(nums) == 1 && nums[0] == val{
		result = len(nums)
	}

	return result
}


func removeElement1(nums []int, val int) int {
	// 使用双指针的做法, 前一个指针寻找不等于2的, 后一个指针找2
	var (
		resultArray []int

		j int
	)

	for i:=0; i < len(nums); i++ {

		if nums[j] == val && nums[i] != val {
			nums[j], nums[i] = nums[i], nums[j]
			j++
		}else if nums[j] == val {

		}else {
			j = i
		}
	}

	resultArray = nums[:j]

	return len(resultArray)
}

func removeElement2(nums []int, val int) int {
	// 使用双指针的做法, 位置前移
	var (
		resultArray []int

		j int
	)

	for i:=0; i < len(nums); i++ {
		if nums[j] == val && nums[i] != val {
			nums[j] = nums[i]
			j++
		} else if nums[i] != val {
			j = i
		}
	}

	resultArray = nums[:j]

	return len(resultArray)
}

func removeElement3(nums []int, val int) int {
	// 使用双指针的做法
	var (
		resultArray []int

		slowIndex int
	)

	for fastIndex:=0; fastIndex < len(nums); fastIndex++ {
		if val != nums[fastIndex] {
			nums[slowIndex] = nums[fastIndex]
			slowIndex++
		}
	}

	resultArray = nums[:slowIndex]

	return len(resultArray)
}

func TestRemoveElement(* testing.T) {
	var (
		nums = []int{1}
		target = 2
	)

	fmt.Println(removeElement3(nums, target))
}


/*
给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。

输入：nums = [-4,-1,0,3,10]
输出：[0,1,9,16,100]
解释：平方后，数组变为 [16,1,0,9,100]
排序后，数组变为 [0,1,9,16,100]

输入：nums = [-7,-3,2,3,11]
输出：[4,9,9,49,121]
*/

func sortedSquares(nums []int) []int {
	// 算法思想为，前后双指针，比较大小，采用头插法

	var (
		result []int
		left, right = 0, len(nums) - 1
	)

	for left <= right {
		leftSquare := nums[left] * nums[left]
		rightSquare := nums[right] * nums[right]
		if leftSquare <= rightSquare {
			result = append([]int{rightSquare}, result...)
			right--
		} else {
			result = append([]int{leftSquare}, result...)
			left++
		}
	}

	return result
}

func TestSortedSquares(t *testing.T) {
	var (
		nums = []int{-4, -1, 0, 3, 10}
	)

	fmt.Println(sortedSquares(nums))
}




