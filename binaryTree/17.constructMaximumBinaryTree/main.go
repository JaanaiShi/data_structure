package main

import "leetcode/common"


func constructMaximumBinaryTree(nums []int) *common.TreeNode {

	if len(nums) == 0 {
		return nil
	}

	num, index := getSliceMax(nums)

	node := common.TreeNode{Val: num}
	leftNum := nums[:index]
	rightNum := []int{}
	node.Left = constructMaximumBinaryTree(leftNum)
	if index == len(nums) {
		rightNum = []int{}
	}else {
		rightNum = nums[index+1:]
	}
	node.Right = constructMaximumBinaryTree(rightNum)

	return &node
}


func getSliceMax(nums []int) (num, index int) {
	for i:=0; i < len(nums); i++ {
		if nums[i] > num {
			num = nums[i]
			index = i
		}
	}
	return
}

func main() {
	nums := []int{7,10,6,3,9,8}
	tree := constructMaximumBinaryTree(nums)

	common.PreOrder(tree)
}