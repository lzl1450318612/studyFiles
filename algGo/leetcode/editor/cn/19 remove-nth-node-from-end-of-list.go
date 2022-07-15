package main

//给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
//
// 示例 1： 
//输入：head = [1,2,3,4,5], n = 2
//输出：[1,2,3,5]
//
// 示例 2： 
//输入：head = [1], n = 1
//输出：[]
// 
// 示例 3：
//输入：head = [1,2], n = 1
//输出：[1]
// 进阶：你能尝试使用一趟扫描实现吗？

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	preNode, tailNode := head, head

	for i := 0; i < n; i++ {
		if tailNode == nil {
			node := head.Next
			head.Next = nil
			head = node
			return head
		}
		tailNode = tailNode.Next
	}
	if tailNode == nil {
		node := head.Next
		head.Next = nil
		head = node
		return head
	}

	for tailNode.Next != nil {
		preNode = preNode.Next
		tailNode = tailNode.Next
	}
	if n == 1 {
		preNode.Next = nil
		return head
	}

	node := preNode.Next.Next
	preNode.Next.Next = nil
	preNode.Next = node

	return head
}

//leetcode submit region end(Prohibit modification and deletion)
