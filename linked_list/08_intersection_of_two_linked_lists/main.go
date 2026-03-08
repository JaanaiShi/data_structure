package main

import "fmt"

// ListNode 链表节点定义
type ListNode struct {
	Val  int
	Next *ListNode
}

// getIntersectionNode 链表相交
// 给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在交点，返回 nil 。
//
// 示例 1：
// 输入：intersectVal = 8, listA = [4,1,8,4,5], listB = [5,6,1,8,4,5], skipA = 2, skipB = 3
// 输出：Intersected at '8'
//
// 示例 2：
// 输入：intersectVal = 2, listA = [1,9,1,2,4], listB = [3,2,4], skipA = 3, skipB = 1
// 输出：Intersected at '2'
//
// 示例 3：
// 输入：intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2
// 输出：null
//
// LeetCode 160: https://leetcode.cn/problems/intersection-of-two-linked-lists/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// TODO: 请实现解法
	return nil
}

func main() {
	fmt.Println("开始测试 linked_list/08_intersection_of_two_linked_lists...")

	// 构造测试用例 1
	// listA = [4,1,8,4,5], listB = [5,6,1,8,4,5]
	common := &ListNode{Val: 8, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}
	headA1 := &ListNode{Val: 4, Next: &ListNode{Val: 1, Next: common}}
	headB1 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 1, Next: common}}}
	got1 := getIntersectionNode(headA1, headB1)
	if got1 == common {
		fmt.Printf("PASS: 测试用例 1\n")
	} else {
		fmt.Printf("FAIL: 测试用例 1, want=%v, got=%v\n", common, got1)
	}

	// 构造测试用例 2
	// 不相交
	headA3 := &ListNode{Val: 2, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	headB3 := &ListNode{Val: 1, Next: &ListNode{Val: 5}}
	got3 := getIntersectionNode(headA3, headB3)
	if got3 == nil {
		fmt.Printf("PASS: 测试用例 2\n")
	} else {
		fmt.Printf("FAIL: 测试用例 2, want=nil, got=%v\n", got3)
	}

	fmt.Println("测试结束。")
}
