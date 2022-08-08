package main

import (
	"fmt"
)

//给你一个正整数 n ，生成一个包含 1 到 n² 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
//
// 示例 1： 
//输入：n = 3
//输出：[[1,2,3],[8,9,4],[7,6,5]]
//
// 示例 2： 
//输入：n = 1
//输出：[[1]]

//leetcode submit region begin(Prohibit modification and deletion)
func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, n)
	}

	i, j := 0, -1

	num := 1

	for {
		flag := false
		for j < n-1 && res[i][j+1] == 0 {
			res[i][j+1] = num
			num++
			flag = true
			j++
		}

		for i < n-1 && res[i+1][j] == 0 {
			res[i+1][j] = num
			num++
			flag = true
			i++
		}

		for j > 0 && res[i][j-1] == 0 {
			res[i][j-1] = num
			num++
			flag = true
			j--
		}

		for i > 0 && res[i-1][j] == 0 {
			res[i-1][j] = num
			num++
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
	fmt.Println(generateMatrix(1))
}