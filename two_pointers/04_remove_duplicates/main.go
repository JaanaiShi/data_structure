package main

import (
	"fmt"
)

// removeDuplicates 使用快慢指针在原地移除有序数组中的重复项。
// 返回移除后数组的新长度。
//
// 题目说明：
// 给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。
// 元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。
// 考虑 nums 的唯一元素的数量为 k ，你需要做以下事情确保你的题解可以被通过：
// 1. 更改数组 nums ，使 nums 的前 k 个元素包含唯一元素，并按照它们最初在 nums 中出现的顺序排列。nums 的其余元素与 nums 的大小不重要。
// 2. 返回 k 。
//
// 示例 1：
// 输入：nums = [1,1,2]
// 输出：2, nums = [1,2,_]
// 解释：函数应该返回新的长度 2 ，并且原数组 nums的前两个元素被修改为 1, 2 。不需要考虑数组中超出新长度后面的元素。
//
// 示例 2：
// 输入：nums = [0,0,1,1,1,2,2,3,3,4]
// 输出：5, nums = [0,1,2,3,4]
// 解释：函数应该返回新的长度 5 ， 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4 。
func removeDuplicates(nums []int) int {
	// TODO: 请实现快慢指针算法
	return 0
}

func main() {
	tests := []struct {
		nums    []int
		want    int   // 期望长度
		wantArr []int // 期望的前缀数组
	}{
		{[]int{1, 1, 2}, 2, []int{1, 2}},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5, []int{0, 1, 2, 3, 4}},
	}

	fmt.Println("开始测试 04_remove_duplicates...")
	for _, tt := range tests {
		// 为了不修改测试用例原始数据，这里拷贝一份
		input := make([]int, len(tt.nums))
		copy(input, tt.nums)

		got := removeDuplicates(input)

		pass := true
		if got != tt.want {
			pass = false
		} else {
			// 检查前 got 个元素是否匹配
			for i := 0; i < got; i++ {
				if input[i] != tt.wantArr[i] {
					pass = false
					break
				}
			}
		}

		if pass {
			fmt.Printf("PASS: nums=%v, wantLen=%d, gotLen=%d, result=%v\n", tt.nums, tt.want, got, input[:got])
		} else {
			fmt.Printf("FAIL: nums=%v, wantLen=%d, gotLen=%d, result=%v\n", tt.nums, tt.want, got, input[:got])
		}
	}
	fmt.Println("测试结束。")
}
