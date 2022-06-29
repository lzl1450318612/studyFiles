package main

//给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。
//
// 示例 1： 
//输入：root = [3,1,4,null,2], k = 1
//输出：1
//
// 示例 2： 
//输入：root = [5,3,6,2,4,null,null,1], k = 3
//输出：3
// 进阶：如果二叉搜索树经常被修改（插入/删除操作）并且你需要频繁地查找第 k 小的值，你将如何优化算法？

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	res := make([]int, 0, k)
	res = searchKthSmallest(root, res, k)
	return res[k-1]
}

func searchKthSmallest(root *TreeNode, res []int, k int) []int {
	if root == nil {
		return res
	}

	if len(res) == k {
		return res
	}

	res = searchKthSmallest(root.Left, res, k)
	res = append(res, root.Val)
	res = searchKthSmallest(root.Right, res, k)
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
