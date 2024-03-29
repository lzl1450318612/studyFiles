package main

//小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为 root 。
// 除了 root 之外，每栋房子有且只有一个“父“房子与之相连。一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。 如果 两个直接
//相连的房子在同一天晚上被打劫 ，房屋将自动报警。
// 给定二叉树的 root 。返回 在不触动警报的情况下 ，小偷能够盗取的最高金额 。
//
// 示例 1:
//输入: root = [3,2,3,null,3,null,1]
//输出: 7
//解释: 小偷一晚能够盗取的最高金额 3 + 3 + 1 = 7
//
// 示例 2:
//输入: root = [3,4,5,1,3,null,1]
//输出: 9
//解释: 小偷一晚能够盗取的最高金额 4 + 5 = 9

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob(root *TreeNode) int {
	stealCur, notStealCur := process(root)
	return max(stealCur, notStealCur)
}

func process(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}

	stealLeft, notStealLeft := process(root.Left)
	stealRight, notStealRight := process(root.Right)

	return root.Val + notStealLeft + notStealRight, arrMax([]int{stealLeft + stealRight, notStealRight + notStealLeft, stealLeft + notStealRight, stealRight + notStealLeft})

}

func arrMax(arr []int) int {
	m := 0
	for _, v := range arr {
		if v > m {
			m = v
		}
	}
	return m
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
