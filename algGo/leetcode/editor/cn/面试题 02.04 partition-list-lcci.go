package main

import (
	"fmt"
)

// 给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
// 你不需要 保留 每个分区中各节点的初始相对位置。
//
// 示例 1： 
//输入：head = [1,4,3,2,5,2], x = 3
//输出：[1,2,2,4,3,5]
//
// 示例 2： 
//输入：head = [2,1], x = 2
//输出：[1,2]
//
// 提示： 
// 链表中节点的数目在范围 [0, 200] 内
// -100 <= Node.val <= 100 
// -200 <= x <= 200 

//leetcode submit region begin(Prohibit modification and deletion)
type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	cur, prev := head, head
	for cur != nil {
		if cur.Val < x {
			tmp := prev.Val
			prev.Val = cur.Val
			cur.Val = tmp
			prev = prev.Next
		}
		cur = cur.Next
	}
	return head
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	PrintLink(partition(InitLink([]int{2, 1}), 2))
}

// PrintLink 打印链表
func PrintLink(head *ListNode) {
	if head == nil {
		return
	}
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

// InitLink 使用arr初始化链表
func InitLink(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	var head, node *ListNode
	for _, v := range arr {
		if head == nil {
			head = &ListNode{
				Val:  v,
				Next: nil,
			}
			node = head
			continue
		}
		node.Next = &ListNode{
			Val:  v,
			Next: nil,
		}
		node = node.Next
	}
	return head
}
