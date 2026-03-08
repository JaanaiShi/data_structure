package main

import "fmt"

// TreeNode 二叉树节点定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// levelOrder 二叉树的层序遍历
// 给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。
//
// 示例 1：
// 输入：root = [3,9,20,null,null,15,7]
// 输出：[[3],[9,20],[15,7]]
//
// 示例 2：
// 输入：root = [1]
// 输出：[[1]]
//
// 示例 3：
// 输入：root = []
// 输出：[]
func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	nodeList := []*TreeNode{root}

	for len(nodeList) != 0 {
		temp := make([]*TreeNode, 0)
		vals := make([]int, 0)
		for _, v := range nodeList {
			vals = append(vals, v.Val)
			if v.Left != nil {
				temp = append(temp, v.Left)
			}
			if v.Right != nil {
				temp = append(temp, v.Right)
			}
		}
		if len(vals) != 0 {
			res = append(res, vals)
		}
		nodeList = temp
	}

	return res
}

// slice2DEqual 判断两个二维整数切片是否相等。
func slice2DEqual(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
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
	want1 := [][]int{{3}, {9, 20}, {15, 7}}

	fmt.Println("开始测试 binary_tree/04_level_order_traversal...")
	got1 := levelOrder(root1)
	if slice2DEqual(got1, want1) {
		fmt.Printf("PASS: 测试用例 1\n")
	} else {
		fmt.Printf("FAIL: 测试用例 1, want=%v, got=%v\n", want1, got1)
	}

	// 测试用例 2: 空树
	var root2 *TreeNode
	want2 := [][]int{}
	got2 := levelOrder(root2)
	if slice2DEqual(got2, want2) {
		fmt.Printf("PASS: 测试用例 2\n")
	} else {
		fmt.Printf("FAIL: 测试用例 2, want=%v, got=%v\n", want2, got2)
	}
	fmt.Println("测试结束。")
}
