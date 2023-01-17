package main

//输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。
//
// 示例1：
// 输入：1->2->4, 1->3->4
//输出：1->1->2->3->4->4
//
// 限制：
// 0 <= 链表长度 <= 1000
//
// 注意：本题与主站 21 题相同：https://leetcode-cn.com/problems/merge-two-sorted-lists/

//leetcode submit region begin(Prohibit modification and deletion)
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	newHead := &ListNode{
		Val:  -3,
		Next: nil,
	}
	node := newHead

	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			node.Next = &ListNode{Val: l2.Val}
			l2 = l2.Next
		} else {
			node.Next = &ListNode{Val: l1.Val}
			l1 = l1.Next
		}
		node = node.Next
	}

	for l1 != nil {
		node.Next = &ListNode{Val: l1.Val}
		node = node.Next
		l1 = l1.Next
	}

	for l2 != nil {
		node.Next = &ListNode{Val: l2.Val}
		node = node.Next
		l2 = l2.Next
	}

	return newHead.Next
}

//leetcode submit region end(Prohibit modification and deletion)
