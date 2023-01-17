package main

//给定一个正整数 n ，将其拆分为 k 个 正整数 的和（ k >= 2 ），并使这些整数的乘积最大化。
// 返回 你可以获得的最大乘积 。
//
// 示例 1:
//输入: n = 2
//输出: 1
//解释: 2 = 1 + 1, 1 × 1 = 1。
//
// 示例 2:
//输入: n = 10
//输出: 36
//解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36。

//leetcode submit region begin(Prohibit modification and deletion)
func integerBreak(n int) int {
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	if n == 3 {
		return 2
	}

	dp := make([]int, n+1)

	dp[0] = 0
	dp[1] = 1
	dp[2] = 2
	dp[3] = 3

	for i := 4; i <= n; i++ {
		maxChengji := 1
		for j := 1; j < i; j++ {
			maxChengji = max(maxChengji, dp[i-j]*dp[j])
		}
		dp[i] = maxChengji
	}

	return dp[n]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
