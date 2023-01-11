package main

import (
	"fmt"
)

//Definition for singly-linked list.
type ListNode struct {
	Val int
	Next *ListNode
}

// 自己写的版本
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var (
		arrayA, arrayB []*ListNode
		j int
	)

	if headA == nil || headB == nil {
		return nil
	}

	for headA != nil {
		arrayA = append(arrayA, headA)

		headA = headA.Next
	}

	for headB != nil {
		arrayB = append(arrayB, headB)
		headB = headB.Next
	}
	
	if len(arrayA) > len(arrayB) {
		j = len(arrayB) - 1
		for i:= len(arrayA) - 1; i >= 0; i-- {
			if j > 0 {
				if arrayA[i] == arrayB[j] {
					
				}else {
					break
				}
			}
			j--
		}
	} else {
		j = len(arrayA) - 1
		for i:= len(arrayB) - 1; i >= 0; i-- {
			if j > 0 {
				if arrayB[i] != arrayA[j] {
					return arrayB[i+1]
				}
			}
			j--
		}
	}

	return nil
}

// 官方版本
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
    var (
        lengthA, lengthB int
		count int
    )
	// 计算长度
	newHeadA := headA
	for newHeadA != nil {
		lengthA++
		newHeadA = newHeadA.Next
	}
	newHeadB := headB
	for newHeadB != nil {
		lengthB++
		newHeadB = newHeadB.Next
	}

	if lengthA > lengthB {
		diff := lengthA - lengthB
		for headA != nil {
			if count == diff {
				if headA != headB {
					headB = headB.Next
				}else {
					return headA
				}
			}else {
				count++
			}
			
			headA = headA.Next
		}
	}else {
		diff := lengthB - lengthA
		for headB != nil {
			if count == diff {
				if headB != headA {
					headA = headA.Next
				}else {
					return headB
				}
			}else {
				count++
			}
			
			headB = headB.Next
		}
	}

	return nil

}

func main() {
	common := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: nil,
		},
	}
	headA := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 1,
				Next: common,
			},
		},
	}
	
	headB := &ListNode{
		Val: 3,
		Next: common,
	}
	

	node := getIntersectionNode1(headB, headA)

	fmt.Println(node.Val)
}