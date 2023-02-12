package main

import (
	"leetcode/common"
)

func removeNthFromEnd(head *common.ListNode, n int) *common.ListNode {
	var (
		length, count = 0, 0
	)
	lenHead := head
	// 1. 获取长度
	for lenHead != nil {
		lenHead = lenHead.Next
		length++
	}
	// 2. 删除节点
	// 增加虚拟节点
	newHead := &common.ListNode{
		Next: head,
	}
	pre := newHead
	result := newHead
	for newHead != nil {
		if count == length - n + 1 {
			pre.Next = newHead.Next
			break
		}
		count++
		pre = newHead
		newHead = newHead.Next
	}

	return result.Next
}

func main() {
	head := &common.ListNode{
		Val: 1,
		Next: &common.ListNode{
			Val: 2,
			Next: &common.ListNode{
				Val: 3,
				Next: &common.ListNode{
					Val: 4,
					Next: &common.ListNode{
						Val: 5,
						Next: &common.ListNode{
							Val: 6,
							Next: nil,
						},
					},
				},
			},
		},
	}

	common.PrintLink(removeNthFromEnd(head, 6))
}