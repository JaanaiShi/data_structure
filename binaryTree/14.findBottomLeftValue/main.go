package main

import (
	"fmt"
	"leetcode/common"
)

var (
	Default_Depth = 0
	result = 0
)



func findBottomLeftValue(root *common.TreeNode) int {
 
	getLeftBottomValue(root, 0)
	return 0
}

func getLeftBottomValue(root *common.TreeNode, depth int) {

	if root.Left == nil && root.Right == nil {
		if depth > Default_Depth {
			Default_Depth = depth
			result = root.Val
		}
	}

	if root.Left != nil {
		depth++
		getLeftBottomValue(root.Left, depth)
		depth--  // 回溯
	}

	if root.Right != nil {
		depth++
		getLeftBottomValue(root.Right, depth)
		depth--  // 回溯
	}


}


func findBottomLeftValueV2(root *common.TreeNode) int {

	var (
		stack []*common.TreeNode
		result int
	)

	if root == nil {
		return 0
	}

	stack = append(stack, root)

	for len(stack) > 0 {
		length := len(stack)
		flag := true
		for i:=0; i < length; i++ {
			node := stack[i]
			if node.Left == nil && node.Right == nil {
				// 判断是否是第一次
				if flag {
					result = node.Val
					flag = false
				}
			}

			if node.Left != nil {
				stack = append(stack, node.Left)
			}

			if node.Right != nil {
				stack = append(stack, node.Right)
			}


		}

		stack = stack[length:]
		
	}
	
	return result
}

func prePrint(root *common.TreeNode) {
	if root == nil {
		return
	}

	fmt.Println(root.Val)



	prePrint(root.Left)
	prePrint(root.Right)
}


func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getDepth(root *common.TreeNode) int {
	if root == nil {
		return 0
	}

	return max(getDepth(root.Left), getDepth(root.Right)) + 1
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
	fmt.Println(findBottomLeftValueV2(tree))
	fmt.Println(result)
}