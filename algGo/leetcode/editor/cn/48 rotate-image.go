package main

//给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。
// 你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。
//
// 示例 1： 
//输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
//输出：[[7,4,1],[8,5,2],[9,6,3]]
//
// 示例 2： 
//输入：matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
//输出：[[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]

//leetcode submit region begin(Prohibit modification and deletion)
func rotate(matrix [][]int) {
	if len(matrix) < 2 {
		return
	}

	// 转置
	for i := 0; i < len(matrix); i++ {
		for j := i + 1; j < len(matrix[0]); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// 镜像
	for i := 0; i < len(matrix); i++ {
		x := 0
		y := len(matrix[0]) - 1

		for x < y {
			matrix[i][x], matrix[i][y] = matrix[i][y], matrix[i][x]
			x++
			y--
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
