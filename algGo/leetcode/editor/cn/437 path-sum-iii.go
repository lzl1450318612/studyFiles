package main

import (
	"fmt"
)

//给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。
// 路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。
//
// 示例 1：
//输入：root = [10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8
//输出：3
//解释：和等于 8 的路径有 3 条，如图所示。
//
// 示例 2：
//输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
//输出：3

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//leetcode submit region begin(Prohibit modification and deletion)
func pathSum(root *TreeNode, targetSum int) int {

	if root == nil {
		return 0
	}

	res := pathSum(root.Left, targetSum)
	res += pathSum(root.Right, targetSum)
	res += dfsSum(root, targetSum)

	return res
}

func dfsSum(root *TreeNode, target int) int {
	if root == nil {
		return 0
	}

	res := 0
	if root.Val == target {
		res++
	}

	return res + dfsSum(root.Left, target-root.Val) + dfsSum(root.Right, target-root.Val)
}

//func pathSum(root *TreeNode, targetSum int) int {
//
//	if root == nil {
//		return 0
//	}
//
//	count := 0
//	originSum := targetSum
//
//	var getPathSum func(root *TreeNode, target int)
//
//	getPathSum = func(root *TreeNode, target int) {
//		if target == root.Val {
//			count++
//			return
//		}
//
//		if root.Left != nil {
//			getPathSum(root.Left, originSum)
//			getPathSum(root.Left, target-root.Val)
//		}
//
//		if root.Right != nil {
//			getPathSum(root.Right, originSum)
//			getPathSum(root.Right, target-root.Val)
//		}
//
//		return
//	}
//
//	getPathSum(root, originSum)
//
//	return count
//}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	//fmt.Println(pathSum(&TreeNode{
	//	Val: 1,
	//	Left: &TreeNode{
	//		Val: -2,
	//		Left: &TreeNode{
	//			Val:  1,
	//			Left: &TreeNode{Val: -1},
	//			//Right: &TreeNode{Val: -2},
	//		},
	//		Right: &TreeNode{
	//			Val:   173,
	//			Right: &TreeNode{Val: 3},
	//		},
	//	},
	//	Right: &TreeNode{
	//		Val:  -3,
	//		Left: &TreeNode{Val: -2},
	//	},
	//}, 3))

	fmt.Println(pathSum(&TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val: -1,
			Left: &TreeNode{
				Val: 1,
			},
		},
		Right: &TreeNode{
			Val: 2,
		},
	}, 0))
}
