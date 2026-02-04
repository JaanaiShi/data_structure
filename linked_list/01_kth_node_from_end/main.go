package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// getKthFromEnd 返回链表的倒数第 k 个节点
func getKthFromEnd(head *ListNode, k int) int {
	pre, cur := head, head
	for cur != nil {
		if k > 0 {
			cur = cur.Next
			k--
			continue
		}
		cur = cur.Next
		pre = pre.Next
	}
	return pre.Val
}

func arrayToListNode(nums []int) *ListNode {
	root := ListNode{
		Val: nums[0],
	}
	temp := &root
	for i := 1; i < len(nums); i++ {
		node := ListNode{
			Val: nums[i],
		}
		temp.Next = &node
		temp = &node
	}
	return &root
}

func main() {
	// 测试用例
	tests := []struct {
		nums []int
		k    int
		want []int // 期望返回的节点及其后续节点的值
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{4, 5}},
		{[]int{1, 2, 3, 4, 5}, 1, []int{5}},
		{[]int{1, 2, 3, 4, 5}, 5, []int{1, 2, 3, 4, 5}},
		{[]int{1}, 1, []int{1}},
	}

	fmt.Println("开始测试 linked_list/01_kth_node_from_end...")
	for _, tt := range tests {
		listNode := arrayToListNode(tt.nums)
		fmt.Println(getKthFromEnd(listNode, tt.k))
	}
	fmt.Println("测试结束。")
}
