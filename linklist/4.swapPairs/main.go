package main

import (
	"fmt"
)

/*
	给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值
的情况下完成本题（即，只能进行节点交换）。
输入：head = [1,2,3,4]
输出：[2,1,4,3]

输入：head = []
输出：[]

输入：head = [1]
输出：[1]
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val int
	Next *ListNode
}

// 交换值，不符合题意
 func swapPairs(head *ListNode) *ListNode {
	var (
		result *ListNode = head
		newHead *ListNode = head
		pre *ListNode = head
		flag  bool
	)

	for newHead != nil {
		if flag {
			pre.Val, newHead.Val  = newHead.Val, pre.Val
			flag = false
		}else {
			flag = true
		}
		pre = newHead
		newHead = newHead.Next
	}

	return result
 }

 func swapPairs1(head *ListNode) *ListNode {
	var (
		result *ListNode = head
		newHead *ListNode = head
		pre *ListNode = head
		flag  bool
	)

	for newHead != nil {
		if flag {
			pre.Val, newHead.Val  = newHead.Val, pre.Val
			flag = false
		}else {
			flag = true
		}
		pre = newHead
		newHead = newHead.Next
	}

	return result
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
					Next: nil,
				},
			},
		},
	}

	ForEach(swapPairs(link))


 }