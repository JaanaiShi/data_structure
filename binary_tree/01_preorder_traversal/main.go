package main

import "fmt"

// TreeNode 二叉树节点定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// preorderTraversal 二叉树的前序遍历
// 给你二叉树的根节点 root ，返回它节点值的 前序 遍历。
//
// 示例 1：
// 输入：root = [1,null,2,3]
// 输出：[1,2,3]
//
// 示例 2：
// 输入：root = []
// 输出：[]
//
// 示例 3：
// 输入：root = [1]
// 输出：[1]
func preorderTraversal1(root *TreeNode) []int {
	// TODO: 请实现解法
	res := make([]int, 0)
	getNodeVal(root, &res)
	return res
}

func getNodeVal(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	if root.Left != nil {
		getNodeVal(root.Left, res)
	}
	if root.Right != nil {
		getNodeVal(root.Right, res)
	}
}

// 迭代法
func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	list := []*TreeNode{root}

	for len(list) != 0 {
		node := list[len(list)-1]
		list = list[:len(list)-1]
		res = append(res, node.Val)
		if node.Right != nil {
			list = append(list, node.Right)
		}
		if node.Left != nil {
			list = append(list, node.Left)
		}
	}
	return res
}

// sliceEqual 判断两个整数切片是否相等。
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
	// 测试用例 1: [1,null,2,3]
	//    1
	//     \
	//      2
	//     /
	//    3
	root1 := &TreeNode{Val: 1}
	root1.Right = &TreeNode{Val: 2}
	root1.Right.Left = &TreeNode{Val: 3}
	want1 := []int{1, 2, 3}

	fmt.Println("开始测试 binary_tree/01_preorder_traversal...")
	got1 := preorderTraversal(root1)
	if sliceEqual(got1, want1) {
		fmt.Printf("PASS: 测试用例 1\n")
	} else {
		fmt.Printf("FAIL: 测试用例 1, want=%v, got=%v\n", want1, got1)
	}

	// 测试用例 2: 空树
	var root2 *TreeNode
	want2 := []int{}
	got2 := preorderTraversal(root2)
	if sliceEqual(got2, want2) {
		fmt.Printf("PASS: 测试用例 2\n")
	} else {
		fmt.Printf("FAIL: 测试用例 2, want=%v, got=%v\n", want2, got2)
	}
	fmt.Println("测试结束。")
}
