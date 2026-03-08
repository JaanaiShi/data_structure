package main

import "fmt"

// ListNode 链表节点定义
type ListNode struct {
	Val  int
	Next *ListNode
}

// detectCycle 环形链表 II
// 给定一个链表的头节点  head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 nil。
//
// 示例 1：
// 输入：head = [3,2,0,-4], pos = 1 (索引从 0 开始, -4 指向 2)
// 输出：返回索引为 1 的链表节点
//
// 示例 2：
// 输入：head = [1,2], pos = 0
// 输出：返回索引为 0 的链表节点
//
// 示例 3：
// 输入：head = [1], pos = -1
// 输出：返回 nil
//
// LeetCode 142: https://leetcode.cn/problems/linked-list-cycle-ii/
func detectCycle(head *ListNode) *ListNode {
	// TODO: 请实现解法
	return nil
}

func main() {
	fmt.Println("开始测试 linked_list/09_linked_list_cycle_ii...")

	// 构造测试用例 1: 环在索引 1 的位置
	n1 := &ListNode{Val: 3}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 0}
	n4 := &ListNode{Val: -4}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n2 // 成环
	got1 := detectCycle(n1)
	if got1 == n2 {
		fmt.Printf("PASS: 测试用例 1\n")
	} else {
		fmt.Printf("FAIL: 测试用例 1, want=%v, got=%v\n", n2, got1)
	}

	// 构造测试用例 2: 环在索引 0 的位置
	n5 := &ListNode{Val: 1}
	n6 := &ListNode{Val: 2}
	n5.Next = n6
	n6.Next = n5 // 成环
	got2 := detectCycle(n5)
	if got2 == n5 {
		fmt.Printf("PASS: 测试用例 2\n")
	} else {
		fmt.Printf("FAIL: 测试用例 2, want=%v, got=%v\n", n5, got2)
	}

	// 构造测试用例 3: 无环
	n7 := &ListNode{Val: 1}
	got3 := detectCycle(n7)
	if got3 == nil {
		fmt.Printf("PASS: 测试用例 3\n")
	} else {
		fmt.Printf("FAIL: 测试用例 3, want=nil, got=%v\n", got3)
	}

	fmt.Println("测试结束。")
}
