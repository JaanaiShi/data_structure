package main

import "fmt"

type Node struct {
     Val int
     Children []*Node
}

   
func maxDepth(root *Node) int {
	var (
		result []int
	)
	if root == nil {
		return 0
	}

	for i:=0; i < len(root.Children); i++ {
		result = append(result, maxDepth(root.Children[i]))
	}

	return  max(result) + 1

}

func maxDepth1(root *Node) int {
	var (
		depth int
		queue []*Node  // 层次遍历 使用队列
	)
	if root == nil {
		return 0
	}
	queue = append(queue, root)

	for len(queue) > 0 {
		depth += 1
		curLen := len(queue)
		for i:=0; i < curLen; i++ {
			nodeList := queue[i].Children
			for _, child := range nodeList{
				queue = append(queue, child)
			}
		}
		queue = queue[curLen:]
	}
	return depth
}

func max(sliceNum []int) int {

	var (
		maxNum = 0
	)

	for i:=0; i < len(sliceNum); i++ {
		if sliceNum[i] > maxNum {
			maxNum = sliceNum[i]
		}
	}

	return maxNum
}


func main() {
	root := &Node{
		Val: 1,
		Children: []*Node{
			{
				Val: 2,
				Children: []*Node{
					{
						Val: 1,
					},
				},
			},
		},
	}

	fmt.Println(maxDepth1(root))
}