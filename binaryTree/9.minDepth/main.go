package main

import (
	"leetcode/common"
)


func minDepth1(root *common.TreeNode) int {
	var (
		leftDepth, rightDepth int
	)
	if root == nil {
		return 0
	}

	// 当一个左子树为空，右不为空，这时并不是最低点
	// 当一个右子树为空，左不为空，这时并不是最低点
	leftDepth = minDepth1(root.Left)
	rightDepth = minDepth1(root.Right)
	if root.Left != nil && root.Right == nil {
		leftDepth++
	}else if root.Left == nil && root.Right != nil{
		rightDepth++
	}

	return min(leftDepth, rightDepth) + 1
}


func min(a, b int) int {
	var (
		minNum int
	)
	if a > b {
		minNum = b
	}else {
		minNum = a
	}

	return minNum
}

func minDepth(root *common.TreeNode) int {
	var (
		minDepth int
		queue []*common.TreeNode
	)

	if root == nil {
		return  0
	}

	queue = append(queue, root)

	for len(queue) > 0 {
		minDepth += 1
		curLen := len(queue)
		for i:=0; i < curLen; i++ {
			
			if queue[i].Left == nil && queue[i].Right == nil {
				return minDepth
			}

			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}

			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}

		}
		queue = queue[curLen:]
	}

	return minDepth
}

func main() {
	var (
		tree = &common.TreeNode{
			Val: 1,
			Left: &common.TreeNode{
				Val: 2,
				Left: &common.TreeNode{
					Val: 3,
				},
			},
		}
	)

	print(minDepth1(tree))
}