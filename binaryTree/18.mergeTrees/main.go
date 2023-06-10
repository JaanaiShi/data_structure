package main

import "leetcode/common"



func mergeTrees(root1 *common.TreeNode, root2 *common.TreeNode) *common.TreeNode {
	var (
		root common.TreeNode

	)
	if root1 == nil && root2 == nil {
		return nil
	}

	if root1 != nil && root2 != nil  {
		root.Val = root1.Val + root2.Val
		root.Left = mergeTrees(root1.Left, root2.Left)
		root.Right = mergeTrees(root1.Right, root2.Right)
	}else if root1 == nil && root2 != nil {
		root.Val = root2.Val
		root.Left = mergeTrees(nil, root2.Left)
		root.Right = mergeTrees(nil, root2.Right)
	}else if root1 != nil && root2 == nil {
		root.Val = root1.Val
		root.Left = mergeTrees(root1.Left, nil)
		root.Right = mergeTrees(root1.Right, nil)
	}
	return &root

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

	root2 := &common.TreeNode{
		Val: 1,
		Left: &common.TreeNode{
			Val: 2,
			Right: &common.TreeNode{
				Val: 4,
				Left: &common.TreeNode{
					Val: 3,
				},
			},
		},
	}


	common.PreOrder(mergeTrees(root1, root2))



}