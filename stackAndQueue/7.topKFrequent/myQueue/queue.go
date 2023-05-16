package myqueue


/*
	单调队列被称为“有序队列”
	要求的是每连续的k个数中的最大（最小）值，很明显，当一个数进入所要“寻找”最大值的范围中时，若这个数比其前面（先进队）的数要大，显然，前面的数会比这个数先出对且不再可能是
	最大值。
*/

type MyQueue struct {
	queue []int
}

func NewMyQueue() *MyQueue {
	return &MyQueue{
		queue: make([]int, 0),
	}
}

// Pop调用que.pop(滑动窗口中移除元素的数值)
func (q *MyQueue) Pop(value int) {
	if !q.Empty() && q.Front() == value {
		q.queue = q.queue[1:]
	}
}

func (q *MyQueue) Back() int {
	return q.queue[q.Length() - 1]
}



// Push 滑动窗口添加元素的数值
func (q *MyQueue) Push(value int) {
	for !q.Empty() &&  value > q.Back() {
		q.queue = q.queue[:len(q.queue) - 1]
	}

	q.queue = append(q.queue, value)
}

func (q *MyQueue) Empty() bool {
	if len(q.queue) > 0 {
		return false
	} else {
		return true
	}
}

// Front 返回我们要的最大值
func (q *MyQueue) Front() int {
	return q.queue[0]

}

func (q *MyQueue) Length() int {
	return len(q.queue)
}
