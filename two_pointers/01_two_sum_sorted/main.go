package main

import (
	"fmt"
	"reflect"
)

// twoSum 使用双指针算法在有序数组中查找两个数，使它们的和等于目标值。
// 题目要求返回的下标是从 1 开始的。
//
// 题目说明：
// 给定一个下标从 1 开始的整数数组 numbers ，该数组已按 非递减顺序排列 ，如果两个数相加等于 target ，则返回这两个数的下标。
// 假设每个输入只对应唯一的答案 ，而且你 不可以 重复使用相同的元素。
// 你所设计的解决方案必须只使用常量级的额外空间。
//
// 示例 1：
// 输入：numbers = [2,7,11,15], target = 9
// 输出：[1,2]
// 解释：2 与 7 之和等于目标数 9 。因此 index1 = 1, index2 = 2 。
//
// 示例 2：
// 输入：numbers = [2,3,4], target = 6
// 输出：[1,3]
// 解释：2 与 4 之和等于目标数 6 。因此 index1 = 1, index2 = 3 。
//
// 示例 3：
// 输入：numbers = [-1,0], target = -1
// 输出：[1,2]
// 解释：-1 与 0 之和等于目标数 -1 。因此 index1 = 1, index2 = 2 。
func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		if numbers[left]+numbers[right] == target {
			break
		} else if numbers[left]+numbers[right] > target {
			right--
		} else {
			left++
		}
	}
	return []int{left + 1, right + 1}
}

func main() {
	// 测试用例
	tests := []struct {
		numbers []int
		target  int
		want    []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
		{[]int{2, 3, 4}, 6, []int{1, 3}},
		{[]int{-1, 0}, -1, []int{1, 2}},
		{[]int{2, 7, 11, 15}, 18, []int{2, 3}},
	}

	fmt.Println("开始测试 01_two_sum_sorted...")
	for _, tt := range tests {
		got := twoSum(tt.numbers, tt.target)
		if reflect.DeepEqual(got, tt.want) {
			fmt.Printf("PASS: numbers=%v, target=%d, want=%v, got=%v\n", tt.numbers, tt.target, tt.want, got)
		} else {
			fmt.Printf("FAIL: numbers=%v, target=%d, want=%v, got=%v\n", tt.numbers, tt.target, tt.want, got)
		}
	}
	fmt.Println("测试结束。")
}
