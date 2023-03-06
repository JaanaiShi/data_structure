package main

import "fmt"

type MyStack struct {
	Queue1 []int
	Queue2 []int
}

func Constructor() MyStack {
	return MyStack{
		Queue1: make([]int, 0), // 输入队列
		Queue2: make([]int, 0), // 输出队列
	}
}

func (this *MyStack) Push(x int) {
	this.Queue1 = append(this.Queue1, x)
}

func (this *MyStack) Pop() int {
	var (
		num int
	)

	queue1Length := len(this.Queue1)

	// 从队列1中取出元素
	for i := 0; i < queue1Length; i++ {
		this.Queue2 = append(this.Queue2, this.Queue1[i])
		if i == queue1Length-1 {
			num = this.Queue1[i]
		}
	}

	this.Queue1 = make([]int, 0)

	// 将队列2的元素放到队列1中
	for i := 0; queue1Length > 1 && i < queue1Length-1; i++ {
		this.Queue1 = append(this.Queue1, this.Queue2[i])
	}

	this.Queue2 = make([]int, 0)

	return num
}

func (this *MyStack) Top() int {
	var (
		num int
	)

	queue1Length := len(this.Queue1)

	// 从队列1中取出元素
	for i := 0; i < queue1Length; i++ {
		this.Queue2 = append(this.Queue2, this.Queue1[i])
		if i == queue1Length-1 {
			num = this.Queue1[i]
		}
	}

	this.Queue1 = make([]int, 0)

	// 将队列2的元素放到队列1中
	for i := 0; i < queue1Length; i++ {
		this.Queue1 = append(this.Queue1, this.Queue2[i])
	}

	this.Queue2 = make([]int, 0)

	return num
}

func (this *MyStack) Empty() bool {
	if len(this.Queue1) == 0 {
		return true
	} else {
		return false
	}
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

func main() {
	obj := Constructor()
	obj.Push(12)
	param_2 := obj.Pop()
	param_3 := obj.Top()
	param_4 := obj.Empty()

	fmt.Println("param_2", param_2)
	fmt.Println("param_3", param_3)
	fmt.Println("param_4", param_4)

}
