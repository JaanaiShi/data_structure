package main

import "fmt"

// TreeNode 二叉树节点定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isBalanced 平衡二叉树
// 给定一个二叉树，判断它是否是高度平衡的二叉树。
// 本题中，一棵高度平衡二叉树定义为：
// 一个二叉树每个节点的左右两个子树的高度差的绝对值不超过 1 。
//
// 示例 1：
// 输入：root = [3,9,20,null,null,15,7]
// 输出：true
//
// 示例 2：
// 输入：root = [1,2,2,3,3,null,null,4,4]
// 输出：false
//
// 示例 3：
// 输入：root = []
// 输出：true
//
// 提示：
// - 树中的节点数在范围 [0, 5000] 内
// - -10^4 <= Node.val <= 10^4
//
// LeetCode 110: https://leetcode.cn/problems/balanced-binary-tree/
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return false
	}
	isHigh := dfs(root)
	if isHigh == -1 {
		return false
	}

	return true
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l, r := dfs(root.Left), dfs(root.Right)
	if l == -1 || r == -1 {
		return -1
	}
	if l-r > 1 || r-l > 1 {
		return -1
	}
	return max(l, r) + 1
}

func main() {
	fmt.Println("开始测试 binary_tree/09_balanced_binary_tree...")

	// 测试用例 1: [3,9,20,null,null,15,7] 平衡树
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
	want1 := true
	got1 := isBalanced(root1)
	if got1 == want1 {
		fmt.Printf("PASS: 测试用例 1\n")
	} else {
		fmt.Printf("FAIL: 测试用例 1, want=%v, got=%v\n", want1, got1)
	}

	// 测试用例 2: [1,2,2,3,3,null,null,4,4] 不平衡树
	//         1
	//        / \
	//       2   2
	//      / \
	//     3   3
	//    / \
	//   4   4
	root2 := &TreeNode{Val: 1}
	root2.Left = &TreeNode{Val: 2}
	root2.Right = &TreeNode{Val: 2}
	root2.Left.Left = &TreeNode{Val: 3}
	root2.Left.Right = &TreeNode{Val: 3}
	root2.Left.Left.Left = &TreeNode{Val: 4}
	root2.Left.Left.Right = &TreeNode{Val: 4}
	want2 := false
	got2 := isBalanced(root2)
	if got2 == want2 {
		fmt.Printf("PASS: 测试用例 2\n")
	} else {
		fmt.Printf("FAIL: 测试用例 2, want=%v, got=%v\n", want2, got2)
	}

	// 测试用例 3: 空树
	var root3 *TreeNode
	want3 := true
	got3 := isBalanced(root3)
	if got3 == want3 {
		fmt.Printf("PASS: 测试用例 3\n")
	} else {
		fmt.Printf("FAIL: 测试用例 3, want=%v, got=%v\n", want3, got3)
	}

	fmt.Println("测试结束。")
}
