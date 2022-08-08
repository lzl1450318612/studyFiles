package main

import (
	"fmt"
)

//给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。
//
// 示例 1： 
//输入：matrix = [[1,1,1],[1,0,1],[1,1,1]]
//输出：[[1,0,1],[0,0,0],[1,0,1]]
//
// 示例 2： 
//输入：matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
//输出：[[0,0,0,0],[0,4,5,0],[0,3,1,0]]
//
// 进阶： 
// 一个直观的解决方案是使用 O(mn) 的额外空间，但这并不是一个好的解决方案。
// 一个简单的改进方案是使用 O(m + n) 的额外空间，但这仍然不是最好的解决方案。 
// 你能想出一个仅使用常量空间的解决方案吗？ 

//leetcode submit region begin(Prohibit modification and deletion)
func setZeroes(matrix [][]int) {
	row0has0 := false
	col0has0 := false

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				if i == 0 && j == 0 {
					row0has0 = true
					col0has0 = true
					continue
				} else if i == 0 {
					row0has0 = true
					continue
				} else if j == 0 {
					col0has0 = true
					continue
				}
				matrix[0][j], matrix[i][0] = 0, 0
			}

		}
	}

	for i := 1; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			for j := 0; j < len(matrix[0]); j++ {
				matrix[i][j] = 0
			}
		}
	}

	for j := 1; j < len(matrix[0]); j++ {
		if matrix[0][j] == 0 {
			for i := 0; i < len(matrix); i++ {
				matrix[i][j] = 0
			}
		}
	}

	if row0has0 {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[0][j] = 0
		}
	}

	if col0has0 {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}

}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	res := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroes(res)

	fmt.Println(res)
}
