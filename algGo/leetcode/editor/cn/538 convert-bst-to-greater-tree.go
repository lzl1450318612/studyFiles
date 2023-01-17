package main

import (
	"fmt"
)

//给出二叉 搜索 树的根节点，该树的节点值各不相同，请你将其转换为累加树（Greater Sum Tree），使每个节点 node 的新值等于原树中大于或等于
// node.val 的值之和。
// 提醒一下，二叉搜索树满足下列约束条件：
// 节点的左子树仅包含键 小于 节点键的节点。
// 节点的右子树仅包含键 大于 节点键的节点。
// 左右子树也必须是二叉搜索树。
// 注意：本题和 1038: https://leetcode-cn.com/problems/binary-search-tree-to-greater-
//sum-tree/ 相同
//
// 示例 1：
// 输入：[4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]
//输出：[30,36,21,36,35,26,15,null,null,null,33,null,null,null,8]
//
// 示例 2：
// 输入：root = [0,null,1]
//输出：[1,null,1]
//
// 示例 3：
// 输入：root = [1,0,2]
//输出：[3,3,2]
//
// 示例 4：
// 输入：root = [3,2,4,1]
//输出：[7,9,4,10]

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	var pre *TreeNode
	var f func(*TreeNode)

	f = func(root *TreeNode) {
		if root.Left == nil && root.Right == nil {
			if pre != nil {
				root.Val += pre.Val
			}
			pre = root
			return
		}

		if root.Right != nil {
			f(root.Right)
		}

		if pre != nil {
			root.Val += pre.Val
		}
		pre = root

		if root.Left != nil {
			f(root.Left)
		}
	}

	f(root)
	return root
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(convertBST(&TreeNode{
		Val: 0,
		Right: &TreeNode{
			Val: 1,
		},
	}))
}
