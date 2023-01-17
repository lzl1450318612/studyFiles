package main

// 给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null 。
// 图示两个链表在节点 c1 开始相交：
// 题目数据 保证 整个链式结构中不存在环。
// 注意，函数返回结果后，链表必须 保持其原始结构 。
//
// 示例 1：
//输入：intersectVal = 8, listA = [4,1,8,4,5], listB = [5,6,1,8,4,5], skipA = 2,
//skipB = 3
//输出：Intersected at '8'
//解释：相交节点的值为 8 （注意，如果两个链表相交则不能为 0）。
//从各自的表头开始算起，链表 A 为 [4,1,8,4,5]，链表 B 为 [5,6,1,8,4,5]。
//在 A 中，相交节点前有 2 个节点；在 B 中，相交节点前有 3 个节点。
//
// 示例 2：
//输入：intersectVal = 2, listA = [1,9,1,2,4], listB = [3,2,4], skipA = 3, skipB =
//1
//输出：Intersected at '2'
//解释：相交节点的值为 2 （注意，如果两个链表相交则不能为 0）。
//从各自的表头开始算起，链表 A 为 [1,9,1,2,4]，链表 B 为 [3,2,4]。
//在 A 中，相交节点前有 3 个节点；在 B 中，相交节点前有 1 个节点。
//
// 示例 3：
//输入：intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2
//输出：null
//解释：从各自的表头开始算起，链表 A 为 [2,6,4]，链表 B 为 [1,5]。
//由于这两个链表不相交，所以 intersectVal 必须为 0，而 skipA 和 skipB 可以是任意值。
//这两个链表不相交，因此返回 null 。

//leetcode submit region begin(Prohibit modification and deletion)
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	node1, node2 := headA, headB

	len1, len2 := 0, 0

	for node1.Next != nil {
		node1 = node1.Next
		len1++
	}

	for node2.Next != nil {
		node2 = node2.Next
		len2++
	}

	if node1 != node2 {
		return nil
	}

	node1, node2 = headA, headB
	if len1 > len2 {
		for i := 0; i < len1-len2; i++ {
			node1 = node1.Next
		}
	} else if len1 < len2 {
		for i := 0; i < len2-len1; i++ {
			node2 = node2.Next
		}
	}

	for node1 != node2 {
		node1 = node1.Next
		node2 = node2.Next
	}
	return node1
}

//leetcode submit region end(Prohibit modification and deletion)
