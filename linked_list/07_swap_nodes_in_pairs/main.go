package main

import "fmt"

// ListNode 链表节点定义
type ListNode struct {
	Val  int
	Next *ListNode
}

// swapPairs 两两交换链表中的节点
// 给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
//
// 示例 1：
// 输入：head = [1,2,3,4]
// 输出：[2,1,4,3]
//
// 示例 2：
// 输入：head = []
// 输出：[]
//
// 示例 3：
// 输入：head = [1]
// 输出：[1]
//
// LeetCode 24: https://leetcode.cn/problems/swap-nodes-in-pairs/
func swapPairs(head *ListNode) *ListNode {
	// TODO: 请实现解法
	return nil
}

// arrayToListNode 将整数切片转换为链表，返回头节点。
func arrayToListNode(nums []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for _, v := range nums {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

// listNodeToArray 将链表转换为整数切片，方便断言比较。
func listNodeToArray(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
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
	fmt.Println("开始测试 linked_list/07_swap_nodes_in_pairs...")

	head1 := arrayToListNode([]int{1, 2, 3, 4})
	want1 := []int{2, 1, 4, 3}
	got1 := listNodeToArray(swapPairs(head1))
	if sliceEqual(got1, want1) {
		fmt.Printf("PASS: 测试用例 1\n")
	} else {
		fmt.Printf("FAIL: 测试用例 1, want=%v, got=%v\n", want1, got1)
	}

	var head2 *ListNode
	var want2 []int
	got2 := listNodeToArray(swapPairs(head2))
	if sliceEqual(got2, want2) {
		fmt.Printf("PASS: 测试用例 2\n")
	} else {
		fmt.Printf("FAIL: 测试用例 2, want=%v, got=%v\n", want2, got2)
	}

	head3 := arrayToListNode([]int{1})
	want3 := []int{1}
	got3 := listNodeToArray(swapPairs(head3))
	if sliceEqual(got3, want3) {
		fmt.Printf("PASS: 测试用例 3\n")
	} else {
		fmt.Printf("FAIL: 测试用例 3, want=%v, got=%v\n", want3, got3)
	}
	fmt.Println("测试结束。")
}
