package main

import "fmt"

// MyLinkedList 设计链表
// 对应 LeetCode 707. Design Linked List
//
// TODO: 请实现双链表或单链表
type MyLinkedList struct {
	// TODO: 实现必要字段
}

// Constructor 初始化 MyLinkedList 对象。
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

// Get 获取链表中第 index 个节点的值。如果索引无效，则返回-1。
func (this *MyLinkedList) Get(index int) int {
	// TODO: 请实现解法
	return -1
}

// AddAtHead 在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
func (this *MyLinkedList) AddAtHead(val int) {
	// TODO: 请实现解法
}

// AddAtTail 将值为 val 的节点追加到链表的最后一个元素。
func (this *MyLinkedList) AddAtTail(val int) {
	// TODO: 请实现解法
}

// AddAtIndex 在链表中的第 index 个节点之前添加值为 val  的节点。
// 如果 index 等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	// TODO: 请实现解法
}

// DeleteAtIndex 如果索引 index 有效，则删除链表中的第 index 个节点。
func (this *MyLinkedList) DeleteAtIndex(index int) {
	// TODO: 请实现解法
}

func main() {
	fmt.Println("开始测试 linked_list/05_design_linked_list...")

	list := Constructor()
	list.AddAtHead(1)
	list.AddAtTail(3)
	list.AddAtIndex(1, 2) // 链表变为 1->2->3
	val1 := list.Get(1)   // 返回 2
	list.DeleteAtIndex(1) // 现在链表是 1->3
	val2 := list.Get(1)   // 返回 3

	if val1 == 2 && val2 == 3 {
		fmt.Println("PASS: 测试用例 1")
	} else {
		fmt.Printf("FAIL: 测试用例 1, get(1)=%d(want 2), get(1)=%d(want 3)\n", val1, val2)
	}

	fmt.Println("测试结束。")
}
