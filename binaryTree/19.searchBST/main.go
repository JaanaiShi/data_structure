package main

import "leetcode/common"

func searchBST(root *common.TreeNode, val int) *common.TreeNode {
	var (
		res *common.TreeNode
	)

	getSonTree(root, &res, val)


	return res
}

func getSonTree(root *common.TreeNode, result **common.TreeNode, val int) {
	if root == nil {
		return
	}

	if *result != nil {
		return
	}

	if root.Val == val {
		*result = root
		return
	}
	if root.Val > val {
		getSonTree(root.Left, result, val)

	}else if root.Val < val {
		getSonTree(root.Right, result, val)
	}
}

func main() {
	root1 := &common.TreeNode{
		Val: 1,
		Left: &common.TreeNode{
			Val: 2,
			Left: &common.TreeNode{
				Val: 3,
			},
			Right: &common.TreeNode{
				Val: 4,
			},
		},
	}

	common.PreOrder(searchBST(root1, 2))
}