package main

import "fmt"

// TreeNode 二叉树节点定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// invertTree 翻转二叉树
// 给你一棵二叉树的根节点 root ，翻转这棵二叉树，并返回其根节点。
//
// 示例 1：
// 输入：root = [4,2,7,1,3,6,9]
// 输出：[4,7,2,9,6,3,1]
//
// 示例 2：
// 输入：root = [2,1,3]
// 输出：[2,3,1]
//
// 示例 3：
// 输入：root = []
// 输出：[]
//
// 提示：
// - 树中节点数目范围在 [0, 100] 内
// - -100 <= Node.val <= 100
//
// LeetCode 226: https://leetcode.cn/problems/invert-binary-tree/
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left != nil || root.Right != nil {
		root.Left, root.Right = root.Right, root.Left
		invertTree(root.Left)
		invertTree(root.Right)
	}
	return root
}

// levelOrder 层序遍历，用于将树转为切片方便比较结果
func levelOrder(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		res = append(res, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return res
}

// sliceEqual 判断两个整数切片是否相等
func sliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("开始测试 binary_tree/11_invert_binary_tree...")

	// 测试用例 1: [4,2,7,1,3,6,9]
	// 翻转前:         翻转后:
	//      4               4
	//     / \             / \
	//    2   7           7   2
	//   / \ / \         / \ / \
	//  1  3 6  9       9  6 3  1
	root1 := &TreeNode{Val: 4}
	root1.Left = &TreeNode{Val: 2}
	root1.Right = &TreeNode{Val: 7}
	root1.Left.Left = &TreeNode{Val: 1}
	root1.Left.Right = &TreeNode{Val: 3}
	root1.Right.Left = &TreeNode{Val: 6}
	root1.Right.Right = &TreeNode{Val: 9}
	want1 := []int{4, 7, 2, 9, 6, 3, 1}
	got1 := levelOrder(invertTree(root1))
	if sliceEqual(got1, want1) {
		fmt.Printf("PASS: 测试用例 1\n")
	} else {
		fmt.Printf("FAIL: 测试用例 1, want=%v, got=%v\n", want1, got1)
	}

	// 测试用例 2: [2,1,3]
	// 翻转前:    翻转后:
	//    2          2
	//   / \        / \
	//  1   3      3   1
	root2 := &TreeNode{Val: 2}
	root2.Left = &TreeNode{Val: 1}
	root2.Right = &TreeNode{Val: 3}
	want2 := []int{2, 3, 1}
	got2 := levelOrder(invertTree(root2))
	if sliceEqual(got2, want2) {
		fmt.Printf("PASS: 测试用例 2\n")
	} else {
		fmt.Printf("FAIL: 测试用例 2, want=%v, got=%v\n", want2, got2)
	}

	// 测试用例 3: 空树
	var root3 *TreeNode
	want3 := []int{}
	got3 := levelOrder(invertTree(root3))
	if sliceEqual(got3, want3) {
		fmt.Printf("PASS: 测试用例 3\n")
	} else {
		fmt.Printf("FAIL: 测试用例 3, want=%v, got=%v\n", want3, got3)
	}

	fmt.Println("测试结束。")
}
