package main

import (
	"leetcode/common"
	"fmt"
)
func maxDepth(root *common.TreeNode) int {
	if root == nil {
		return 0
	}

	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
} 


func maxDepth1(root *common.TreeNode) int {
	var (
		depth int
		queue []*common.TreeNode
	)

	queue = append(queue, root)
	for len(queue) > 0 {
		depth += 1
		curLen := len(queue)
		for i:=0; i < curLen; i++ {
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[curLen:]
	}
	return depth
}

func max(a, b int) int {
	var (
		maxNum int
	)
	if a > b {
		maxNum =  a
	}else {
		maxNum =  b
	}

	return maxNum
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
			Right: &common.TreeNode{
				Val: 2,
			},
		}
	)

	fmt.Println(maxDepth1(tree))
}