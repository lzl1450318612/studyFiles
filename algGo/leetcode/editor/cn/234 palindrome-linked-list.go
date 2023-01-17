package main

//给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。
//
// 示例 1：
//输入：head = [1,2,2,1]
//输出：true
//
// 示例 2：
//输入：head = [1,2]
//输出：false
//
// 进阶：你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next != nil {
			fast = fast.Next.Next
		} else {
			fast = fast.Next
		}
	}

	slow = reverseList(slow)
	temp := slow

	for slow != nil && head != temp {
		if head.Val != slow.Val {
			return false
		}
		head = head.Next
		slow = slow.Next
	}

	return true
}

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
