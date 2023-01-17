package main

import (
	"fmt"
	"math"
)

//给你一个整数 n ，返回 和为 n 的完全平方数的最少数量 。
// 完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。
//
// 示例 1：
//输入：n = 12
//输出：3
//解释：12 = 4 + 4 + 4
//
// 示例 2：
//输入：n = 13
//输出：2
//解释：13 = 4 + 9

//leetcode submit region begin(Prohibit modification and deletion)
func numSquares(n int) int {
	if n == 0 {
		return 0
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1

	for i := 2; i <= n; i++ {
		minCount := math.MaxInt32
		for j := 1; j < i; j++ {
			if j*j == i {
				minCount = 1
				break
			}
			if j*j > i {
				break
			}
			minCount = min(minCount, dp[i-j*j]+1)
		}
		dp[i] = minCount
	}

	return dp[n]
}

func findMaxPower(n int, dp []int) int {
	if n < 0 {
		return -1
	}
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if dp[n] != -1 {
		return dp[n]
	}
	minCount := math.MaxInt32
	for i := 1; i < n; i++ {
		a := i * i
		temp := findMaxPower(n-a, dp)
		if temp == -1 {
			continue
		}
		minCount = min(temp, minCount)
	}
	dp[n] = minCount + 1
	return minCount + 1

}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(numSquares(52))
}
