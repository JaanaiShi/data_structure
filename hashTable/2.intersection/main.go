package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	var (
		num1Map = make(map[int]struct{}, 0)
		num2Map = make(map[int]struct{}, 0)
		result  []int
	)

	for i := 0; i < len(nums1); i++ {
		num1Map[nums1[i]] = struct{}{}
	}

	for i := 0; i < len(nums2); i++ {
		num2Map[nums2[i]] = struct{}{}
	}

	for k, _ := range num1Map {
		if _, ok := num2Map[k]; ok {
			result = append(result, k)
		}
	}
	return result
}

func main() {
	var (
		nums1 = []int{1,2,2,1,3,3,3,34,4,1}
		nums2 = []int{2,2,1,34}
	)

	fmt.Println(intersection(nums1, nums2))
}
