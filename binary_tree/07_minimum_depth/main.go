package main

import (
	"fmt"
	"math"
)

// TreeNode 二叉树节点定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// minDepth 二叉树的最小深度
// 给定一个二叉树，找出其最小深度。
// 最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
// 说明：叶子节点是指没有子节点的节点。
//
// 示例 1：
// 输入：root = [3,9,20,null,null,15,7]
// 输出：2
//
// 示例 2：
// 输入：root = [2,null,3,null,4,null,5,null,6]
// 输出：5
//
// 提示：
// - 树中节点数的范围在 [0, 10^5] 内
// - -1000 <= Node.val <= 1000
//
// LeetCode 111: https://leetcode.cn/problems/minimum-depth-of-binary-tree/
func minDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	minList := make([]int, 0)
	depth := 1
	dfs(root, &minList, depth)
	res := math.MaxInt
	for _, v := range minList {
		if res > v {
			res = v
		}
	}
	return res
}

func dfs(root *TreeNode, res *[]int, depth int) {
	if root.Left == nil && root.Right == nil {
		*res = append(*res, depth)
		return
	}
	if root.Left != nil {
		dfs(root.Left, res, depth+1)
	}
	if root.Right != nil {
		dfs(root.Right, res, depth+1)
	}
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	nodeList := []*TreeNode{root}
	res := 1

	for len(nodeList) != 0 {
		temp := make([]*TreeNode, 0)
		for _, v := range nodeList {
			if v.Left == nil && v.Right == nil {
				return res
			}
			if v.Left != nil {
				temp = append(temp, v.Left)
			}
			if v.Right != nil {
				temp = append(temp, v.Right)
			}
		}
		nodeList = temp
		res++
	}
	return res
}

func main() {
	fmt.Println("开始测试 binary_tree/07_minimum_depth...")

	// 测试用例 1: [3,9,20,null,null,15,7]
	//      3
	//     / \
	//    9  20
	//       / \
	//      15  7
	// 最小深度为 2（3 -> 9）
	root1 := &TreeNode{Val: 3}
	root1.Left = &TreeNode{Val: 9}
	root1.Right = &TreeNode{Val: 20}
	root1.Right.Left = &TreeNode{Val: 15}
	root1.Right.Right = &TreeNode{Val: 7}
	want1 := 2
	got1 := minDepth(root1)
	if got1 == want1 {
		fmt.Printf("PASS: 测试用例 1\n")
	} else {
		fmt.Printf("FAIL: 测试用例 1, want=%d, got=%d\n", want1, got1)
	}

	// 测试用例 2: [2,null,3,null,4,null,5,null,6] 右斜树
	//    2
	//     \
	//      3
	//       \
	//        4
	//         \
	//          5
	//           \
	//            6
	// 最小深度为 5
	root2 := &TreeNode{Val: 2}
	root2.Right = &TreeNode{Val: 3}
	root2.Right.Right = &TreeNode{Val: 4}
	root2.Right.Right.Right = &TreeNode{Val: 5}
	root2.Right.Right.Right.Right = &TreeNode{Val: 6}
	want2 := 5
	got2 := minDepth(root2)
	if got2 == want2 {
		fmt.Printf("PASS: 测试用例 2\n")
	} else {
		fmt.Printf("FAIL: 测试用例 2, want=%d, got=%d\n", want2, got2)
	}

	// 测试用例 3: 空树
	var root3 *TreeNode
	want3 := 0
	got3 := minDepth(root3)
	if got3 == want3 {
		fmt.Printf("PASS: 测试用例 3\n")
	} else {
		fmt.Printf("FAIL: 测试用例 3, want=%d, got=%d\n", want3, got3)
	}

	fmt.Println("测试结束。")
}
