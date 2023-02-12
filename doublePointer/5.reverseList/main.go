package main

import "fmt"

// Definition for singly-linked list.
 type ListNode struct {
     Val int
     Next *ListNode
 }
 
func reverseList(head *ListNode) *ListNode {
    var (
        result *ListNode
        slowNode = head
        flag = true
    )
    for head != nil {
        newHead := &ListNode{}
        if slowNode != head {
            newHead = head.Next
            head.Next = slowNode
            if flag{
                slowNode.Next = nil
                flag = false
            }
            slowNode = head
        }else {
            newHead = head.Next
        }
        if  newHead == nil {
            result = head
            break
        }

        head = newHead

    }
    return result
}

func reverseList2(head *ListNode) *ListNode {
    var (
        slowNode *ListNode   // 定义一个变量
    )

    for head != nil {
        node := head
        nextNode := head.Next
        head.Next = slowNode
        slowNode = node
        head  = nextNode
    }

    return slowNode
}

func PrintLink(head *ListNode) {
    for head != nil {
        fmt.Println(head.Val)
        head = head.Next
    }
}


func main() {
    head := ListNode{
        Val: 1,
        Next: &ListNode{
            Val: 2,
            Next: &ListNode{
                Val: 3,
                Next: nil,
            },
        },
    }

    PrintLink(reverseList2(&head))
}