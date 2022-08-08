package main

import (
	"fmt"
)

//给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
//
// 示例 1： 
//输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
//输出：[1,2,3,6,9,8,7,4,5]
//
// 示例 2： 
//输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
//输出：[1,2,3,4,8,12,11,10,9,5,6,7]

//leetcode submit region begin(Prohibit modification and deletion)
func spiralOrder(matrix [][]int) []int {
	width := len(matrix[0])
	height := len(matrix)

	visited := make([][]bool, height)
	for i := 0; i < height; i++ {
		visited[i] = make([]bool, width)
	}

	i, j := 0, -1

	res := make([]int, 0, width*height)

	for {
		flag := false
		for j < width-1 && !visited[i][j+1] {
			res = append(res, matrix[i][j+1])
			visited[i][j+1] = true
			flag = true
			j++
		}

		for i < height-1 && !visited[i+1][j] {
			res = append(res, matrix[i+1][j])
			visited[i+1][j] = true
			flag = true
			i++
		}

		for j > 0 && !visited[i][j-1] {
			res = append(res, matrix[i][j-1])
			visited[i][j-1] = true
			flag = true
			j--
		}

		for i > 0 && !visited[i-1][j] {
			res = append(res, matrix[i-1][j])
			visited[i-1][j] = true
			flag = true
			i--
		}
		if !flag {
			break
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(spiralOrder([][]int{
		{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12},
	}))
}
