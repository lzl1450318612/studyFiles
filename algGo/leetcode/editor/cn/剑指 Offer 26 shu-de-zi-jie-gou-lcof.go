package main

//输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)
// B是A的子结构， 即 A中有出现和B相同的结构和节点值。
//
// 例如:
//给定的树 A:
// 3
// / \
// 4 5
// / \
// 1 2
//给定的树 B：
// 4
// /
// 1
//返回 true，因为 B 与 A 的一个子树拥有相同的结构和节点值。
//
// 示例 1：
// 输入：A = [1,2,3], B = [3,1]
//输出：false
//
// 示例 2：
// 输入：A = [3,4,5,1,2], B = [4,1]
//输出：true

//leetcode submit region begin(Prohibit modification and deletion)
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}

	return judgeIsSubTree(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func judgeIsSubTree(a *TreeNode, b *TreeNode) bool {
	if b == nil {
		return true
	}

	if a == nil {
		return false
	}

	return a.Val == b.Val && judgeIsSubTree(a.Left, b.Left) && judgeIsSubTree(a.Right, b.Right)
}

//leetcode submit region end(Prohibit modification and deletion)
