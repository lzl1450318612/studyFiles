package main

//给你一个整数 n ，请你生成并返回所有由 n 个节点组成且节点值从 1 到 n 互不相同的不同 二叉搜索树 。可以按 任意顺序 返回答案。
//
// 示例 1： 
//输入：n = 3
//输出：[[1,null,2,null,3],[1,null,3,2],[2,1,3],[3,1,null,null,2],[3,2,null,1]]
//
// 示例 2： 
//输入：n = 1
//输出：[[1]]

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
	return genTrees(1, n)

}

func genTrees(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}

	res := make([]*TreeNode, 0)
	for i := start; i <= end; i++ {
		leftSubTrees := genTrees(start, i-1)
		rightSubTrees := genTrees(i+1, end)

		for _, leftSubTree := range leftSubTrees {
			for _, rightSubTree := range rightSubTrees {
				curNode := &TreeNode{Val: i}
				curNode.Left = leftSubTree
				curNode.Right = rightSubTree
				res = append(res, curNode)
			}
		}
	}
	return res
}
//leetcode submit region end(Prohibit modification and deletion)
