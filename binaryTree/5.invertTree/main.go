package main

import "leetcode/common"


func invertTree(root *common.TreeNode) *common.TreeNode {
	var (
		stack []*common.TreeNode
	)

	if root == nil {
		return root
	}

	stack = append(stack, root)

	for len(stack) > 0 {
		node := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]

		if node == nil {
			continue
		}

		if node.Left == nil && node.Right == nil {
			continue
		}

		node.Left, node.Right = node.Right, node.Left
		stack = append(stack, node.Left)
		stack = append(stack, node.Right)
	}

	return root

}

func invertTree1(root *common.TreeNode) *common.TreeNode {
	recursion(root)

	return root
}


func recursion(root *common.TreeNode)  {


	if root == nil {
		return 
	}

	if root.Left == nil && root.Right == nil {
		return
	}

	root.Left, root.Right = root.Right, root.Left

	recursion(root.Left)
	recursion(root.Right)
}
