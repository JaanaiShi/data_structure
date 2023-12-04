package main

import (
	"leetcode/common"
)


func removeNthFromEnd(head *common.ListNode, n int) *common.ListNode {
	var (
		pre,res *common.ListNode
		count int 
	)

	newHead := &common.ListNode{
		Next: head,
	}

	pre, res = newHead, newHead

	for newHead != nil {
		if count == n + 1 {
			if newHead.Next.Next == nil {
				pre.Next = pre.Next.Next
				break	
			}
			newHead = newHead.Next
			pre = pre.Next
		}else {
			count++
			newHead = newHead.Next

		}
	}

	return res.Next
}

func main() {
	head := &common.ListNode{
		Val: 1,
		Next: &common.ListNode{
			Val: 6,
			Next: &common.ListNode{
				Val: 7,
				Next: &common.ListNode{
					Val: 5,
					Next: nil,
				},
			},
		},
	}

	common.PrintLink(removeNthFromEnd(head, 2))

}