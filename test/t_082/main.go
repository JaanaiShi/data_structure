package main

import "container/list"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode1(root *TreeNode) []int {
	result := []int{}
	stack := list.New()
	cur := root
	maxCount := 0
	count := 0
	var pre *TreeNode
	for cur != nil || stack.Len() != 0 {
		if cur != nil {
			stack.PushBack(cur)
			cur = cur.Left
		} else {
			cur := stack.Remove(stack.Back()).(*TreeNode)
			if pre == nil || pre.Val == cur.Val {
				count++
			} else {
				count = 1
			}
			pre = cur
			if count > maxCount {
				maxCount = count
				result = result[0:0]
				result = append(result, cur.Val)
			} else if count == maxCount {
				result = append(result, cur.Val)
			}
			// 右
			cur = cur.Right
		}
	}
	return result
}

func main() {
	
}