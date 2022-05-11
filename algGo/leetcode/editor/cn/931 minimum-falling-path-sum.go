package main

import (
	"fmt"
	"math"
)

//给你一个 n x n 的 方形 整数数组 matrix ，请你找出并返回通过 matrix 的下降路径 的 最小和 。
// 下降路径 可以从第一行中的任何元素开始，并从每一行中选择一个元素。在下一行选择的元素和当前行所选元素最多相隔一列（即位于正下方或者沿对角线向左或者向右的第
//一个元素）。具体来说，位置 (row, col) 的下一个元素应当是 (row + 1, col - 1)、(row + 1, col) 或者 (row + 1
//, col + 1) 。 
//
// 示例 1： 
//输入：matrix = [[2,1,3],[6,5,4],[7,8,9]]
//输出：13
//解释：如图所示，为和最小的两条下降路径
// 
//
// 示例 2： 
//输入：matrix = [[-19,57],[-40,-5]]
//输出：-59
//解释：如图所示，为和最小的下降路径
//
// 提示： 
// n == matrix.length == matrix[i].length
// 1 <= n <= 100 
// -100 <= matrix[i][j] <= 100 

//leetcode submit region begin(Prohibit modification and deletion)
func minFallingPathSum(matrix [][]int) int {
	height := len(matrix)
	width := len(matrix[0])

	for i := 1; i < height; i++ {
		for j := 0; j < width; j++ {
			if j == 0 {
				matrix[i][j] = min(matrix[i-1][j]+matrix[i][j], matrix[i-1][j+1]+matrix[i][j])
			} else if j == width-1 {
				matrix[i][j] = min(matrix[i-1][j-1]+matrix[i][j], matrix[i-1][j]+matrix[i][j])
			} else {
				matrix[i][j] = min(min(matrix[i-1][j-1]+matrix[i][j], matrix[i-1][j]+matrix[i][j]), matrix[i-1][j+1]+matrix[i][j])
			}
		}
	}

	ret := math.MaxInt32
	for _, a := range matrix[height-1] {
		if a < ret {
			ret = a
		}
	}

	return ret
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)

func init() {
	Register(912, func() {
		matrix := [][]int{{2, 3, 7}, {4, 6, 2}, {8, 1, 9}}
		fmt.Println(minFallingPathSum(matrix))
	})
}
