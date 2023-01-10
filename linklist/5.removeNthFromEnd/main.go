package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}
 
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var (
        dammy = &ListNode{
            Next: head,
        }
		nodeArray []*ListNode
	)
	newHead := dammy
	for newHead != nil {
		nodeArray = append(nodeArray, newHead)
		newHead = newHead.Next
	}
	if n > len(nodeArray) - 1 {
		
	}else if n < 0 {
		
	}else {
		q := len(nodeArray) - n 
		if nodeArray[q].Next == nil {
			nodeArray[q-1].Next = nil
		}else {
			nodeArray[q -1].Next = nodeArray[q+1]
		}
	} 

	return dammy.Next
}

/*
	1. 
*/
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	var (
        dammy = &ListNode{
            Next: head,
        }
		count int
	)

	newHead, pre := dammy, dammy
	for newHead != nil {
		if count >= n+1 {
			pre = pre.Next
		}
		newHead = newHead.Next
		count++
	}

	pre.Next = pre.Next.Next

	

	return dammy.Next
}

func ForEach(head *ListNode) {
	var (
		result []int
	)

	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}

	fmt.Println(result)

}

 var (
	link1 = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: nil,
				},
			},
		},
	}
	link2 = &ListNode{
		Val: 1,
		Next: nil,
	}
	link3 = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: nil,
		},
	}
 )

func main() {
	ForEach(removeNthFromEnd1(link3, 1))
}