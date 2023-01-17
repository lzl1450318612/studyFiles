package main

import (
	"fmt"
)

//给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
//
// 示例 1：
//输入：head = [1,2,3,4,5]
//输出：[5,4,3,2,1]
//
// 示例 2：
//输入：head = [1,2]
//输出：[2,1]
//
// 示例 3：
//输入：head = []
//输出：[]
//
// 提示：
// 链表中节点的数目范围是 [0, 5000]
// -5000 <= Node.val <= 5000
//
// 进阶：链表可以选用迭代或递归方式完成反转。你能否用两种方法解决这道题？

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
type ListNode struct {
	Val  int
	Next *ListNode
}

//leetcode submit region begin(Prohibit modification and deletion)
//func reverseList(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return head
//	}
//	var prev *ListNode
//	node := head
//
//	for node != nil {
//		temp := node.Next
//		node.Next = prev
//		prev = node
//		node = temp
//	}
//	return prev
//}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	res, _ := reverseReclusive(head)
	return res
}

func reverseReclusive(head *ListNode) (*ListNode, *ListNode) {
	if head.Next == nil {
		return head, head
	}
	newHead, tail := reverseReclusive(head.Next)
	tail.Next = head
	head.Next = nil
	return newHead, head
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	fmt.Println(reverseList(head).Val)
}
