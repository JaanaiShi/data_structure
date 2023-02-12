package main

import "leetcode/common"


func swapPairs(head *common.ListNode) *common.ListNode {
	var (
		flag  bool = true
	)

	newHead := &common.ListNode{
		Next: head,
	}

	pre, cur := newHead, head

	for cur != nil {
		if flag {
			next := cur.Next
			if next != nil {
				// 交换节点
				cur.Next = next.Next
				next.Next = cur
				
				pre.Next = next
				
				// 定义当前节点
				cur = next
			}
			flag = false
		}else {
			pre = cur
			flag = true
		}

		cur = cur.Next
	}

	return newHead.Next

}

func main() {
    head := common.ListNode{
        Val: 1,
        Next: &common.ListNode{
            Val: 2,
            Next: &common.ListNode{
                Val: 3,
                Next: &common.ListNode{
					Val: 4,
					Next: nil,
				},
            },
        },
    }

    common.PrintLink(swapPairs(&head))
}