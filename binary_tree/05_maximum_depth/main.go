package main

import "fmt"

// TreeNode 二叉树节点定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// maxDepth 二叉树的最大深度
// 给定一个二叉树 root ，返回其最大深度。
// 二叉树的 最大深度 是指从根节点到最远叶子节点的最长路径上的节点数。
//
// 示例 1：
// 输入：root = [3,9,20,null,null,15,7]
// 输出：3
//
// 示例 2：
// 输入：root = [1,null,2]
// 输出：2
func maxDepth(root *TreeNode) int {
	// TODO: 请实现解法
	res := 0
	if root == nil {
		return res
	}
	nodeList := []*TreeNode{root}
	for len(nodeList) != 0 {
		res++
		temp := make([]*TreeNode, 0)
		for _, v := range nodeList {
			if v.Left != nil {
				temp = append(temp, v.Left)
			}
			if v.Right != nil {
				temp = append(temp, v.Right)
			}
		}
		nodeList = temp
	}

	return res
}

func main() {
	// 测试用例 1: [3,9,20,null,null,15,7]
	//      3
	//     / \
	//    9  20
	//       / \
	//      15  7
	root1 := &TreeNode{Val: 3}
	root1.Left = &TreeNode{Val: 9}
	root1.Right = &TreeNode{Val: 20}
	root1.Right.Left = &TreeNode{Val: 15}
	root1.Right.Right = &TreeNode{Val: 7}
	want1 := 3

	fmt.Println("开始测试 binary_tree/05_maximum_depth...")
	got1 := maxDepth(root1)
	if got1 == want1 {
		fmt.Printf("PASS: 测试用例 1\n")
	} else {
		fmt.Printf("FAIL: 测试用例 1, want=%d, got=%d\n", want1, got1)
	}

	// 测试用例 2: 空树
	var root2 *TreeNode
	want2 := 0
	got2 := maxDepth(root2)
	if got2 == want2 {
		fmt.Printf("PASS: 测试用例 2\n")
	} else {
		fmt.Printf("FAIL: 测试用例 2, want=%d, got=%d\n", want2, got2)
	}
	fmt.Println("测试结束。")
}
