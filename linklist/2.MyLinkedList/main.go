package main

import (
	"fmt"
)

/*
设计链表的实现。您可以选择使用单链表或双链表。单链表中的节点应该具有两个属性：val 和 next。val 是当前节点的值，next 是指向下一个节点的指针/引用。如果要使用双向链表，则还需要一个属性 prev 以指示链表中的上一个节点。假设链表中的所有节点都是 0-index 的。

在链表类中实现这些功能：

get(index)：获取链表中第 index 个节点的值。如果索引无效，则返回-1。
addAtHead(val)：在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
addAtTail(val)：将值为 val 的节点追加到链表的最后一个元素。
addAtIndex(index,val)：在链表中的第 index 个节点之前添加值为 val  的节点。如果 index 等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
deleteAtIndex(index)：如果索引 index 有效，则删除链表中的第 index 个节点。


MyLinkedList linkedList = new MyLinkedList();
linkedList.addAtHead(1);
linkedList.addAtTail(3);
linkedList.addAtIndex(1,2);   //链表变为1-> 2-> 3
linkedList.get(1);            //返回2
linkedList.deleteAtIndex(1);  //现在链表是1-> 3
linkedList.get(1);            //返回3

*/

type Node struct {
	Val int
	Next *Node
}

type MyLinkedList struct {
	Link *Node
	Size int
}


func Constructor() MyLinkedList {
	return MyLinkedList{
		Link: &Node{},
		Size: 0,
	}
}

// 构建虚拟节点
func (this *MyLinkedList) Get(index int) int {
	var (
		count = 0
		head = this.Link.Next // 头节点
	)

	if this.Size == 0 {
		return -1
	}

	if index < 0 || index > this.Size {
		return -1
	}


	// 遍历节点
	for head != nil {
		if index == count {
			return head.Val
		}

		head = head.Next
		count++
	}
	return -1
}


func (this *MyLinkedList) AddAtHead(val int)  {
	var (
		head = this.Link.Next
	)

	node := &Node{
		Val: val,
		Next: head,
	}

	this.Link.Next = node
	this.Size++
}


func (this *MyLinkedList) AddAtTail(val int)  {
	var (
		head = this.Link.Next
	)

	node := &Node{
		Val: val,
		Next: nil,
	}

	if head == nil {
		this.Link.Next = node
		this.Size++
		return
	}

	for head != nil {
		if head.Next == nil {
			head.Next = node
			this.Size++
			break
		}
		head = head.Next
	}
}


func (this *MyLinkedList) AddAtIndex(index int, val int)  {
	var (
		pre = this.Link
		head = this.Link.Next
		count = 0
	)

	node := &Node{
		Val: val,
		Next: nil,
	}

	// index小于0
	if index <= 0 {
		this.AddAtHead(val)
		return
	}

	// index长度等于节点个数
	if index == this.Size {
		this.AddAtTail(val)
		return
	}


	// 头部节点为空
	if head == nil && index == 0 {
		this.Link.Next = node
		this.Size++
		return
	}

	for head != nil {
		if index == count {
			node.Next = head
			pre.Next = node
			this.Size++
			break
		}
		pre = pre.Next
		head = head.Next
		count++
	}

}


func (this *MyLinkedList) DeleteAtIndex(index int)  {
	var (
		head = this.Link.Next
		pre = this.Link
		count = 0
	)



	for head != nil {
		if index == count {
			pre.Next = head.Next
			this.Size--
			break
		}
		pre = head
		head = head.Next
		count++
	}
}

func (this *MyLinkedList) ForEach() []int {
	var (
		head = this.Link.Next
		result []int
	)

	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}

	return result
}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

 func testDelete() {
	obj := Constructor()
	obj.AddAtHead(2);
	obj.AddAtHead(4)
	obj.AddAtTail(3);
	obj.AddAtTail(4)
	fmt.Println(obj.ForEach())

	// 删除开始节点
	obj.DeleteAtIndex(0)
	fmt.Println(obj.ForEach())

	obj.DeleteAtIndex(obj.Size - 1)
	fmt.Println(obj.ForEach())

	obj.DeleteAtIndex(1)
	fmt.Println(obj.ForEach())
 }

 func testGet() {
	obj := Constructor()
	obj.AddAtHead(2);
	obj.AddAtHead(4)
	obj.AddAtTail(3);
	obj.AddAtTail(4)
	fmt.Println(obj.ForEach())

	fmt.Println(obj.Get(0))
	fmt.Println(obj.Get(obj.Size - 1))
	fmt.Println(obj.Get(1))
 }

 func testAddAtIndex() {
	fmt.Printf("\n测试指定位置添加\n")
	obj := Constructor()
	obj.AddAtIndex(0, 1)
	obj.AddAtIndex(0, 2)
	fmt.Println(obj.ForEach())
 }

 func testFatul() {
	obj := Constructor()
	obj.AddAtHead(7)
	obj.AddAtHead(2)
	obj.AddAtHead(1)
	obj.AddAtIndex(3, 0)
	fmt.Println(obj.ForEach())
	obj.DeleteAtIndex(2)
	obj.AddAtHead(6)
	obj.AddAtTail(4)
	fmt.Println(obj.ForEach())
	fmt.Println(obj.Get(4))

	obj.AddAtHead(4)
	obj.AddAtIndex(5,0)
	obj.AddAtHead(6)
 }

 func testFatul1() {
	obj := Constructor()
	obj.AddAtIndex(0, 10)
	obj.AddAtIndex(0, 20)
	obj.AddAtIndex(1, 30)
	fmt.Println(obj.ForEach())
	fmt.Println(obj.Get(0))

 }

 func testFatal2() {
	obj := Constructor()
	obj.AddAtTail(1)
	obj.AddAtTail(3)
	obj.AddAtIndex(1, 2)
	fmt.Println(obj.ForEach())
	fmt.Println(obj.Get(1))
	obj.DeleteAtIndex(1)
	fmt.Println(obj.ForEach())
	fmt.Println(obj.Get(2))
 }

 func testFatal3() {
	obj := Constructor()
	obj.AddAtIndex(1, 0)
	fmt.Println(obj.ForEach())
	fmt.Println(obj.Get(0))

 }

 func main() {
	fmt.Println("测试删除")
	testDelete()

	fmt.Printf("\n测试Get\n")
	testGet()

	testFatal3()
 }
