package main

import (
	"fmt"
	"leetcode/common"
)


func preorderTraversal(root *common.TreeNode) []int {
	return recursion(root)
}

func recursion(root *common.TreeNode) []int {
	var (
		result []int
	)
	if root != nil {
		result = append(result, root.Val)
	}else {
		return result
	}

	leftResult := recursion(root.Left)
	rightResult := recursion(root.Right)

	result = append(result, leftResult...)
	result = append(result, rightResult...)

	return result
}

func preorderTraversal1(root *common.TreeNode) []int {
	var (
		result []int
		stack []*common.TreeNode   
	)
	if root == nil {
		return result
	}
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		result = append(result, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return result
}

func main() {
	var (
		tree = &common.TreeNode{
			Val: 1,
			Left: &common.TreeNode{
				Val: 2,
				Left: &common.TreeNode{
					Val: 4,
				},
				Right: &common.TreeNode{
					Val: 5,
				},
			},
			Right: &common.TreeNode{
				Val: 3,
			},
		}
	)

	fmt.Println(preorderTraversal1(tree))
}
