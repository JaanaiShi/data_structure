package main

import "fmt"

// TreeNode 二叉树节点定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isSymmetric 对称二叉树
// 给你一个二叉树的根节点 root ，检查它是否轴对称。
//
// 示例 1：
// 输入：root = [1,2,2,3,4,4,3]
// 输出：true
//
// 示例 2：
// 输入：root = [1,2,2,null,3,null,3]
// 输出：false
//
// 提示：
// - 树中节点数目在范围 [1, 1000] 内
// - -100 <= Node.val <= 100
//
// LeetCode 101: https://leetcode.cn/problems/symmetric-tree/
func isSymmetric1(root *TreeNode) bool {

	fn := func(data []*TreeNode) bool {
		if len(data)%2 == 1 {
			return false
		}
		for i := 0; i < len(data)/2; i++ {
			if (data[i] == nil && data[len(data)-i-1] != nil) ||
				(data[i] != nil && data[len(data)-i-1] == nil) {
				return false
			}
			if data[i] != nil && data[len(data)-i-1] != nil &&
				data[i].Val != data[len(data)-i-1].Val {
				return false
			}
		}
		return true
	}

	nodeList := []*TreeNode{root}
	for len(nodeList) != 0 {
		temp := make([]*TreeNode, 0)
		for _, v := range nodeList {
			if v == nil {
				continue
			}
			temp = append(temp, v.Left)
			temp = append(temp, v.Right)
		}
		if !fn(temp) {
			return false
		}
		nodeList = temp
	}
	return true
}

func isSymmetric(root *TreeNode) bool {
	return dfs(root.Left, root.Right)
}

func dfs(left *TreeNode, right *TreeNode) bool {
	if left == nil && right != nil {
		return false
	} else if left != nil && right == nil {
		return false
	} else if left == nil && right == nil {
		return true
	} else if left != nil && right != nil && left.Val != right.Val {
		return false
	}
	a1 := dfs(left.Left, right.Right)
	a2 := dfs(left.Right, right.Left)
	return a1 && a2
}

// 判断是否对成

func main() {
	fmt.Println("开始测试 binary_tree/06_symmetric_tree...")

	// 测试用例 1: [1,2,2,3,4,4,3] 对称树
	//         1
	//        / \
	//       2   2
	//      / \ / \
	//     3  4 4  3
	root1 := &TreeNode{Val: 1}
	root1.Left = &TreeNode{Val: 2}
	root1.Right = &TreeNode{Val: 2}
	root1.Left.Left = &TreeNode{Val: 3}
	root1.Left.Right = &TreeNode{Val: 4}
	root1.Right.Left = &TreeNode{Val: 4}
	root1.Right.Right = &TreeNode{Val: 3}
	want1 := true
	got1 := isSymmetric(root1)
	if got1 == want1 {
		fmt.Printf("PASS: 测试用例 1\n")
	} else {
		fmt.Printf("FAIL: 测试用例 1, want=%v, got=%v\n", want1, got1)
	}

	// 测试用例 2: [1,2,2,null,3,null,3] 非对称树
	//         1
	//        / \
	//       2   2
	//        \   \
	//         3   3
	root2 := &TreeNode{Val: 1}
	root2.Left = &TreeNode{Val: 2}
	root2.Right = &TreeNode{Val: 2}
	root2.Left.Right = &TreeNode{Val: 3}
	root2.Right.Right = &TreeNode{Val: 3}
	want2 := false
	got2 := isSymmetric(root2)
	if got2 == want2 {
		fmt.Printf("PASS: 测试用例 2\n")
	} else {
		fmt.Printf("FAIL: 测试用例 2, want=%v, got=%v\n", want2, got2)
	}

	// 测试用例 3: 单节点
	root3 := &TreeNode{Val: 1}
	want3 := true
	got3 := isSymmetric(root3)
	if got3 == want3 {
		fmt.Printf("PASS: 测试用例 3\n")
	} else {
		fmt.Printf("FAIL: 测试用例 3, want=%v, got=%v\n", want3, got3)
	}

	fmt.Println("测试结束。")
}
