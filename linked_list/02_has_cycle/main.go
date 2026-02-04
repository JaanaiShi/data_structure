package main

import (
	"fmt"
)

// ListNode 定义链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// hasCycle 判断链表是否有环
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for slow != fast {
		if fast.Next == nil || fast.Next.Next == nil {
			return false
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return true
}

func hasCycleV2(head *ListNode) bool {
	nodeMap := make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := nodeMap[head]; ok {
			return true
		} else {
			nodeMap[head] = struct{}{}
		}
		head = head.Next
	}
	return false
}

func main() {
	// 构造测试用例 1: 无环链表 1->2->3->4
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4

	// 构造测试用例 2: 有环链表 1->2->3->4->2...
	node1_cycle := &ListNode{Val: 1}
	node2_cycle := &ListNode{Val: 2}
	node3_cycle := &ListNode{Val: 3}
	node4_cycle := &ListNode{Val: 4}
	node1_cycle.Next = node2_cycle
	node2_cycle.Next = node3_cycle
	node3_cycle.Next = node4_cycle
	node4_cycle.Next = node2_cycle // 环

	// 构造测试用例 3: 单节点无环
	node1_single := &ListNode{Val: 1}

	tests := []struct {
		name string
		head *ListNode
		want bool
	}{
		{"No Cycle", node1, false},
		{"Has Cycle", node1_cycle, true},
		{"Single Node", node1_single, false},
		{"Nil Node", nil, false},
	}

	fmt.Println("开始测试 linked_list/02_has_cycle...")
	for _, tt := range tests {
		got := hasCycle(tt.head)
		if got == tt.want {
			fmt.Printf("PASS: %s, want=%v, got=%v\n", tt.name, tt.want, got)
		} else {
			fmt.Printf("FAIL: %s, want=%v, got=%v\n", tt.name, tt.want, got)
		}
	}
	fmt.Println("测试结束。")
}
