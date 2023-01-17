package main

//请完成一个函数，输入一个二叉树，该函数输出它的镜像。
// 例如输入：
// 4
// / \
// 2 7
// / \ / \
//1 3 6 9
//镜像输出：
// 4
// / \
// 7 2
// / \ / \
//9 6 3 1
//
// 示例 1：
// 输入：root = [4,2,7,1,3,6,9]
//输出：[4,7,2,9,6,3,1]
//
// 注意：本题与主站 226 题相同：https://leetcode-cn.com/problems/invert-binary-tree/

//leetcode submit region begin(Prohibit modification and deletion)
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = mirrorTree(root.Right), mirrorTree(root.Left)
	return root
}

//leetcode submit region end(Prohibit modification and deletion)
