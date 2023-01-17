package main

import (
	"fmt"
)

//给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），每段绳子的长度记为 k[0],k[1]...k[m-1] 。
//请问 k[0]*k[1]*...*k[m-1] 可能的最大乘积是多少？例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18
//。
// 示例 1：
// 输入: 2
//输出: 1
//解释: 2 = 1 + 1, 1 × 1 = 1
//
// 示例 2:
// 输入: 10
//输出: 36
//解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36
// 注意：本题与主站 343 题相同：https://leetcode-cn.com/problems/integer-break/

//leetcode submit region begin(Prohibit modification and deletion)
func cuttingRope(n int) int {
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
func main() {
	fmt.Println(cuttingRope(6))
}
