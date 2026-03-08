package main

import "fmt"

// ListNode 链表节点定义
type ListNode struct {
	Val  int
	Next *ListNode
}

// reverseList 翻转链表
// 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
//
// 示例 1：
// 输入：head = [1,2,3,4,5]
// 输出：[5,4,3,2,1]
//
// 示例 2：
// 输入：head = [1,2]
// 输出：[2,1]
//
// 示例 3：
// 输入：head = []
// 输出：[]
//
// LeetCode 206: https://leetcode.cn/problems/reverse-linked-list/
func reverseList(head *ListNode) *ListNode {
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
	fmt.Println("开始测试 linked_list/06_reverse_linked_list...")

	head1 := arrayToListNode([]int{1, 2, 3, 4, 5})
	want1 := []int{5, 4, 3, 2, 1}
	got1 := listNodeToArray(reverseList(head1))
	if sliceEqual(got1, want1) {
		fmt.Printf("PASS: 测试用例 1\n")
	} else {
		fmt.Printf("FAIL: 测试用例 1, want=%v, got=%v\n", want1, got1)
	}

	head2 := arrayToListNode([]int{1, 2})
	want2 := []int{2, 1}
	got2 := listNodeToArray(reverseList(head2))
	if sliceEqual(got2, want2) {
		fmt.Printf("PASS: 测试用例 2\n")
	} else {
		fmt.Printf("FAIL: 测试用例 2, want=%v, got=%v\n", want2, got2)
	}

	var head3 *ListNode
	var want3 []int
	got3 := listNodeToArray(reverseList(head3))
	if sliceEqual(got3, want3) {
		fmt.Printf("PASS: 测试用例 3\n")
	} else {
		fmt.Printf("FAIL: 测试用例 3, want=%v, got=%v\n", want3, got3)
	}
	fmt.Println("测试结束。")
}
