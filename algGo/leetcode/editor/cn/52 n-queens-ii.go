package main

import (
	"fmt"
	"math"
)

//n 皇后问题 研究的是如何将 n 个皇后放置在 n × n 的棋盘上，并且使皇后彼此之间不能相互攻击。
// 给你一个整数 n ，返回 n 皇后问题 不同的解决方案的数量。
//
// 示例 1： 
//输入：n = 4
//输出：2
//解释：如上图所示，4 皇后问题存在两个不同的解法。
//
// 示例 2： 
//输入：n = 1
//输出：1

//leetcode submit region begin(Prohibit modification and deletion)
func totalNQueens(n int) int {
	if n == 0 {
		return 0
	}

	rowResults := make([]int, n)
	for i := 0; i < len(rowResults); i++ {
		rowResults[i] = math.MinInt32
	}
	count := countNQueensSolution(rowResults, n, 0, 0)
	return count
}

func countNQueensSolution(rowResults []int, n, curLine, count int) int {

	if curLine == n {
		count++
		return count
	}

	for row := 0; row < n; row++ {
		f := true
		for j := 0; j < len(rowResults); j++ {
			if rowResults[j] == row || math.Abs(float64(j-curLine)) == math.Abs(float64(rowResults[j]-row)) {
				f = false
				break
			}
		}
		if !f {
			continue
		}
		rowResults[curLine] = row
		count = countNQueensSolution(rowResults, n, curLine+1, count)
		rowResults[curLine] = math.MinInt32
	}
	return count
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(totalNQueens(1))
}
