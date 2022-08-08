package utils

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// PrintLink 打印链表
func PrintLink(head *ListNode) {
	if head == nil {
		return
	}
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

// InitLink 使用arr初始化链表
func InitLink(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	var head, node *ListNode
	for _, v := range arr {
		if head == nil {
			head = &ListNode{
				Val:  v,
				Next: nil,
			}
			node = head
			continue
		}
		node.Next = &ListNode{
			Val:  v,
			Next: nil,
		}
		node = node.Next
	}
	return head
}