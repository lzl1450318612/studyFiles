package main

import (
	"fmt"
)

//给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
// 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
// 此外，你可以假设该网格的四条边均被水包围。
//
// 示例 1： 
//输入：grid = [
//  ["1","1","1","1","0"],
//  ["1","1","0","1","0"],
//  ["1","1","0","0","0"],
//  ["0","0","0","0","0"]
//]
//输出：1
//
// 示例 2： 
//输入：grid = [
//  ["1","1","0","0","0"],
//  ["1","1","0","0","0"],
//  ["0","0","1","0","0"],
//  ["0","0","0","1","1"]
//]
//输出：3

//leetcode submit region begin(Prohibit modification and deletion)
func numIslands(grid [][]byte) int {
	height := len(grid) - 1
	width := len(grid[0]) - 1

	count := 0
	for i := 0; i <= height; i++ {
		for j := 0; j <= width; j++ {
			if grid[i][j] == '1' {
				count++
				grid = infect(grid, width, height, i, j)
			}
		}
	}
	return count
}

func infect(grid [][]byte, width, height int, i, j int) [][]byte {
	if grid[i][j] != '1' {
		return grid
	}
	grid[i][j] = '2'
	if i+1 <= height {
		infect(grid, width, height, i+1, j)
	}

	if i-1 >= 0 {
		infect(grid, width, height, i-1, j)
	}

	if j-1 >= 0 {
		infect(grid, width, height, i, j-1)
	}

	if j+1 <= width {
		infect(grid, width, height, i, j+1)
	}
	return grid
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(numIslands([][]byte{
		{1, 1, 1, 1, 0},
		{1, 1, 0, 1, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}))
}
