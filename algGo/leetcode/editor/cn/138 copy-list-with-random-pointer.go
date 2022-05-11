package main

import (
	"fmt"
)

// 给你一个长度为 n 的链表，每个节点包含一个额外增加的随机指针 random ，该指针可以指向链表中的任何节点或空节点。
//
// 构造这个链表的 深拷贝。 深拷贝应该正好由 n 个 全新 节点组成，其中每个新节点的值都设为其对应的原节点的值。新节点的 next 指针和 random 
//指针也都应指向复制链表中的新节点，并使原链表和复制链表中的这些指针能够表示相同的链表状态。复制链表中的指针都不应指向原链表中的节点 。 
//
// 例如，如果原链表中有 X 和 Y 两个节点，其中 X.random --> Y 。那么在复制链表中对应的两个节点 x 和 y ，同样有 x.random --> y 。
//
// 返回复制链表的头节点。 
//
// 用一个由 n 个节点组成的链表来表示输入/输出中的链表。每个节点用一个 [val, random_index] 表示： 
//
// val：一个表示 Node.val 的整数。 
// random_index：随机指针指向的节点索引（范围从 0 到 n-1）；如果不指向任何节点，则为 null 。 
//
// 你的代码 只 接受原链表的头节点 head 作为传入参数。 
//
// 示例 1： 
//输入：head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
//输出：[[7,null],[13,0],[11,4],[10,2],[1,0]]
//
// 示例 2： 
//输入：head = [[1,1],[2,1]]
//输出：[[1,1],[2,1]]
//
// 示例 3： 
//输入：head = [[3,null],[3,0],[3,null]]
//输出：[[3,null],[3,0],[3,null]]
//
// 提示： 
// 0 <= n <= 1000
// -10⁴ <= Node.val <= 10⁴ 
// Node.random 为 null 或指向链表中的节点。 
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

//leetcode submit region begin(Prohibit modification and deletion)

// 方法1 map 很简单但是空间复杂度高
//func copyRandomList(head *Node) *Node {
//	if head == nil {
//		return nil
//	}
//
//	copyMap := make(map[*Node]*Node, 0)
//
//	node := head
//	for node != nil {
//		copyMap[node] = &Node{
//			Val:    node.Val,
//			Next:   nil,
//			Random: nil,
//		}
//		node = node.Next
//	}
//
//	node = head
//	for node != nil {
//		copyMap[node].Next = copyMap[node.Next]
//		copyMap[node].Random = copyMap[node.Random]
//		node = node.Next
//	}
//
//	return copyMap[head]
//}

// 方法2 原链表上操作
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	node := head

	// 在每个节点后面插入一个它自己的复制节点
	for node != nil {
		node.Next = &Node{
			Val:    node.Val,
			Next:   node.Next,
			Random: nil,
		}
		node = node.Next.Next
	}

	node = head

	// 复制节点的random指向原random节点的下一个
	for node != nil {
		if node.Random != nil {
			node.Next.Random = node.Random.Next
		}
		node = node.Next.Next
	}

	copyHead := head.Next

	node = head
	copyNode := copyHead

	// 断开原节点
	for copyNode.Next != nil {
		node.Next = node.Next.Next
		copyNode.Next = copyNode.Next.Next
		node = node.Next
		copyNode = copyNode.Next
	}
	node.Next = nil

	return copyHead
}

//leetcode submit region end(Prohibit modification and deletion)

func init() {
	Register(138, func() {
		node1 := &Node{
			Val:    1,
			Next:   nil,
			Random: nil,
		}

		node2 := &Node{
			Val:    2,
			Next:   nil,
			Random: nil,
		}

		node3 := &Node{
			Val:    3,
			Next:   nil,
			Random: nil,
		}

		node1.Next = node2
		node2.Next = node3

		node1.Random = node3
		node3.Random = node2

		printLink := func(head *Node) {
			fmt.Println("Path by next: ")

			node := head
			for node != nil {
				fmt.Print("->")
				fmt.Print(node.Val)
				node = node.Next
			}
			fmt.Println()

			fmt.Println("Path by random: ")

			node = head
			for node != nil {
				fmt.Print("->")
				fmt.Print(node.Val)
				node = node.Random
			}
			fmt.Println()
		}

		fmt.Println("Origin Link: ")
		printLink(node1)

		fmt.Println("------------------------------")

		fmt.Println("Copy Link: ")

		copyHead := copyRandomList(node1)
		printLink(copyHead)

	})
}
