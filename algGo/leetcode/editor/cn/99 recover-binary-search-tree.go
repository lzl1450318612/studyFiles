package main

//给你二叉搜索树的根节点 root ，该树中的 恰好 两个节点的值被错误地交换。请在不改变其结构的情况下，恢复这棵树 。
//
// 示例 1：
//输入：root = [1,3,null,null,2]
//输出：[3,1,null,null,2]
//解释：3 不能是 1 的左孩子，因为 3 > 1 。交换 1 和 3 使二叉搜索树有效。
//
// 示例 2：
//输入：root = [3,1,4,null,null,2]
//输出：[2,1,4,null,null,3]
//解释：2 不能在 3 的右子树中，因为 2 < 3 。交换 2 和 3 使二叉搜索树有效。
//
// 进阶：使用 O(n) 空间复杂度的解法很容易实现。你能想出一个只使用 O(1) 空间的解决方案吗？

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//leetcode submit region begin(Prohibit modification and deletion)
//var t1, t2, preNode *TreeNode
//
//func recoverTree(root *TreeNode) {
//	searchIncorrectNode(root)
//	temp := t1.Val
//	t1.Val = t2.Val
//	t2.Val = temp
//}
//
//func searchIncorrectNode(root *TreeNode) {
//	if root == nil {
//		return
//	}
//
//	searchIncorrectNode(root.Left)
//	if preNode != nil && root.Val < preNode.Val {
//		if t1 == nil {
//			t1 = preNode
//		}
//		t2 = root
//	}
//	preNode = root
//	searchIncorrectNode(root.Right)
//
//}
func recoverTree(root *TreeNode) {
	stack := []*TreeNode{}
	var x, y, pred *TreeNode
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pred != nil && root.Val < pred.Val {
			y = root
			if x == nil {
				x = pred
			} else {
				break
			}
		}
		pred = root
		root = root.Right
	}
	x.Val, y.Val = y.Val, x.Val
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	recoverTree(&TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 2,
			},
		},
	})
}
