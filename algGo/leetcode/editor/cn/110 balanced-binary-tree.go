package main

//给定一个二叉树，判断它是否是高度平衡的二叉树。
// 本题中，一棵高度平衡二叉树定义为：
// 一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1 。
//
// 示例 1： 
//输入：root = [3,9,20,null,null,15,7]
//输出：true
//
// 示例 2： 
//输入：root = [1,2,2,3,3,null,null,4,4]
//输出：false
//
// 示例 3： 
//输入：root = []
//输出：true

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	_, balance := process(root)
	return balance
}

func process(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}

	h1, isBalanced1 := process(root.Left)
	h2, isBalanced2 := process(root.Right)

	height := treeMax(h1, h2) + 1

	return height, isBalanced1 && isBalanced2 && !(h1 - h2 > 1 || h2 - h1 > 1)
}

func treeMax(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
