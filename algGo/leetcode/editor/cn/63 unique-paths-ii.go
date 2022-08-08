package main

import (
	"fmt"
)

//一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
// 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish”）。
// 现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
// 网格中的障碍物和空位置分别用 1 和 0 来表示。
//
// 示例 1：
//输入：obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]
//输出：2
//解释：3x3 网格的正中间有一个障碍物。
//从左上角到右下角一共有 2 条不同的路径：
//1. 向右 -> 向右 -> 向下 -> 向下
//2. 向下 -> 向下 -> 向右 -> 向右
//
// 示例 2： 
//输入：obstacleGrid = [[0,1],[0,0]]
//输出：1

//leetcode submit region begin(Prohibit modification and deletion)
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	height := len(obstacleGrid)
	width := len(obstacleGrid[0])

	dp := make([][]int, height)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, width)
	}

	if obstacleGrid[0][0] != 1 {
		dp[0][0] = 1
	}

	for i := 1; i < len(dp[0]); i++ {
		if obstacleGrid[0][i] == 1 {
			continue
		}
		dp[0][i] = dp[0][i-1]
	}

	for i := 1; i < len(dp); i++ {
		if obstacleGrid[i][0] == 1 {
			continue
		}
		dp[i][0] = dp[i-1][0]
	}

	for i := 1; i < height; i++ {
		for j := 1; j < width; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
				continue
			}

			up := obstacleGrid[i-1][j]
			left := obstacleGrid[i][j-1]

			if up == 1 && left == 1 {
				dp[i][j] = 0
			} else if up == 1 {
				dp[i][j] = dp[i][j-1]
			} else if left == 1 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}

		}
	}

	return dp[height-1][width-1]
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(uniquePathsWithObstacles([][]int{
		{1, 0},
	}))
}
