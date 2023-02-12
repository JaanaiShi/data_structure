package common

import "fmt"



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