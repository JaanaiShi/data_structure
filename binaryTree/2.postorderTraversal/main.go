package main

import (
	"fmt"
	"leetcode/common"
)

func postorderTraversal(root *common.TreeNode) []int {
	return recursion(root)
}

func recursion(root *common.TreeNode) []int {
	var (
		result []int
	)

	if root == nil {
		return result
	}

	if root.Left == nil && root.Right == nil {
		result = append(result, root.Val)
	}else if root.Left != nil && root.Right == nil {
		leftResult := recursion(root.Left)
		root.Left = nil
		result = append(result, leftResult...)
	}else {
		rightResult := recursion(root.Right)
		root.Right = nil
		result = append(result, rightResult...)
	}
	return result
}

func postorderTraversal1(root *common.TreeNode) []int {
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
		if node.Left == nil && node.Right == nil {
			result = append(result, node.Val)
			stack = stack[:len(stack) - 1]
			node = nil
		} else {
			if node.Right != nil {
				stack = append(stack, node.Right)
				node.Right = nil
			}
	
			if node.Left != nil {
				stack = append(stack, node.Left)
				node.Left = nil
			}
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

	fmt.Println(postorderTraversal1(tree))
}
