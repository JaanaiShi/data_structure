package main

import "leetcode/common"

func getIntersectionNode(headA, headB *common.ListNode) *common.ListNode {
    
	// 1. 寻找headA，headB的长度
	// 2. 比较长度

	var (
		lengthA, lengthB int
		result, fastNode *common.ListNode
	)
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

	if lengthA >= lengthB {
		newHeadA := headA
		count := 0
		for newHeadA != nil {
			if count == lengthA - lengthB {
				fastNode = newHeadA
				break
			}
			count++
			newHeadA = newHeadA.Next
		}

		for fastNode != nil {
			if fastNode == headB {
				result = fastNode
				break
			}
			fastNode = fastNode.Next
			headB = headB.Next
		}

	}else {
		newHeadB := headB
		count := 0
		for newHeadB != nil {
			if count == lengthB - lengthA {
				fastNode = newHeadB
				break
			}
			count++
			newHeadB = newHeadB.Next
		}

		for fastNode != nil {
			if fastNode == headA {
				result = fastNode
				break
			}

			fastNode = fastNode.Next
			headA = headA.Next
		}
	}

	return result

	
}

func main() {
	commonNode := &common.ListNode{
		Val: 1,
		Next: &common.ListNode{
			Val: 2,
			Next: nil,
		},
	}

	headA := &common.ListNode{
		Val: 8,
		Next: &common.ListNode{
			Val: 7,
			Next: commonNode,
		},
	}

	headB := &common.ListNode{
		Val: 3,
		Next: commonNode,
	}

	common.PrintLink(getIntersectionNode(headA, headB))
}