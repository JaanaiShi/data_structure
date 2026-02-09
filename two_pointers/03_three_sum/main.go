package main

import (
	"fmt"
)

// threeSum 使用排序 + 双指针算法找到所有和为 0 的不重复三元组。
//
// 题目说明：
// 给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，
// 同时还满足 nums[i] + nums[j] + nums[k] == 0 。请你返回所有和为 0 且不重复的三元组。
// 注意：答案中不可以包含重复的三元组。
//
// 示例 1：
// 输入：nums = [-1,0,1,2,-1,-4]
// 输出：[[-1,-1,2],[-1,0,1]]
// 解释：
// nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
// nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
// nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
// 不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
// 注意，输出的顺序和三元组的顺序并不重要。
//
// 示例 2：
// 输入：nums = [0,1,1]
// 输出：[]
// 解释：唯一可能的三元组和不为 0 。
//
// 示例 3：
// 输入：nums = [0,0,0]
// 输出：[[0,0,0]]
// 解释：唯一可能的三元组和为 0 。
func threeSum(nums []int) [][]int {
	// 边界检查
	if len(nums) < 3 {
		return [][]int{}
	}

	// 排序
	sort(nums)
	res := make([][]int, 0)

	for i := 0; i < len(nums)-2; i++ {
		// 剪枝：如果当前数字大于0，后面不可能找到和为0的三元组
		if nums[i] > 0 {
			break
		}

		// ✅ 去重优化1：跳过重复的 nums[i]
		// 如果当前元素与前一个元素相同，跳过避免重复计算
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				// 找到一组解
				res = append(res, []int{nums[i], nums[left], nums[right]})

				// ✅ 去重优化2：跳过重复的 left
				// 找到解后，跳过所有相同的 left 值
				for left < right && nums[left] == nums[left+1] {
					left++
				}

				// ✅ 去重优化3：跳过重复的 right
				// 找到解后，跳过所有相同的 right 值
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				// 同时移动 left 和 right，寻找下一组可能的解
				left++
				right--
			} else if sum > 0 {
				// 和太大，右指针左移
				right--
			} else {
				// 和太小，左指针右移
				left++
			}
		}
	}
	return res
}

func sort(nums []int) {
	// 冒泡排序
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}

func main() {
	tests := []struct {
		nums []int
		want [][]int
		desc string
	}{
		{
			[]int{-1, 0, 1, 2, -1, -4},
			[][]int{{-1, -1, 2}, {-1, 0, 1}},
			"基础测试用例",
		},
		{
			[]int{0, 1, 1},
			[][]int{},
			"无解用例",
		},
		{
			[]int{0, 0, 0},
			[][]int{{0, 0, 0}},
			"全零用例",
		},
		{
			[]int{-2, 0, 0, 2, 2},
			[][]int{{-2, 0, 2}},
			"测试去重：多个重复元素",
		},
		{
			[]int{-1, -1, -1, 2, 2, 2},
			[][]int{{-1, -1, 2}},
			"测试去重：大量重复",
		},
	}

	for i, tt := range tests {
		got := threeSum(tt.nums)
		fmt.Printf("Test %d: %s\n", i+1, tt.desc)
		fmt.Printf("  Input: %v\n", tt.nums)
		fmt.Printf("  Want:  %v\n", tt.want)
		fmt.Printf("  Got:   %v\n", got)

		// 简单判断结果是否正确
		if len(got) == len(tt.want) {
			fmt.Printf("  ✅ 通过\n\n")
		} else {
			fmt.Printf("  ❌ 失败\n\n")
		}
	}
}
