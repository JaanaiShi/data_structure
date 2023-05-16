package main

import "leetcode/common"



func levelOrder(root *common.TreeNode) [][]int {
	var (
		result [][]int
		queue []*common.TreeNode
		temp  []int
	)

	if root == nil {
		return result
	}

	queue = append(queue, root)

	for len(queue) > 0 {
		length := len(queue)
		for i:=0; i < length; i++ {
			node := queue[0]
			temp = append(temp, node.Val)
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
	
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}		
		result = append(result, temp)
		temp = []int{}
	}
	return result
}


func main() {

}