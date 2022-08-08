package main

//给你两棵二叉树的根节点 p 和 q ，编写一个函数来检验这两棵树是否相同。
// 如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。
//
// 示例 1：
//输入：p = [1,2,3], q = [1,2,3]
//输出：true
//
// 示例 2： 
//输入：p = [1,2], q = [1,null,2]
//输出：false
//
// 示例 3： 
//输入：p = [1,2,1], q = [1,1,2]
//输出：false

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}

		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
