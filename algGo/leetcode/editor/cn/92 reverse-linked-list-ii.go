package main

import (
	"fmt"
)

//给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链
//表节点，返回 反转后的链表 。
//
// 示例 1： 
//输入：head = [1,2,3,4,5], left = 2, right = 4
//输出：[1,4,3,2,5]
//
// 示例 2： 
//输入：head = [5], left = 1, right = 1
//输出：[5]

type ListNode struct {
	Val  int
	Next *ListNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return nil
	}

	if left == right {
		return head
	}

	startNode, endNode := head, head
	if left == 1 {
		fakeNode := &ListNode{Next: head}

		startNode, endNode = fakeNode, head
		head = fakeNode
	}

	for i := 1; i < left-1; i++ {
		startNode = startNode.Next
	}

	for i := 1; i < right; i++ {
		endNode = endNode.Next
	}

	tailNode := endNode.Next
	endNode.Next = nil
	startNode.Next = reverseLinkList(startNode.Next)
	for startNode.Next != nil {
		startNode = startNode.Next
	}
	startNode.Next = tailNode

	if head.Val == 0 {
		return head.Next
	}
	return head
}

func reverseLinkList(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}

	node := reverseLinkList(head.Next)
	head.Next.Next = head
	head.Next = nil

	return node
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(reverseBetween(&ListNode{
		Val: 3,
		Next: &ListNode{
			Val:  5,
			Next: nil,
		},
	}, 1, 2))
}
