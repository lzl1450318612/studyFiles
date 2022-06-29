package main

import (
	"fmt"
	"math"
)

//n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
// 给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
// 每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
//
// 示例 1： 
//输入：n = 4
//输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
//解释：如上图所示，4 皇后问题存在两个不同的解法。
//
// 示例 2：
//输入：n = 1
//输出：[["Q"]]

//leetcode submit region begin(Prohibit modification and deletion)
func solveNQueens(n int) [][]string {
	solutions := make([][]int, 0)
	flag := make([]int, n)
	for i := 0; i < len(flag); i++ {
		flag[i] = math.MinInt32
	}
	solutions = findNQueensSolution(flag, n, 0, solutions)
	return genQueenPattern(solutions, n)
}

func findNQueensSolution(flag []int, n, curLine int, solutions [][]int) [][]int {

	if curLine == n {
		temp := make([]int, 0, len(flag))
		for _, i := range flag {
			temp = append(temp, i)
		}
		solutions = append(solutions, temp)
		return solutions
	}

	for i := 0; i < n; i++ {
		f := true
		for j := 0; j < len(flag); j++ {
			if flag[j] == i || math.Abs(float64(j-curLine)) == math.Abs(float64(flag[j]-i)) {
				f = false
				break
			}
		}
		if !f {
			continue
		}
		flag[curLine] = i
		solutions = findNQueensSolution(flag, n, curLine+1, solutions)
		flag[curLine] = math.MinInt32
	}

	return solutions
}

func genQueenPattern(solutions [][]int, n int) [][]string {
	res := make([][]string, 0, n)

	for _, solution := range solutions {
		lines := make([]string, 0, len(solutions))
		for i := 0; i < len(solution); i++ {
			line := ""
			for j := 0; j < n; j++ {
				if j != solution[i] {
					line += "."
				} else {
					line += "Q"
				}
			}
			lines = append(lines, line)
		}
		res = append(res, lines)
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func printQueenPattern(pattern [][]string) {
	fmt.Println("--------------------")
	for _, p := range pattern {
		for _, s := range p {
			fmt.Print(s + " ")
		}
		fmt.Println("")
	}
}

func main() {
	printQueenPattern(solveNQueens(4))
}
