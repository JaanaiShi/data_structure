package main

func twoSum(nums []int, target int) []int {
	var (
		result []int
	)

	if len(nums) <= 1 {
		return result
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result = append(result, i)
				result = append(result, j)
				return result
			}
		}
	}

	return result
}

func twoSum2(nums []int, target int) []int {
	var (
		result []int
		numMap = make(map[int]int, 0)
	)

	for i:=0; i < len(nums); i++ {
		diff := target - nums[i]
		if v, ok := numMap[diff]; ok {
			result = append(result, i)
			result = append(result, v)
			return result
		}else {
			numMap[nums[i]] = i
		}
	}

	return result
}

func main() {

}
