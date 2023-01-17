package main

//请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。
// 例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
// 1
// / \
// 2 2
// / \ / \
//3 4 4 3
//但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
// 1
// / \
// 2 2
// \ \
// 3 3
//
// 示例 1：
// 输入：root = [1,2,2,3,4,4,3]
//输出：true
//
// 示例 2：
// 输入：root = [1,2,2,null,3,null,3]
//输出：false
// 注意：本题与主站 101 题相同：https://leetcode-cn.com/problems/symmetric-tree/

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

	return dfsJudge(root.Left, root.Right)
}

func dfsJudge(a *TreeNode, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}
	return a.Val == b.Val && dfsJudge(a.Left, b.Right) && dfsJudge(a.Right, b.Left)
}

//leetcode submit region end(Prohibit modification and deletion)
