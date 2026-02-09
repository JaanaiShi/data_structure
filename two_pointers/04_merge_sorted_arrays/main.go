package main

import (
	"fmt"
	"reflect"
)

// mergeSimple 合并两个有序数组，返回一个新的有序数组。
// 不使用内建排序，使用双指针算法。
//
// 题目要求：
// 合并两个有序数组（不允许使用内建排序）
func mergeSimple(nums1 []int, nums2 []int) []int {
	pre, cur := 0, 0
	res := make([]int, 0, len(nums1))
	for pre < len(nums1) && cur < len(nums2) {
		if nums1[pre] <= nums2[cur] {
			res = append(res, nums1[pre])
			pre++
		} else {
			res = append(res, nums2[cur])
			cur++
		}
	}
	if pre < len(nums1) {
		res = append(res, nums1[pre:]...)
	}
	if cur < len(nums2) {
		res = append(res, nums2[cur:]...)
	}
	return res
}

func main() {
	tests := []struct {
		nums1 []int
		nums2 []int
		want  []int
	}{
		{[]int{1, 3, 5}, []int{2, 4, 6}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{1, 2, 3}, []int{}, []int{1, 2, 3}},
		{[]int{}, []int{4, 5, 6}, []int{4, 5, 6}},
		{[]int{1, 1, 2}, []int{1, 3, 4}, []int{1, 1, 1, 2, 3, 4}},
		{[]int{10}, []int{1, 2, 3}, []int{1, 2, 3, 10}},
	}

	fmt.Println("开始测试 two_pointers/04_merge_sorted_arrays...")
	for _, tt := range tests {
		got := mergeSimple(tt.nums1, tt.nums2)
		if reflect.DeepEqual(got, tt.want) {
			fmt.Printf("PASS: nums1=%v, nums2=%v, want=%v, got=%v\n", tt.nums1, tt.nums2, tt.want, got)
		} else {
			fmt.Printf("FAIL: nums1=%v, nums2=%v, want=%v, got=%v\n", tt.nums1, tt.nums2, tt.want, got)
		}
	}
	fmt.Println("测试结束。")
}
