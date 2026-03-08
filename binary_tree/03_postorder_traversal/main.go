package main

import "fmt"

// TreeNode 二叉树节点定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// postorderTraversal 二叉树的后序遍历
// 给你一棵二叉树的根节点 root ，返回其中节点值的 后序遍历 。
//
// 示例 1：
// 输入：root = [1,null,2,3]
// 输出：[3,2,1]
//
// 示例 2：
// 输入：root = []
// 输出：[]
//
// 示例 3：
// 输入：root = [1]
// 输出：[1]
func postorderTraversal1(root *TreeNode) []int {
	// TODO: 请实现解法
	res := make([]int, 0)
	getNodeVal(root, &res)
	return res
}

func getNodeVal(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	getNodeVal(root.Left, res)
	getNodeVal(root.Right, res)
	*res = append(*res, root.Val)
}

func postorderTraversal(root *TreeNode) []int {
	// TODO: 请实现解法
	res := make([]int, 0)

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
	want1 := []int{3, 2, 1}

	fmt.Println("开始测试 binary_tree/03_postorder_traversal...")
	got1 := postorderTraversal(root1)
	if sliceEqual(got1, want1) {
		fmt.Printf("PASS: 测试用例 1\n")
	} else {
		fmt.Printf("FAIL: 测试用例 1, want=%v, got=%v\n", want1, got1)
	}

	// 测试用例 2: 空树
	var root2 *TreeNode
	want2 := []int{}
	got2 := postorderTraversal(root2)
	if sliceEqual(got2, want2) {
		fmt.Printf("PASS: 测试用例 2\n")
	} else {
		fmt.Printf("FAIL: 测试用例 2, want=%v, got=%v\n", want2, got2)
	}
	fmt.Println("测试结束。")
}
