package main

// Definition for singly-linked list.
type ListNode struct {
	Val int
	Next *ListNode
}
 
func detectCycle(head *ListNode) *ListNode {
	var (
		newHead = head
		nodeMap = make(map[*ListNode]struct{}, 0)
	)

	for newHead != nil {
		if _, ok := nodeMap[newHead]; ok {
			return newHead
		}else {
			nodeMap[newHead] = struct{}{}
		}
		newHead = newHead.Next
	}
	return nil
}

func detectCycle1(head *ListNode) *ListNode {
	var (
		newHead = head
		nodeMap = make(map[*ListNode]struct{}, 0)
	)

	for newHead != nil {
		if _, ok := nodeMap[newHead]; ok {
			return newHead
		}else {
			nodeMap[newHead] = struct{}{}
		}
		newHead = newHead.Next
	}
	return nil
}

func main() {
	head := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 4,
					Next: ,
				},
			},
		},
	}
}