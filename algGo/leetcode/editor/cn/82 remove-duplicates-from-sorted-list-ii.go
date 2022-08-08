package main

//给定一个已排序的链表的头 head ， 删除原始链表中所有重复数字的节点，只留下不同的数字 。返回 已排序的链表 。
//
// 示例 1： 
//输入：head = [1,2,3,3,4,4,5]
//输出：[1,2,5]
//
// 示例 2： 
//输入：head = [1,1,1,2,3]
//输出：[2,3]

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

	// 使用一个假的头节点
	dummy := &ListNode{0, head}

	node := dummy

	for node.Next != nil && node.Next.Next != nil {
		if node.Next.Val == node.Next.Next.Val {
			v := node.Next.Val
			cur := node.Next
			for cur.Next != nil && cur.Next.Val == v {
				cur = cur.Next
			}
			node.Next = cur.Next
		} else {
			node = node.Next
		}
	}

	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)
