package main

import "leetcode/common"

// func detectCycle(head *common.ListNode) *common.ListNode {
//     var (
// 		fastNode *common.ListNode
// 	)

// 	demHead := &common.ListNode{
// 		Next: head,
// 	}

// 	for demHead != nil {
// 		fastNode  = demHead.Next
// 		for fastNode != nil {
// 			if demHead == fastNode {
// 				return demHead
// 			}
// 			fastNode = fastNode.Next
// 		}
// 		demHead = demHead.Next
		
// 	}

// 	return nil
// }


func detectCycle(head *common.ListNode) *common.ListNode {
	slow , fast := head, head
	// 第一个for循换判断是否是一个环
	for fast != nil && fast.Next != nil {
		slow = slow.Next   // slow 走一步
		fast = fast.Next.Next // fast 走两步

		if slow == fast {
			// 第二个for循环 通过规律，寻找换的入口
			for slow != head {
				slow = slow.Next
				head = head.Next
			}
			return head
		}
	}

	return nil
}



func main() {
	cycleNode := &common.ListNode{
		Val: 11,
		Next: nil,
	}
	head := &common.ListNode{
		Val: 1,
		Next: &common.ListNode{
			Val: 2,
			Next: &common.ListNode{
				Val: 3,
				Next: nil,
			},
		},
	}

	commonNodeList := &common.ListNode{
		Val: 8,
		Next: &common.ListNode{
			Val: 9,
			Next: &common.ListNode{
				Val: 10,
				Next: cycleNode,
			},
		},
	}

	cycleNode.Next = commonNodeList

	head.Next.Next = cycleNode

	common.PrintLink(detectCycle(head))
}