package main

//请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。
// 例如:
//给定二叉树: [3,9,20,null,null,15,7],
//     3
//   / \
//  9  20
//    /  \
//   15   7
// 返回其层次遍历结果：
// [
//  [3],
//  [20,9],
//  [15,7]
//]

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

	if root == nil {
		return res
	}

	queue := []*TreeNode{root}
	level := 0

	for len(queue) != 0 {
		levelRes := make([]int, 0)
		temp := make([]*TreeNode, 0)
		for _, treeNode := range queue {
			levelRes = append(levelRes, treeNode.Val)
			if treeNode.Left != nil {
				temp = append(temp, treeNode.Left)
			}
			if treeNode.Right != nil {
				temp = append(temp, treeNode.Right)
			}
		}
		queue = temp
		if level%2 != 0 {
			for i, j := 0, len(levelRes)-1; i < j; {
				levelRes[i], levelRes[j] = levelRes[j], levelRes[i]
				i++
				j--
			}
		}
		res = append(res, levelRes)
		level++
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
