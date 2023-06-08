package common

import "fmt"


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

type ListNode struct {
	Val int
	Next *ListNode
}

func PrintLink(head *ListNode) {
    for head != nil {
        fmt.Println(head.Val)
        head = head.Next
    }
}


func PreOrder(root *TreeNode) {

    if root == nil {
        return 
    }

    fmt.Println(root.Val)

    PreOrder(root.Left)

    PreOrder(root.Right)


}