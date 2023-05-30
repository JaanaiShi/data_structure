package main

import (
	"fmt"
	"leetcode/common"
)

/*
	区分两个概念：深度和高度
	深度：二叉树的深度是从跟几点开始（其深度为1）自顶向下逐层累加的。求深度可以用前序遍历，因为是从上往下查
	高度：二叉树的高度是从叶子节点开始（其高度为1）自底向上逐层累加的。求高度可以用后序遍历，因为是从下往上查
*/

func isBalanced(root *common.TreeNode) bool {
	height := getHigh(root)
	if height == -1 {
		return false
	}

	return true
}


// 返回该节点为根节点的二叉树的高度，如果不是平衡二叉树则返回-1
func getHigh(root *common.TreeNode) int {
	if root == nil {
		return 0
	}

	l, r := getHigh(root.Left), getHigh(root.Right)

	if l == -1 || r == -1 {
		return -1
	}

	if l - r > 1 || r - l > 1{
		return -1
	}

	return max(l, r) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


func getHightV2(root *common.TreeNode) int {
	var (
		stack []*common.TreeNode
		l, r int
		result []int
	)

	if root == nil {
		return 0
	}

	l += 1
	r += 1

	stack = append(stack, root)

	for len(stack) > 0 {
		node := stack[len(stack) - 1]
		if node.Left == nil && node.Right == nil {
			stack = stack[:len(stack) - 1]
			result = append(result, node.Val)
		}

		if node.Right != nil {
			r++
			stack = append(stack, node.Right)
			node.Right = nil
		}

		if node.Left != nil {
			l++
			stack = append(stack, node.Left)
			node.Left = nil
		}

		if l - r > 1 && r - l > 1 {
			return -1
		}
	}

	fmt.Println(result)
	fmt.Printf("l: %d r: %d\n", l, r)
	return max(l, r)
}

// TODO 迭代法未实现
func getDepth(root *common.TreeNode) {




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

	fmt.Println(getHightV2(tree))
}
