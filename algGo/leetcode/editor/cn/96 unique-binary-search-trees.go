package main

//给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。
//
// 示例 1： 
//输入：n = 3
//输出：5
//
// 示例 2： 
//输入：n = 1
//输出：1

//leetcode submit region begin(Prohibit modification and deletion)
func numTrees(n int) int {
	dp := make([]int, n+1)

	dp[0] = 1
	dp[1] = 1

	// i代表从1-i中，不同二叉搜索树的种数
	for i := 2; i <= n; i++ {
		// j代表以1-i中每个节点作为根节点时二叉树分别的种数
		for j := 1; j <= i; j++ {
			// 从1-i中，不同二叉搜索树的种数 = 以每个数作为根节点时，左子树的种数 * 右子树的种数；
			// 因为不关心二叉树的取值只数个数，所以例如：1，2，3，4，5 选取3为根节点时，4，5形成的子树个数和1，2形成的子树个数相等，有对称性，所以可以直接使用dp[i-j]作为右边子树的个数
			dp[i] += dp[j-1] * dp[i-j]
		}
	}

	return dp[n]
}

//leetcode submit region end(Prohibit modification and deletion)
