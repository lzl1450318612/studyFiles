package main

//给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。
// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
// 例如，在下面的 3×4 的矩阵中包含单词 "ABCCED"（单词中的字母已标出）。
//
// 示例 1：
//输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word =
//"ABCCED"
//输出：true
//
// 示例 2：
//输入：board = [["a","b"],["c","d"]], word = "abcd"
//输出：false

//leetcode submit region begin(Prohibit modification and deletion)
func exist(board [][]byte, word string) bool {
	if len(board) == 0 {
		return false
	}

	visited := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		visited[i] = make([]bool, len(board[0]))
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(board, i, j, word, 0, visited) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, x, y int, word string, wordIdx int, visited [][]bool) bool {

	if wordIdx == len(word) {
		return true
	}

	if x > len(board)-1 || x < 0 || y > len(board[0])-1 || y < 0 {
		return false
	}

	if visited[x][y] {
		return false
	}

	if board[x][y] != word[wordIdx] {
		return false
	}

	visited[x][y] = true

	f := dfs(board, x+1, y, word, wordIdx+1, visited) || dfs(board, x, y+1, word, wordIdx+1, visited) || dfs(board, x-1, y, word, wordIdx+1, visited) || dfs(board, x, y-1, word, wordIdx+1, visited)
	visited[x][y] = false
	return f
}

//leetcode submit region end(Prohibit modification and deletion)
