package main

//给你二叉树的根结点 root ，请你将它展开为一个单链表：
// 展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
// 展开后的单链表应该与二叉树 先序遍历 顺序相同。
//
// 示例 1：
//输入：root = [1,2,5,3,4,null,6]
//输出：[1,null,2,null,3,null,4,null,5,null,6]
//
// 示例 2：
//输入：root = []
//输出：[]
//
// 示例 3：
//输入：root = [0]
//输出：[0]
//
// 进阶：你可以使用原地算法（O(1) 额外空间）展开这棵树吗？

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	root, _ = processFlattenTree(root)
}

func processFlattenTree(root *TreeNode) (*TreeNode, *TreeNode) {
	if root == nil {
		return nil, nil
	}
	if root.Left == nil && root.Right == nil {
		return root, root
	}
	leftHead, leftTail := processFlattenTree(root.Left)
	rightHead, rightTail := processFlattenTree(root.Right)

	root.Left = nil

	if leftHead == nil {
		return root, rightTail
	}

	root.Right = leftHead
	leftTail.Right = rightHead

	if rightHead == nil {
		return root, leftTail
	}

	return root, rightTail
}

//leetcode submit region end(Prohibit modification and deletion)
