package main

//给定一个整数数组 prices，其中第 prices[i] 表示第 i 天的股票价格 。
// 设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
// 卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
//
// 示例 1:
//输入: prices = [1,2,3,0,2]
//输出: 3
//解释: 对应的交易状态为: [买入, 卖出, 冷冻期, 买入, 卖出]
//
// 示例 2:
//输入: prices = [1]
//输出: 0

//leetcode submit region begin(Prohibit modification and deletion)
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	dp := make([][]int, len(prices))

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 3)
	}

	dp[0][1] = -prices[0]
	dp[0][2] = -prices[0]

	for i := 1; i < len(dp); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][2])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
		dp[i][2] = dp[i-1][1] + prices[i]
	}

	return max(dp[len(prices)-1][0], dp[len(prices)-1][2])
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
