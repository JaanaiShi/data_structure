package main

import (
	"fmt"
	"leetcode/common"
)


func isValidBST(root *common.TreeNode) bool {
	var (
		res = new(bool)
	)

	*res = true
	
	fmt.Println(inOrder(root, res))
	return *res
}

func inOrder(root *common.TreeNode, status *bool) ([]int) {
	var (
		res, left, right []int
	)

	if status != nil && !*status {
		return res
	}

	if root == nil {
		return []int{}
	}

	if root.Left != nil {
		left = inOrder(root.Left, status)
	}
	res = append(res, left...)

	if len(res) != 0 {
		if res[len(res) - 1] > root.Val {
			*status = false
			return res
		}
	}
	res = append(res, root.Val)


	if root.Right != nil {
		right = inOrder(root.Right, status)
	}
	res = append(res, right...)
	return res
}


func main() {
	// root1 := &common.TreeNode{
	// 	Val: 8,
	// 	Left: &common.TreeNode{
	// 		Val: 3,
	// 		Left: &common.TreeNode{
	// 			Val: 2,
	// 		},
	// 		Right: &common.TreeNode{
	// 			Val: 4,
	// 			Right: &common.TreeNode{
	// 				Val: 10,
	// 			},
	// 		},
	// 	},
	// }

	root2 := &common.TreeNode{
		Val: 5,
		Left: &common.TreeNode{
			Val: 1,
		},
		Right: &common.TreeNode{
			Val: 4,
			Left: &common.TreeNode{
				Val: 3,
			},
			Right: &common.TreeNode{
				Val: 6,
			},
		},
	}

	fmt.Println(isValidBST(root2))
}