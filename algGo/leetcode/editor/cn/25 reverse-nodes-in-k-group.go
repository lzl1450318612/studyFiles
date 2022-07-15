package main

//给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。
// k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
// 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
//
// 示例 1： 
//输入：head = [1,2,3,4,5], k = 2
//输出：[2,1,4,3,5]
// 
// 示例 2：
//输入：head = [1,2,3,4,5], k = 3
//输出：[3,2,1,4,5]
//
// 进阶：你可以设计一个只用 O(1) 额外内存空间的算法解决此问题吗？

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if k < 2 {
		return head
	}

	node, tail := head, head
	var last *ListNode

	for tail != nil && tail.Next != nil {
		for i := 0; i < k-1; i++ {
			if tail == nil||tail.Next==nil {
				return head
			}
			tail = tail.Next
		}
		temp := tail.Next
		tail.Next = nil
		node, tail = reverseLink(node)
		if last == nil {
			head = node
		} else {
			last.Next = node
		}
		tail.Next = temp
		node = temp
		last = tail
		tail = tail.Next
	}
	return head
}

func reverseLink(head *ListNode) (*ListNode, *ListNode) {
	reverseHead, reverseTail := head, head
	for reverseTail != nil && reverseTail.Next != nil {
		reverseHead, reverseTail = reverse(reverseHead, reverseTail, reverseTail.Next)
	}
	return reverseHead, reverseTail
}

func reverse(head *ListNode, tail *ListNode, reverseNode *ListNode) (*ListNode, *ListNode) {
	if reverseNode == nil {
		return head, tail
	}

	temp := reverseNode.Next
	reverseNode.Next = head
	tail.Next = temp

	return reverseNode, tail
}

//leetcode submit region end(Prohibit modification and deletion)
