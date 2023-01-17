package main

//输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。
//
// 示例 1：
// 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
//输出：[1,2,3,6,9,8,7,4,5]
//
// 示例 2：
// 输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
//输出：[1,2,3,4,8,12,11,10,9,5,6,7]
//
// 注意：本题与主站 54 题相同：https://leetcode-cn.com/problems/spiral-matrix/

//leetcode submit region begin(Prohibit modification and deletion)
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	m, n := len(matrix), len(matrix[0])

	visited := make([][]bool, len(matrix))
	for i := 0; i < len(matrix); i++ {
		visited[i] = make([]bool, len(matrix[i]))
	}

	res := make([]int, 0, m+n)

	x, y := 0, 0

	for x+1 < n && !visited[x+1][y] {
		for x < n && !visited[x][y] {
			res = append(res, matrix[x][y])
			visited[x][y] = true
			x++
		}

		for y < m && !visited[x][y] {
			res = append(res, matrix[x][y])
			visited[x][y] = true
			y++
		}

		for x >= 0 && !visited[x][y] {
			res = append(res, matrix[x][y])
			visited[x][y] = true
			x--
		}

		for y >= 0 {
			res = append(res, matrix[x][y])
			visited[x][y] = true
			y--
		}
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
