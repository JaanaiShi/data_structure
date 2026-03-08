package main

import "fmt"

// TreeNode 二叉树节点定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// countNodes 完全二叉树的节点个数
// 给你一棵 完全二叉树 的根节点 root ，求出该树的节点个数。
// 完全二叉树 的定义如下：在完全二叉树中，除了最底层节点可能没填满外，
// 其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置。
// 若最底层为第 h 层（从第 1 层开始计起），则该层包含 1~ 2^(h-1) 个节点。
//
// 示例 1：
// 输入：root = [1,2,3,4,5,6]
// 输出：6
//
// 示例 2：
// 输入：root = []
// 输出：0
//
// 示例 3：
// 输入：root = [1]
// 输出：1
//
// 提示：
// - 树中节点的数目范围是 [0, 5 * 10^4]
// - 0 <= Node.val <= 5 * 10^4
// - 题目数据保证输入的树是 完全二叉树
//
// 进阶：遍历树来统计节点是一种时间复杂度为 O(n) 的简单解决方案。
//
//	你可以设计一个更快的算法吗？
//
// LeetCode 222: https://leetcode.cn/problems/count-complete-tree-nodes/
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 1
	res += countNodes(root.Left)
	res += countNodes(root.Right)
	return res
}

func countNodes1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftH, rightH := 0, 0
	leftNode, rightNode := root.Left, root.Right
	for leftNode != nil {
		leftH++
		leftNode = leftNode.Left
	}

	for rightNode != nil {
		rightH++
		rightNode = rightNode.Right
	}

	if leftH == rightH {
		return 2<<leftH - 1
	}
	return countNodes(root.Left) + countNodes(root.Right) + 1
}

func main() {
	fmt.Println("开始测试 binary_tree/08_count_complete_tree_nodes...")

	// 测试用例 1: [1,2,3,4,5,6]
	//       1
	//      / \
	//     2   3
	//    / \ /
	//   4  5 6
	root1 := &TreeNode{Val: 1}
	root1.Left = &TreeNode{Val: 2}
	root1.Right = &TreeNode{Val: 3}
	root1.Left.Left = &TreeNode{Val: 4}
	root1.Left.Right = &TreeNode{Val: 5}
	root1.Right.Left = &TreeNode{Val: 6}
	want1 := 6
	got1 := countNodes(root1)
	if got1 == want1 {
		fmt.Printf("PASS: 测试用例 1\n")
	} else {
		fmt.Printf("FAIL: 测试用例 1, want=%d, got=%d\n", want1, got1)
	}

	// 测试用例 2: 空树
	var root2 *TreeNode
	want2 := 0
	got2 := countNodes(root2)
	if got2 == want2 {
		fmt.Printf("PASS: 测试用例 2\n")
	} else {
		fmt.Printf("FAIL: 测试用例 2, want=%d, got=%d\n", want2, got2)
	}

	// 测试用例 3: 单节点
	root3 := &TreeNode{Val: 1}
	want3 := 1
	got3 := countNodes(root3)
	if got3 == want3 {
		fmt.Printf("PASS: 测试用例 3\n")
	} else {
		fmt.Printf("FAIL: 测试用例 3, want=%d, got=%d\n", want3, got3)
	}

	fmt.Println("测试结束。")
}
