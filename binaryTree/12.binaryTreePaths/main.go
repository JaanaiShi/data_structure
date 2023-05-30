package main

import (
	"fmt"
	"leetcode/common"
	"strconv"
)

// binaryTreePaths 递归
func binaryTreePaths(root *common.TreeNode) []string {
	
	var (
		result  []string
		path string
		resultMap =  make(map[string]struct{}, 0)
	)

	if root == nil {
		return result
	}
	
	getPath(root, path, resultMap)

	for k, _ := range resultMap {
		result = append(result, k)
	}

	return result

}

func getPath(root *common.TreeNode, path string, result map[string]struct{}) {

	if path == "" {
		path = strconv.Itoa(root.Val)
	}else {
		path = path + "->" + strconv.Itoa(root.Val)
	}
	if root.Left == nil && root.Right == nil {
		result[path] = struct{}{}
	}

	if root.Left != nil {
		getPath(root.Left, path, result)
	}

	if root.Right != nil {
		getPath(root.Right, path, result)
	}	
}


func binaryTreePathsV2(root *common.TreeNode) []string {
	
	var (
		result  []string
		stackPath []string
		stack []*common.TreeNode
	)

	if root == nil {
		return result
	}
	stack = append(stack, root)
	stackPath = append(stackPath, strconv.Itoa(root.Val))

	for len(stack) > 0 {
		node := stack[len(stack) - 1]
		path := stackPath[len(stackPath) - 1]

		stack = stack[:len(stack) - 1]
		stackPath = stackPath[:len(stackPath) - 1]
		if node.Left == nil && node.Right == nil {
			result = append(result, path)
			
		}
		if node.Right != nil {
			rPath := path + "->" + strconv.Itoa(node.Right.Val)
			stackPath = append(stackPath, rPath)
			stack = append(stack, node.Right)
		}

		if node.Left != nil {
			lpath := path + "->" + strconv.Itoa(node.Left.Val)
			stackPath = append(stackPath, lpath)
			stack = append(stack, node.Left)
		}
	}

	return result
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

	fmt.Println(binaryTreePathsV2(tree))
}