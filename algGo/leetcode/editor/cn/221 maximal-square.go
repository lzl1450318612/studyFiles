package main

//在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。
//
// 示例 1：
//输入：matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"]
//,["1","0","0","1","0"]]
//输出：4
//
// 示例 2：
//输入：matrix = [["0","1"],["1","0"]]
//输出：1
//
// 示例 3：
//输入：matrix = [["0"]]
//输出：0

//leetcode submit region begin(Prohibit modification and deletion)
func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	dp := make([][]int, len(matrix)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(matrix[0])+1)
	}

	maxSide := 0

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				dp[i+1][j+1] = min(min(dp[i][j+1], dp[i+1][j]), dp[i][j]) + 1
				maxSide = max(maxSide, dp[i+1][j+1])
			}
		}
	}

	return maxSide * maxSide
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
