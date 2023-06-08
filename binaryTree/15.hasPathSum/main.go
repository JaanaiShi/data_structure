package main

import (
	"fmt"
	"leetcode/common"
)


var (
	result []int
)


func hasPathSum(root *common.TreeNode, targetSum int) bool {
	var (
		result = new(bool)
		e = false
	)

	*result = e
	if root == nil {
        return false
    }

	getPrePathSum(root, 0, targetSum, result)


	return *result
}

func getPrePathSum(root *common.TreeNode, sum, target int, result *bool) {
	sum += root.Val
	if root.Left == nil && root.Right == nil {
		if target == sum {
			*result = true
		}
		return
	}
	if root.Left != nil {
		getPrePathSum(root.Left, sum, target, result)

	}
	if root.Right != nil {
		getPrePathSum(root.Right, sum, target, result)
	}
}

func main() {
	var (
		tree = &common.TreeNode{
			Val: 1,
			Left: &common.TreeNode{
				Val: 2,
			},
		}
	)
	fmt.Println(hasPathSum(tree, 3))

}