package main

import (
	"math/big"
)

//一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
// 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
// 问总共有多少条不同的路径？
//
// 示例 1： 
//输入：m = 3, n = 7
//输出：28 
//
// 示例 2： 
//输入：m = 3, n = 2
//输出：3
//解释：
//从左上角开始，总共有 3 条路径可以到达右下角。
//1. 向右 -> 向下 -> 向下
//2. 向下 -> 向下 -> 向右
//3. 向下 -> 向右 -> 向下
// 
// 示例 3：
//输入：m = 7, n = 3
//输出：28
// 
// 示例 4：
//输入：m = 3, n = 3
//输出：6 

//leetcode submit region begin(Prohibit modification and deletion)
//func uniquePaths(m int, n int) int {
//	dp := make([][]int, m)
//	for i := 0; i < len(dp); i++ {
//		dp[i] = make([]int, n)
//	}
//
//	for i := 0; i < len(dp[0]); i++ {
//		dp[0][i] = 1
//	}
//
//	for i := 0; i < len(dp); i++ {
//		dp[i][0] = 1
//	}
//
//	for i := 1; i < len(dp); i++ {
//		for j := 1; j < len(dp[0]); j++ {
//			dp[i][j] = dp[i-1][j] + dp[i][j-1]
//		}
//	}
//
//	return dp[m-1][n-1]
//}

func uniquePaths(m int, n int) int {
	// 因为只能向右或者向下走，所以走到右下角必然走了m+n-2步，因为在m行中，每行都可以选择往下走，所以问题就变成了求C(m+n-2 m-1)
	// C(m+n-2 m-1) = (m+n-2)!/((m-1)!*((m+n-2)-(m-1))!) = (m+n-2)!/((m-1)!*(n-1)!)
	return int(new(big.Int).Binomial(int64(m+n-2), int64(n-1)).Int64())
}

//leetcode submit region end(Prohibit modification and deletion)
