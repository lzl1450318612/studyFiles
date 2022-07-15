package main

import (
	"container/heap"
)

//给你一个链表数组，每个链表都已经按升序排列。
// 请你将所有链表合并到一个升序链表中，返回合并后的链表。
//
// 示例 1： 
// 输入：lists = [[1,4,5],[1,3,4],[2,6]]
//输出：[1,1,2,3,4,4,5,6]
//解释：链表数组如下：
//[
//  1->4->5,
//  1->3->4,
//  2->6
//]
//将它们合并到一个有序链表中得到。
//1->1->2->3->4->4->5->6
//
// 示例 2： 
// 输入：lists = []
//输出：[]
//
// 示例 3： 
// 输入：lists = [[]]
//输出：[]

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	var res *ListNode
	var tail *ListNode

	h := &listHeap{}

	for _, head := range lists {
		if head != nil {
			heap.Push(h, head)
		}
	}

	for h.Len() > 0 {
		node := heap.Pop(h).(*ListNode)
		newHead := removeHead(node)
		if newHead != nil {
			heap.Push(h, newHead)
		}
		res, tail = addNode(res, node, tail)
	}

	return res
}

type listHeap []*ListNode

func (l listHeap) Len() int {
	return len(l)
}

func (l listHeap) Less(i, j int) bool {
	return l[i].Val < l[j].Val
}

func (l listHeap) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l *listHeap) Push(x interface{}) {
	*l = append(*l, x.(*ListNode))
}

func (l *listHeap) Pop() interface{} {
	old := *l
	n := len(old)
	x := old[n-1]
	*l = old[0 : n-1]
	return x
}

func removeHead(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	node := head.Next
	head.Next = nil
	return node
}

func addNode(head, node, tail *ListNode) (*ListNode, *ListNode) {
	if head == nil {
		return node, node
	}
	tail.Next = node
	return head, node
}

//leetcode submit region end(Prohibit modification and deletion)
