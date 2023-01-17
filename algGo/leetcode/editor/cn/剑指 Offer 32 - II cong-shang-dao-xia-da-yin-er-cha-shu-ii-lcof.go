package main

//从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。
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
//  [9,20],
//  [15,7]
//]
//
// 注意：本题与主站 102 题相同：https://leetcode-cn.com/problems/binary-tree-level-order-

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
	//
	//res := make([][]int, 0)
	//
	//if root == nil {
	//	return res
	//}
	//
	//return dfsTree(root, 0, res)
	res := make([][]int, 0)

	if root == nil {
		return res
	}

	queue := []*TreeNode{root}

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
		res = append(res, levelRes)
	}
	return res
}

func dfsTree(root *TreeNode, level int, res [][]int) [][]int {

	if len(res) <= level {
		res = append(res, []int{})
	}

	res[level] = append(res[level], root.Val)

	if root.Left != nil {
		res = dfsTree(root.Left, level+1, res)
	}

	if root.Right != nil {
		res = dfsTree(root.Right, level+1, res)
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
