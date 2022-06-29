package main

import (
	"fmt"
	"math"
)

//给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
// 有效 二叉搜索树定义如下：
// 节点的左子树只包含 小于 当前节点的数。
// 节点的右子树只包含 大于 当前节点的数。
// 所有左子树和右子树自身必须也是二叉搜索树。
//
// 示例 1：
//输入：root = [2,1,3]
//输出：true
//
// 示例 2：
//输入：root = [5,1,4,null,null,3,6]
//输出：false
//解释：根节点的值是 5 ，但是右子节点的值是 4 。
//
// 提示：
// 树中节点数目范围在[1, 10⁴] 内
// -2³¹ <= Node.val <= 2³¹ - 1

//leetcode submit region begin(Prohibit modification and deletion)

func isValidBST(root *TreeNode) bool {
	return judge(root, math.MinInt64, math.MaxInt64)
}

func judge(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}

	if root.Val <= min || root.Val >= max {
		return false
	}

	return judge(root.Left, min, root.Val) && judge(root.Right, root.Val, max)
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	node1 := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}

	node3 := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}

	node2 := &TreeNode{
		Val:   2,
		Left:  node1,
		Right: node3,
	}

	fmt.Println(isValidBST(node2))
}
