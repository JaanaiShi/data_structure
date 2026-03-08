package main

import (
	"fmt"
)

// ListNode 链表节点定义
type ListNode struct {
	Val  int
	Next *ListNode
}

// removeNthFromEnd 删除链表中倒数第 n 个节点，返回链表头节点。
// 对应 LeetCode 19. Remove Nth Node From End of List
//
// 示例 1：[1,2,3,4,5], n=2 -> [1,2,3,5]
// 示例 2：[1],         n=1 -> []
// 示例 3：[1,2],       n=1 -> [1]
//
// 提示：
//   - 链表中节点的数目为 sz，1 <= sz <= 30
//   - 0 <= Node.val <= 100
//   - 1 <= n <= sz
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 构建一个空头
	root := &ListNode{
		Val:  0,
		Next: head,
	}
	pre, low, fast := root, root, root
	for fast != nil {
		fast = fast.Next
		if n > 0 {
			n--
			continue
		}
		pre = low
		low = low.Next
	}
	pre.Next = low.Next
	return root.Next
}

// -------------------- 以下为辅助函数，无需修改 --------------------

// arrayToListNode 将整数切片转换为链表，返回头节点。
func arrayToListNode(nums []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for _, v := range nums {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

// listNodeToArray 将链表转换为整数切片，方便断言比较。
func listNodeToArray(head *ListNode) []int {
	result := []int{}
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

// sliceEqual 判断两个整数切片是否相等。
func sliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	// 测试用例
	tests := []struct {
		nums []int
		n    int
		want []int
	}{
		// 题目示例
		{[]int{1, 2, 3, 4, 5}, 2, []int{1, 2, 3, 5}},
		{[]int{1}, 1, []int{}},
		{[]int{1, 2}, 1, []int{1}},
		// 边界：删除头节点
		{[]int{1, 2}, 2, []int{2}},
		// 删除尾节点
		{[]int{1, 2, 3}, 1, []int{1, 2}},
	}

	fmt.Println("开始测试 linked_list/03_remove_nth_from_end...")
	for _, tt := range tests {
		head := arrayToListNode(tt.nums)
		got := listNodeToArray(removeNthFromEnd(head, tt.n))
		if sliceEqual(got, tt.want) {
			fmt.Printf("PASS: nums=%v, n=%d -> %v\n", tt.nums, tt.n, got)
		} else {
			fmt.Printf("FAIL: nums=%v, n=%d, want=%v, got=%v\n", tt.nums, tt.n, tt.want, got)
		}
	}
	fmt.Println("测试结束。")
}
