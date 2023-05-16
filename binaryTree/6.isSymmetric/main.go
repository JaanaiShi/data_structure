package main

import (
	"fmt"
	"leetcode/common"
)

func isSymmetric(root *common.TreeNode) bool {
	var (
		result bool
		queue []*common.TreeNode
	)

	if root == nil {
		result = true
		return result
	}

	queue = append(queue, root.Left)
	queue = append(queue, root.Right)

	for len(queue) > 0 {
		length := len(queue)
		i := 0
		middleware := length / 2
		for i < middleware {
			if queue[i] != nil && queue[length - i - 1] == nil {
				return false
			} else if queue[i] == nil && queue[length - i - 1] != nil {
				return false
			} else if queue[i] == nil && queue[length - i - 1] == nil {
				i++
				continue
			} else if queue[i].Val == queue[length - i - 1].Val {
				queue = append(queue, queue[i].Left)
				queue = append(queue, queue[i].Right)
				queue = append(queue, queue[length - i - 1].Left)
				queue = append(queue, queue[length - i - 1].Right)
				i++
				continue
			}else {
				return result
			} 
		}
		queue = queue[length:]
	}

	result = true

	return result
}


func main() {
	var (
		tree = &common.TreeNode{
			Val: 1,
			Left: &common.TreeNode{
				Val: 2,
			},
			Right: &common.TreeNode{
				Val: 2,
			},
		}
	)

	fmt.Println(isSymmetric(tree))
}