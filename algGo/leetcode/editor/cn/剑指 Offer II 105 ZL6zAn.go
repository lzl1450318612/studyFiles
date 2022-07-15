package main

//给定一个由 0 和 1 组成的非空二维数组 grid ，用来表示海洋岛屿地图。
// 一个 岛屿 是由一些相邻的 1 (代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在水平或者竖直方向上相邻。你可以假设 grid 的四个边缘都被
//0（代表水）包围着。 
// 找到给定的二维数组中最大的岛屿面积。如果没有岛屿，则返回面积为 0 。
//
// 示例 1: 
//输入: grid = [[0,0,1,0,0,0,0,1,0,0,0,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,1,1,0,1
//,0,0,0,0,0,0,0,0],[0,1,0,0,1,1,0,0,1,0,1,0,0],[0,1,0,0,1,1,0,0,1,1,1,0,0],[0,0,0
//,0,0,0,0,0,0,0,1,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,0,0,0,0,0,0,1,1,0,0,0,0]]
//输出: 6
//解释: 对于上面这个给定矩阵应返回 6。注意答案不应该是 11 ，因为岛屿只能包含水平或垂直的四个方向的 1 。 
//
// 示例 2: 
//输入: grid = [[0,0,0,0,0,0,0,0]]
//输出: 0 

//leetcode submit region begin(Prohibit modification and deletion)
func maxAreaOfIsland(grid [][]int) int {
	height := len(grid) - 1
	width := len(grid[0]) - 1

	maxCount := 0
	for i := 0; i <= height; i++ {
		for j := 0; j <= width; j++ {
			if grid[i][j] == 1 {
				c := infect(grid, width, height, i, j, 0)
				if c > maxCount {
					maxCount = c
				}
			}
		}
	}
	return maxCount
}

func infect(grid [][]int, width, height int, i, j int, count int) int {
	if grid[i][j] != 1 {
		return count
	}
	grid[i][j] = 2
	count++
	if i+1 <= height {
		count = infect(grid, width, height, i+1, j, count)
	}

	if i-1 >= 0 {
		count = infect(grid, width, height, i-1, j, count)
	}

	if j-1 >= 0 {
		count = infect(grid, width, height, i, j-1, count)
	}

	if j+1 <= width {
		count = infect(grid, width, height, i, j+1, count)
	}
	return count
}

//leetcode submit region end(Prohibit modification and deletion)
