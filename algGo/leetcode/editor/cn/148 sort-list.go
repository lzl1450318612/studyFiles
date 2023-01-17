package main

//给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
//
// 示例 1：
//输入：head = [4,2,1,3]
//输出：[1,2,3,4]
//
// 示例 2：
//输入：head = [-1,5,3,4,0]
//输出：[-1,0,3,4,5]
//
// 示例 3：
//输入：head = []
//输出：[]
//
// 进阶：你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	return splitMerge(head)
}

func splitMerge(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head
	var pre *ListNode
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	pre.Next = nil

	l := splitMerge(head)
	r := splitMerge(slow)

	return mergeList(l, r)
}

func mergeList(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var resHead *ListNode
	var resCur *ListNode

	if l1.Val <= l2.Val {
		resHead, resCur = l1, l1
		l1 = l1.Next
	} else {
		resHead, resCur = l2, l2
		l2 = l2.Next
	}

	node1, node2 := l1, l2

	for node1 != nil && node2 != nil {
		if node1.Val <= node2.Val {
			resCur.Next = node1
			node1 = node1.Next
		} else {
			resCur.Next = node2
			node2 = node2.Next
		}
		resCur = resCur.Next
	}

	if node1 != nil {
		resCur.Next = node1
	}

	if node2 != nil {
		resCur.Next = node2
	}

	return resHead
}

//leetcode submit region end(Prohibit modification and deletion)
