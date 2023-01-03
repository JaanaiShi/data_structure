package main

import "fmt"

/*
题意：删除链表中等于给定值 val 的所有节点。

示例 1：
输入：head = [1,2,6,3,4,5,6], val = 6
输出：[1,2,3,4,5]

示例 2：
输入：head = [], val = 1
输出：[]

示例 3：
输入：head = [7,7,7,7], val = 7
输出：[]

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	// 新定义的节点必须初始化
	var (
		newHead, pre *ListNode = head, head
	)

	if head == nil {
		return nil
	}

	for newHead != nil {
		if newHead.Val == val {
			pre.Next = newHead.Next
			newHead = newHead.Next
			
		}else {
			pre = newHead  // 保存上一个链表的地址
			if newHead.Next != nil {
				newHead = newHead.Next
			}
		}
	}
	// 处理头节点的值相等的情况
	if head.Val == val {
		head = head.Next
	}

	return head
}

// 方法二
func removeElements1(head *ListNode, val int) *ListNode {
	var(
		virtual = &ListNode{
			Next: head,
		}
	)
	pre := virtual
	newHead  := virtual.Next

	for newHead != nil {
		if newHead.Val == val {
			pre.Next = newHead.Next
		}else {
			// 将当前节点赋值给pre
			pre = newHead
		}
		newHead = newHead.Next
	}
	// 去掉虚拟节点
	return virtual.Next
}

func main() {
	// 1,2,6,3,4,5,6
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val: 5,
					Next: nil,
				},
			},
		},
	}

	// head := &ListNode{
	// 	Val: 1,
	// 	Next: nil,
	// }

	head = removeElements1(head, 0)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}


/*
curl -X 'POST' \
  'http://140.143.179.194:11113/v1/component_vulns' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "component_name": "openssl",
  "component_version": "1.1.1m",
  "package_manager": ""
}'

*/