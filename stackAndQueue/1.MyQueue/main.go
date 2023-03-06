package main

import (
	"fmt"
)

type MyQueue struct {
	Input  []int // 输入栈
	OutPut []int // 输出栈
}

func Constructor() MyQueue {
	return MyQueue{
		Input:  make([]int, 0),
		OutPut: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.Input = append(this.Input, x)
}

func (this *MyQueue) Pop() int {

	if this.Empty() {
		return -1
	}

	inputLen := len(this.Input)

	// 将输入栈的内容放到输出栈中
	for i := inputLen - 1; i >= 0; i-- {
		this.OutPut = append(this.OutPut, this.Input[i])
	}
	this.Input = make([]int, 0)
	// 取出需要的数
	num := this.OutPut[inputLen-1]

	// 将输出栈的内容重新放回到输入栈中
	for i := inputLen - 2; i >= 0; i-- {
		this.Input = append(this.Input, this.OutPut[i])
	}
	

	this.OutPut = make([]int, 0)

	return num
}

func (this *MyQueue) Peek() int {

	if this.Empty() {
		return -1
	}

	inputLen := len(this.Input)

	// 将输入栈的内容放到输出栈中
	for i := inputLen - 1; i >= 0; i-- {
		this.OutPut = append(this.OutPut, this.Input[i])
	}
	this.Input = make([]int, 0)
	// 取出需要的数
	num := this.OutPut[inputLen-1]

	// 将输出栈的内容重新放回到输入栈中
	for i := inputLen - 1; i >= 0; i-- {
		this.Input = append(this.Input, this.OutPut[i])
	}
	

	this.OutPut = make([]int, 0)

	return num
}

func (this *MyQueue) Empty() bool {
	if len(this.Input) == 0 {
		return true
	} else {
		return false
	}
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

func main() {
	obj := Constructor();
	obj.Push(1);
	param_3 := obj.Peek();
	param_4 := obj.Empty();
	fmt.Println("param_3", param_3)
	fmt.Println("param_4", param_4)

}
