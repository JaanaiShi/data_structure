package main

import "fmt"

func sort(nums []int) {

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}

func fourSum(nums []int, target int) [][]int {
	var (
		result [][]int
	)
	// 排序
	sort(nums)

	for i := 0; i < len(nums); i++ {

		// i去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j := i + 1
		for ; j < len(nums); j++ {

			// j去重
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			left := j + 1
			right := len(nums) - 1

			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})

					// 过滤left 和right
					for left < right && nums[left] == nums[left+1] {
						left++
					}

					for left < right && nums[right] == nums[right-1] {
						right--
					}

					left++
					right--
				} else if sum > target {
					right--
				} else {
					left++
				}
			}
		}
	}

	return result
}

func main() {
	nums := []int{2,2,2,2,2}

	fmt.Println(fourSum(nums, 8))
}
