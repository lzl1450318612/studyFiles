package main

import (
	"fmt"
	"math"
)

//给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
// 你应当 保留 两个分区中每个节点的初始相对位置。
//
// 示例 1： 
//输入：head = [1,4,3,2,5,2], x = 3
//输出：[1,2,2,4,3,5]
// 
// 示例 2：
//输入：head = [2,1], x = 2
//输出：[1,2]

type ListNode struct {
	Val  int
	Next *ListNode
}

//leetcode submit region begin(Prohibit modification and deletion)
func partition(head *ListNode, x int) *ListNode {
	fakeHead := &ListNode{
		Val:  math.MinInt32,
		Next: head,
	}

	slow, fast := fakeHead, fakeHead

	for slow.Next != nil && slow.Next.Val < x {
		slow = slow.Next
	}

	fast = slow

	for fast.Next != nil {
		for fast.Next != nil && fast.Next.Val >= x {
			fast = fast.Next
		}

		if fast.Next == nil {
			return fakeHead.Next
		}

		temp := fast.Next
		fast.Next = temp.Next
		temp.Next = slow.Next
		slow.Next = temp
		slow = slow.Next
	}

	return fakeHead.Next
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val:  2,
							Next: nil,
						},
					},
				},
			},
		},
	}

	fmt.Println(partition(head, 2))
}
