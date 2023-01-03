package main

import "fmt"

/*
1. 双指针法（快慢指针法）：通过一个快指针和慢指针在一个for循环下完成两个for循环的工作
定义快慢指针：
 * 快指针：寻找新数组的元素，新数组就是不含有目标元素的数组
 * 慢指针：指向更新 新数组下标的位置
*/


func removeElement(nums []int, val int) int {
	var count = 0
	a := 0
	for i:= 0; i < len(nums); i++ {
		for j:=0; j < len(nums) - i - 1; j++ {
			if nums[j] == val {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}		
		}
	}
	
	for i:= 0; i < len(nums); i++ {
		if nums[i] == val {
			a = i
			count = len(nums) -i
			break
		}
	}
	if a != 0 {
		nums = nums[:a]
	}
	fmt.Println(nums)
	

	

	return count
	
}
// removeElement1
func removeElement1(nums []int, val int) int {
	length := len(nums)
	res := 0
	for i:=0; i< length; i++ {
		if nums[i] != val {
			nums[res] = nums[i]
			res++
		}
	}
	nums = nums[:res]

	return res
	
}

func MySort(arr []int) {
    for i := 0; i < len(arr); i++ { // 5 // 控制循环几次 每次都要从头开始，所以i,j 必须为0
        for j := 0; j < len(arr)-i-1; j++ { // 4 控制排好一次需要几次,减去i，是因为i从1开始表示有多少已经排好了
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
        // break
    }

}

func main() {
	nums := []int{0,1,2,2,3,0,4,2}
	val := 2
	fmt.Println(removeElement(nums, val))
}