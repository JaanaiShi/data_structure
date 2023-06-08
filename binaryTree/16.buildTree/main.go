package main

import (
	"leetcode/common"
)


func buildTree(inorder []int, postorder []int) *common.TreeNode {
	return getTree(inorder, postorder)
}


func getTree( left, right []int) *common.TreeNode {
	var (
		node common.TreeNode
	)
	if len(right) == 0 {
		return nil
	}

	rootNode := right[len(right) - 1]
	node.Val = rootNode
	if len(right) == 1 {
		return &node
	}

	// 切割点
	toCutLeft := 0
	for i:=0; i < len(left); i++ {
		if left[i] == rootNode {
			toCutLeft = i
			break
		}
	}


	// 寻找右
	toCutRight := len(left[:toCutLeft])

	left1 := left[:toCutLeft]
	left2 := left[toCutLeft+1:]
	right1 := right[:toCutRight]
	right2 := right[toCutRight:len(right) - 1]


	node.Left = getTree(left1, right1)
	node.Right = getTree(left2,right2)


	return &node
}


func main() {
	left := []int{7,10,6,3,9,8}
	right := []int{7,6,10,9,8,3}

	tree := buildTree(left, right)
	common.PreOrder(tree)
}