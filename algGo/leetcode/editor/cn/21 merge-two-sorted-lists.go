package main

//将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//
// 示例 1： 
//输入：l1 = [1,2,4], l2 = [1,3,4]
//输出：[1,1,2,3,4,4]
//
// 示例 2： 
//输入：l1 = [], l2 = []
//输出：[]
//
// 示例 3： 
//输入：l1 = [], l2 = [0]
//输出：[0]
//
// 提示： 
// 两个链表的节点数目范围是 [0, 50]
// -100 <= Node.val <= 100 
// l1 和 l2 均按 非递减顺序 排列 

//leetcode submit region begin(Prohibit modification and deletion)
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	var ret *ListNode
	var tail *ListNode

	n1, n2 := list1, list2
	for n1 != nil && n2 != nil {
		if n1.Val <= n2.Val {
			ret, tail = addNodeToList(ret, tail, n1)
			n1 = n1.Next
		} else {
			ret, tail = addNodeToList(ret, tail, n2)
			n2 = n2.Next
		}
	}

	for n1 != nil {
		ret, tail = addNodeToList(ret, tail, n1)
		n1 = n1.Next
	}

	for n2 != nil {
		ret, tail = addNodeToList(ret, tail, n2)
		n2 = n2.Next
	}

	return ret
}

func addNodeToList(head, tail, node *ListNode) (*ListNode, *ListNode) {
	if head == nil {
		return node, node
	}

	tail.Next = node
	tail = tail.Next
	return head, tail
}

//leetcode submit region end(Prohibit modification and deletion)
