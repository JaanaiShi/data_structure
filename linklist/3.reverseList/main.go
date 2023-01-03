package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
  输入：head = [1,2,3,4,5]
  输出：[5,4,3,2,1]

  输入：head = [1,2]
  输出：[2,1]

  输入：head = []
  输出：[]
*/

type ListNode struct {
    Val int
	Next *ListNode
}

 func reverseList(head *ListNode) *ListNode {
	//执行用时：0 ms  内存消耗：2.7 MB
	var (
		newLink *ListNode = &ListNode{}  // 虚拟节点
		nodeArray []*ListNode
	)

	for head != nil {
		node := ListNode{
			Val: head.Val,
			Next: nil,
		}
		nodeArray = append(nodeArray, &node)
		head = head.Next
	}
	newHead := newLink
	for i:=len(nodeArray) - 1; i >= 0; i-- {
		node := nodeArray[i]
		newHead.Next = node
		newHead = newHead.Next
	}

	return newLink.Next
 }

func reverseList1(head *ListNode) *ListNode {
	// 改变链条的转向
	var (
		pre = &ListNode{}  // 虚拟的头节点
		newHead = head
		flag   bool
	)
	
	if newHead == nil {
		return nil
	}

	for newHead != nil {
		// 1. 保留当前节点
		node := newHead
		// 2. 将循环变量指向下一个node
		newHead = newHead.Next
		// 3. 将当前node的Next指向上一个node
		node.Next = pre
		// 4. 指定上一个节点为当前节点
		pre = node
		// 5. 去掉头部虚拟节点(只执行一次)
		if !flag {
			pre.Next = nil
			flag = true
		}
	}
	result := pre
	// 去掉虚拟节点

	return result

}

// 官方-双指针法
func reverseList2(head *ListNode) *ListNode {
	var pre *ListNode    // 只是声明了变量，没有进行初始化，所以为nil
	cur := head
	for cur != nil {
		// 保存下一个节点的地址
		next := cur.Next
		// 反转链表，当前节点的
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func recurList(cur, new *ListNode) *ListNode {
	// 退出条件
	if cur.Next == nil {
		return new
	}

	// 赋值
	node := cur.Next
	cur.Next = new
	recurList(node, cur) 

	return cur

}

func reverseList3(head *ListNode) *ListNode {	
	var new *ListNode

	new = recurList(head, new)

	return new
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

 

 func main() {
	link := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: nil,
					},
				},
			},
		},
	}

	ForEach(reverseList3(link))
 }