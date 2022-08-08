package main

//给定一个已排序的链表的头 head ， 删除所有重复的元素，使每个元素只出现一次 。返回 已排序的链表 。
//
// 示例 1： 
//输入：head = [1,1,2]
//输出：[1,2]
//
// 示例 2： 
//输入：head = [1,1,2,3,3]
//输出：[1,2,3]

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	node := head

	for node.Next != nil {
		if node.Val == node.Next.Val {
			node.Next = node.Next.Next
		} else {
			node = node.Next
		}
	}

	return head
}

//leetcode submit region end(Prohibit modification and deletion)
