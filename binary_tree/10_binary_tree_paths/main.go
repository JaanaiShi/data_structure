package main

import (
	"fmt"
	"strconv"
	"strings"
)

// TreeNode 二叉树节点定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// binaryTreePaths 二叉树的所有路径
// 给你一个二叉树的根节点 root ，按 任意顺序 ，返回所有从根节点到叶子节点的路径。
// 叶子节点 是指没有子节点的节点。
//
// 示例 1：
// 输入：root = [1,2,3,null,5]
// 输出：["1->2->5","1->3"]
//
// 示例 2：
// 输入：root = [1]
// 输出：["1"]
//
// 提示：
// - 树中节点的数目在范围 [1, 100] 内
// - -100 <= Node.val <= 100
//
// LeetCode 257: https://leetcode.cn/problems/binary-tree-paths/
func binaryTreePaths(root *TreeNode) []string {
	res := make([]string, 0)
	path := strconv.Itoa(root.Val)
	dfs(root, &res, path)
	return res
}

func dfs(root *TreeNode, res *[]string, path string) {
	if root.Left == nil && root.Right == nil {
		*res = append(*res, path)
		return
	}
	if root.Left != nil {
		temp := path + "->" + strconv.Itoa(root.Left.Val)
		dfs(root.Left, res, temp)
	}
	if root.Right != nil {
		temp := path + "->" + strconv.Itoa(root.Right.Val)
		dfs(root.Right, res, temp)
	}
}

// sortedEqual 判断两个字符串切片排序后是否相等（忽略顺序）
func sortedEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	// 简单判断：将两个切片都拼接后比较（实际测试建议排序后比较）
	count := make(map[string]int)
	for _, s := range a {
		count[s]++
	}
	for _, s := range b {
		count[s]--
		if count[s] < 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("开始测试 binary_tree/10_binary_tree_paths...")

	// 测试用例 1: [1,2,3,null,5]
	//      1
	//     / \
	//    2   3
	//     \
	//      5
	// 路径：["1->2->5", "1->3"]
	root1 := &TreeNode{Val: 1}
	root1.Left = &TreeNode{Val: 2}
	root1.Right = &TreeNode{Val: 3}
	root1.Left.Right = &TreeNode{Val: 5}
	want1 := []string{"1->2->5", "1->3"}
	got1 := binaryTreePaths(root1)
	if sortedEqual(got1, want1) {
		fmt.Printf("PASS: 测试用例 1\n")
	} else {
		fmt.Printf("FAIL: 测试用例 1, want=%v, got=%v\n",
			strings.Join(want1, ", "), strings.Join(got1, ", "))
	}

	// 测试用例 2: 单节点
	root2 := &TreeNode{Val: 1}
	want2 := []string{"1"}
	got2 := binaryTreePaths(root2)
	if sortedEqual(got2, want2) {
		fmt.Printf("PASS: 测试用例 2\n")
	} else {
		fmt.Printf("FAIL: 测试用例 2, want=%v, got=%v\n",
			strings.Join(want2, ", "), strings.Join(got2, ", "))
	}

	fmt.Println("测试结束。")
}
