package main

import (
	"studyFiles/algGo/utils"
)

//给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
//
// 请你将两个数相加，并以相同形式返回一个表示和的链表。 
//
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。 
//
//输入：l1 = [2,4,3], l2 = [5,6,4]
//输出：[7,0,8]
//解释：342 + 465 = 807.
//
//输入：l1 = [0], l2 = [0]
//输出：[0]
//
//输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
//输出：[8,9,9,9,0,0,0,1]
//
// 每个链表中的节点数在范围 [1, 100] 内 
// 0 <= Node.val <= 9 
// 题目数据保证列表表示的数字不含前导零

//leetcode submit region begin(Prohibit modification and deletion)

func addTwoNumbers(l1 *utils.ListNode, l2 *utils.ListNode) *utils.ListNode {
	switch {
	case l1 == nil && l2 == nil:
		return nil
	case l1 == nil:
		return l2
	case l2 == nil:
		return l1
	}

	var res *utils.ListNode = nil
	var current *utils.ListNode = nil

	overflow := false
	for l1 != nil || l2 != nil || overflow {
		a := 0

		if overflow {
			a++
		}
		overflow = false

		if l1 != nil {
			a += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			a += l2.Val
			l2 = l2.Next
		}

		if a >= 10 {
			overflow = true
			a -= 10
		}

		if current == nil {
			current = &utils.ListNode{
				Val:  a,
				Next: nil,
			}
			res = current
		} else {
			current.Next = &utils.ListNode{
				Val:  a,
				Next: nil,
			}
			current = current.Next
		}
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)

//func main() {
//	l1 := &ListNode{
//		Val: 2,
//		Next: &ListNode{
//			Val: 4,
//			Next: &ListNode{
//				Val:  3,
//				Next: nil,
//			},
//		},
//	}
//
//	l2 := &ListNode{
//		Val: 5,
//		Next: &ListNode{
//			Val: 6,
//			Next: &ListNode{
//				Val:  4,
//				Next: nil,
//			},
//		},
//	}
//
//	res := addTwoNumbers(l1, l2)
//
//	for res != nil {
//		fmt.Println(res.Val)
//		res = res.Next
//	}
//}
