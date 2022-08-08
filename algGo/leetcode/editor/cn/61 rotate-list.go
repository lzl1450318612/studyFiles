package main

import (
	"fmt"
)

//给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。
//
// 示例 1： 
//输入：head = [1,2,3,4,5], k = 2
//输出：[4,5,1,2,3]
//
// 示例 2： 
//输入：head = [0,1,2], k = 4
//输出：[2,0,1]

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

//leetcode submit region begin(Prohibit modification and deletion)
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	count := 1
	node := head

	for node.Next != nil {
		count++
		node = node.Next
	}

	if count == 1 {
		return head
	}

	node.Next = head
	targetCount := count - k%count - 1
	node = head
	for i := 0; i < targetCount; i++ {
		node = node.Next
	}
	temp := node.Next
	node.Next = nil
	return temp
}

//leetcode submit region end(Prohibit modification and deletion)

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

func main() {
	fmt.Println(rotateRight(InitLink([]int{1, 2, 3, 4, 5}), 2))
}
