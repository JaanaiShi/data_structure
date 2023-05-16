package main

import "leetcode/common"

func countNodes(root *common.TreeNode) int {
	var (
		count int
		queue  []*common.TreeNode
	)

	if root == nil {
		return count
	}

	queue = append(queue, root)
	count++
	for  len(queue) > 0 {
		curLen := len(queue)

		for i:=0; i < curLen; i++ {
			node := queue[i]
			if node.Left != nil {
				count++
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				count++
				queue = append(queue, node.Right)
			}
		}
		queue = queue[curLen:]
	}

	return count

}

func countNodes2(root *common.TreeNode) int {
	var (
		res int
	)
	if root == nil {
		return res
	}

	res = 1

	if root.Left != nil {
		res += countNodes2(root.Left)
	}

	if root.Right != nil {
		res += countNodes2(root.Right)
	}

	return res
}


// 递归遍历，根据完全二叉树的特性去遍历的
func countNodes1(root *common.TreeNode) int {
	
	if root == nil{
		return 0
	}

	leftH, rightH := 0, 0
	leftNode, rightNode := root.Left, root.Right

	for leftNode != nil {
		leftNode = leftNode.Left
		leftH++
	}

	for rightNode != nil {
		rightNode = rightNode.Right
		rightH++
	}

	if leftH == rightH {
		return (2 << leftH) - 1
	}

	return countNodes1(root.Left) + countNodes1(root.Right) + 1
}

