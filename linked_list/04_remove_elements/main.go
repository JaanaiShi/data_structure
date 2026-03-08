package main

import "fmt"

// ListNode 链表节点定义
type ListNode struct {
	Val  int
	Next *ListNode
}

// removeElements 移除链表元素
// 给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。
//
// 示例 1：
// 输入：head = [1,2,6,3,4,5,6], val = 6
// 输出：[1,2,3,4,5]
//
// 示例 2：
// 输入：head = [], val = 1
// 输出：[]
//
// 示例 3：
// 输入：head = [7,7,7,7], val = 7
// 输出：[]
//
// LeetCode 203: https://leetcode.cn/problems/remove-linked-list-elements/
func removeElements(head *ListNode, val int) *ListNode {
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
	fmt.Println("开始测试 linked_list/04_remove_elements...")

	head1 := arrayToListNode([]int{1, 2, 6, 3, 4, 5, 6})
	want1 := []int{1, 2, 3, 4, 5}
	got1 := listNodeToArray(removeElements(head1, 6))
	if sliceEqual(got1, want1) {
		fmt.Printf("PASS: 测试用例 1\n")
	} else {
		fmt.Printf("FAIL: 测试用例 1, want=%v, got=%v\n", want1, got1)
	}

	var head2 *ListNode
	var want2 []int
	got2 := listNodeToArray(removeElements(head2, 1))
	if sliceEqual(got2, want2) {
		fmt.Printf("PASS: 测试用例 2\n")
	} else {
		fmt.Printf("FAIL: 测试用例 2, want=%v, got=%v\n", want2, got2)
	}

	head3 := arrayToListNode([]int{7, 7, 7, 7})
	var want3 []int
	got3 := listNodeToArray(removeElements(head3, 7))
	if sliceEqual(got3, want3) {
		fmt.Printf("PASS: 测试用例 3\n")
	} else {
		fmt.Printf("FAIL: 测试用例 3, want=%v, got=%v\n", want3, got3)
	}
	fmt.Println("测试结束。")
}
