package main

//给你一个二叉树的根节点 root ， 检查它是否轴对称。
//
// 示例 1：
//输入：root = [1,2,2,3,4,4,3]
//输出：true
//
// 示例 2：
//输入：root = [1,2,2,null,3,null,3]
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
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return searchSymmetric(root.Left, root.Right)
}

func searchSymmetric(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}

	if root1 == nil || root2 == nil || root1.Val != root2.Val {
		return false
	}
	return searchSymmetric(root1.Left, root2.Right) && searchSymmetric(root1.Right, root2.Left)
}

//leetcode submit region end(Prohibit modification and deletion)
