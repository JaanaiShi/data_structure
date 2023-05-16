package main

import (
	"fmt"
	"leetcode/common"
)

func inorderTraversal(root *common.TreeNode) []int {
	return recursion(root)
}


func recursion(root *common.TreeNode) []int {
	var (
		result []int
	)

	if root == nil {
		return result
	}

	leftResult := recursion(root.Left)
	root.Left = nil
	result = append(result, leftResult...)

	if root.Left == nil {
		result = append(result, root.Val)
	}

	rightResult := recursion(root.Right)
	root.Right = nil
	result = append(result, rightResult...)

	return result
}

func inorderTraversal1(root *common.TreeNode) []int {
	var (
		result []int
		stack  []*common.TreeNode
	)

	if root == nil {
		return result
	}

	stack = append(stack, root)

	for len(stack) > 0 {

		node := stack[len(stack) - 1]

		if node.Left == nil  {
			result = append(result, node.Val)
			stack = stack[:len(stack) - 1]

			if node.Right != nil {
				stack = append(stack, node.Right)
				node.Right = nil
			}
		} else {

			if node.Right != nil {
				stack =  append(stack[:len(stack) - 1], node.Right)		
				stack = append(stack, node)
				node.Right = nil
			}

			
			stack = append(stack, node.Left)
			node.Left = nil
			
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

	fmt.Println(inorderTraversal1(tree))
}
