package main

//定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。
//
// 示例:
// 输入: 1->2->3->4->5->NULL
//输出: 5->4->3->2->1->NULL
// 注意：本题与主站 206 题相同：https://leetcode-cn.com/problems/reverse-linked-list/

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead, _ := reverseReclusive(head)
	return newHead
}

func reverseReclusive(head *ListNode) (*ListNode, *ListNode) {
	if head == nil {
		return nil, nil
	}
	newHead, tail := reverseReclusive(head.Next)
	if tail != nil {
		tail.Next = head
	}
	if newHead == nil {
		newHead = head
	}
	head.Next = nil
	return newHead, head
}

//leetcode submit region end(Prohibit modification and deletion)
