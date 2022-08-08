package main

import (
	"fmt"
)

//给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
// 说明：每次只能向下或者向右移动一步。
//
// 示例 1： 
//输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
//输出：7
//解释：因为路径 1→3→1→1→1 的总和最小。
//
// 示例 2： 
//输入：grid = [[1,2,3],[4,5,6]]
//输出：12

//leetcode submit region begin(Prohibit modification and deletion)
func minPathSum(grid [][]int) int {
	height := len(grid)
	width := len(grid[0])

	dp := make([][]int, height)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, width)
	}

	dp[0][0] = grid[0][0]

	for i := 1; i < height; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	for i := 1; i < width; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}

	for i := 1; i < height; i++ {
		for j := 1; j < width; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[height-1][width-1]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(minPathSum([][]int{
		{1, 3, 1}, {1, 5, 1}, {4, 2, 1},
	}))
}
