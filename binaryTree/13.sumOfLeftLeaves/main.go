package main

import (
	"fmt"
	"leetcode/common"
)


func sumOfLeftLeaves(root *common.TreeNode) int {
	var (
		sum int
	)
	return getLeftVal(root, sum)
}

func getLeftVal(root *common.TreeNode, sum int) int {

	if root == nil {
		return sum
	}

	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		sum += root.Left.Val
	}
	lSum := getLeftVal(root.Left, sum)
	rSum := getLeftVal(root.Right, sum)

	return lSum + rSum - sum
}



func sumOfLeftLeavesV2(root *common.TreeNode) int {
	var (
		sum int
		stack []*common.TreeNode
	)

	if root == nil {
		return 0
	}

	stack = append(stack, root)

	for len(stack) > 0 {
		node := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]

		if node.Left != nil &&  node.Left.Left == nil && node.Left.Right == nil {
			sum  += node.Left.Val
		}

		if node.Right != nil {
			stack = append(stack, node.Right)
		}

		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}


	return sum
}





func main() {
	var (
		tree = &common.TreeNode{
			Val: 1,
			Left: &common.TreeNode{
				Val: 2,
				Left: &common.TreeNode{
					Val: 4,
					Left: &common.TreeNode{
						Val: 6,
					},
				},
				Right: &common.TreeNode{
					Val: 5,
					Right: &common.TreeNode{
						Val: 9,
					},
				},
			},
			Right: &common.TreeNode{
				Val: 3,
				Left: &common.TreeNode{
					Val: 7,
				},
				Right: &common.TreeNode{
					Val: 8,
				},
			},
		}
	)

	fmt.Println(sumOfLeftLeavesV2(tree))
}
