package main

//给你一棵二叉树的根节点 root ，翻转这棵二叉树，并返回其根节点。
//
// 示例 1：
//输入：root = [4,2,7,1,3,6,9]
//输出：[4,7,2,9,6,3,1]
//
// 示例 2：
//输入：root = [2,1,3]
//输出：[2,3,1]
//
// 示例 3：
//输入：root = []
//输出：[]

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}
	return invertProcess(root)
}

func invertProcess(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left != nil || root.Right != nil {
		root.Left, root.Right = root.Right, root.Left
	}
	root.Left = invertProcess(root.Left)
	root.Right = invertTree(root.Right)
	return root
}

//leetcode submit region end(Prohibit modification and deletion)
