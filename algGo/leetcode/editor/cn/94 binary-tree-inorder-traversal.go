package main

//给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。
//
// 示例 1： 
//输入：root = [1,null,2,3]
//输出：[1,3,2]
//
// 示例 2： 
//输入：root = []
//输出：[]
//
// 示例 3： 
//输入：root = [1]
//输出：[1]

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	node := root
	stack := make([]*TreeNode, 0)
	for node != nil || len(stack) > 0 {
		if node != nil {
			stack = append(stack, node)
			node = node.Left
		} else {
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, a.Val)
			node = a.Right
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
