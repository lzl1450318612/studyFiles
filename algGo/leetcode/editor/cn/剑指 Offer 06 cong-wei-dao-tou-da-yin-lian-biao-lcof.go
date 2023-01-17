package main

//输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。
//
// 示例 1：
// 输入：head = [1,3,2]
//输出：[2,3,1]

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) []int {
	if head == nil {
		return []int{}
	}

	if head.Next == nil {
		return []int{head.Val}
	}
	res := reversePrint(head.Next)
	res = append(res, head.Val)
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
