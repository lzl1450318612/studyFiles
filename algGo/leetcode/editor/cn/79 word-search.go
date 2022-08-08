package main

import (
	"fmt"
	"strings"
)

//给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。 
// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
//
// 示例 1： 
//输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word =
//"ABCCED"
//输出：true
//
// 示例 2： 
//输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word =
//"SEE"
//输出：true
//
// 示例 3： 
//输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word =
//"ABCB"
//输出：false
//
// 进阶：你可以使用搜索剪枝的技术来优化解决方案，使其在 board 更大的情况下可以更快解决问题？ 

//leetcode submit region begin(Prohibit modification and deletion)
func exist(board [][]byte, word string) bool {

	if len(board)*len(board[0]) < len(word) {
		return false
	}

	if len(word) == 0 || len(board) == 0 || len(board[0]) == 0 {
		return false
	}

	wordArr := strings.Split(word, "")

	visited := make([][]bool, len(board))

	for i := 0; i < len(board); i++ {
		visited[i] = make([]bool, len(board[0]))
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if string(board[i][j]) == wordArr[0] {
				visited[i][j] = true
				if dfsMatrixExist(board, i, j, wordArr, 0, visited) {
					return true
				}
				visited[i][j] = false
			}
		}
	}
	return false
}

func dfsMatrixExist(board [][]byte, curRow, curCol int, wordArr []string, wordIdx int, visited [][]bool) bool {
	if wordArr[wordIdx] != string(board[curRow][curCol]) {
		return false
	}

	if wordIdx >= len(wordArr)-1 {
		return true
	}

	if curRow > len(board)-1 || curRow < 0 {
		return false
	}

	if curCol > len(board[0])-1 || curCol < 0 {
		return false
	}

	a, b, c, d := false, false, false, false

	if curRow > 0 && !visited[curRow-1][curCol] {
		visited[curRow-1][curCol] = true
		a = dfsMatrixExist(board, curRow-1, curCol, wordArr, wordIdx+1, visited)
		visited[curRow-1][curCol] = false
	}

	if curRow < len(board)-1 && !visited[curRow+1][curCol] {
		visited[curRow+1][curCol] = true
		b = dfsMatrixExist(board, curRow+1, curCol, wordArr, wordIdx+1, visited)
		visited[curRow+1][curCol] = false
	}

	if curCol > 0 && !visited[curRow][curCol-1] {
		visited[curRow][curCol-1] = true
		c = dfsMatrixExist(board, curRow, curCol-1, wordArr, wordIdx+1, visited)
		visited[curRow][curCol-1] = false
	}

	if curCol < len(board[0])-1 && !visited[curRow][curCol+1] {
		visited[curRow][curCol+1] = true
		d = dfsMatrixExist(board, curRow, curCol+1, wordArr, wordIdx+1, visited)
		visited[curRow][curCol+1] = false
	}

	return a || b || c || d
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(exist([][]byte{{'a', 'a'}, {'A', 'A'}}, "aaa"))
}
