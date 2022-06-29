package main

//给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。
//
// 示例 1： 
//输入：root = [3,9,20,null,null,15,7]
//输出：[[3],[9,20],[15,7]]
//
// 示例 2： 
//输入：root = [1]
//输出：[[1]]
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
func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)

	order := make([]*TreeNode, 0)
	if root == nil {
		return res
	}
	order = append(order, root)

	for len(order) != 0 {
		layerCount := len(order)
		items := make([]int, 0, layerCount)
		for ; layerCount > 0; layerCount-- {
			node := order[0]
			order = order[1:]
			items = append(items, node.Val)
			if node.Left != nil {
				order = append(order, node.Left)
			}
			if node.Right != nil {
				order = append(order, node.Right)
			}
		}
		res = append(res, items)
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
