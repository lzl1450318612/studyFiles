package main

import (
	"fmt"
)

//给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。
// 示例 :
//给定二叉树
//           1
//         / \
//        2   3
//       / \
//      4   5
// 返回 3, 它的长度是路径 [4,2,1,3] 或者 [5,2,1,3]。
//
// 注意：两结点之间的路径长度是以它们之间边的数目表示。

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
func diameterOfBinaryTree(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return 0
	}

	a, b := getTreeMax(root)
	return max(a, b) - 1
}

func getTreeMax(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	leftNotThroughPath, leftThroughPath := getTreeMax(root.Left)
	rightNotThroughPath, rightThroughRoot := getTreeMax(root.Right)

	return max(leftNotThroughPath, rightNotThroughPath) + 1, max(max(leftThroughPath, rightThroughRoot), leftNotThroughPath+rightNotThroughPath+1)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{Val: 3},
	}

	fmt.Println(diameterOfBinaryTree(root))
}
