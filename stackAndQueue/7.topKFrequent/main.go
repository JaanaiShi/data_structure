package main


func topKFrequent(nums []int, k int) []int {
	var (
		result []int
		numsMap  = make(map[int]int, 0)
	)

	for i:=0; i < len(nums); i++ {
		if v, ok := numsMap[nums[i]]; ok {
			numsMap[nums[i]] = v + 1
		}else {
			numsMap[nums[i]] = 1
		}
	}

	for len(result) < k {
		max, maxLength := 0, 0
		
		for k, v := range numsMap {
			if v > maxLength {
				max = k
				maxLength = v
			}
		}

		delete(numsMap, max)

		result = append(result, max)
	}
	return result
}

func main() {

}
                       