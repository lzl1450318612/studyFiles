package main

//给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
// 示例 1：
//输入：head = [1,2,3,4]
//输出：[2,1,4,3]
//
// 示例 2： 
//输入：head = []
//输出：[]
// 
// 示例 3：
//输入：head = [1]
//输出：[1]

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	var last *ListNode
	node := head

	for node != nil && node.Next != nil {
		if node == head {
			temp := node.Next
			node.Next = temp.Next
			temp.Next = node
			last = node
			head = temp
		} else {
			temp := node.Next
			node.Next = temp.Next
			temp.Next = node
			last.Next = temp
			last = node
		}

		node = node.Next
	}
	return head
}

//leetcode submit region end(Prohibit modification and deletion)
