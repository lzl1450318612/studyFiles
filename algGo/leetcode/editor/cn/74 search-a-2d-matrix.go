package main

import (
	"fmt"
)

//编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：
// 每行中的整数从左到右按升序排列。
// 每行的第一个整数大于前一行的最后一个整数。 
//
// 示例 1： 
//输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
//输出：true
//
// 示例 2： 
//输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
//输出：false

//leetcode submit region begin(Prohibit modification and deletion)
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	if len(matrix) == 1 && len(matrix[0]) == 1 {
		return matrix[0][0] == target
	}

	rowNum := rowBinarySearch(matrix, target, 0, len(matrix)-1)
	if rowNum == -1 {
		return false
	}

	return colBinarySearch(matrix, target, rowNum, 0, len(matrix[0])-1)
}

func rowBinarySearch(matrix [][]int, target int, startRow, endRow int) int {
	rowLen := len(matrix[0])

	midRow := startRow + (endRow-startRow)/2
	if target == matrix[midRow][rowLen-1] {
		return midRow
	}

	if startRow >= endRow {
		return startRow
	}

	if target > matrix[midRow][rowLen-1] {
		return rowBinarySearch(matrix, target, midRow+1, endRow)
	} else if target < matrix[midRow][rowLen-1] {
		return rowBinarySearch(matrix, target, startRow, midRow)
	}
	return startRow
}

func colBinarySearch(matrix [][]int, target int, row, startIdx, endIdx int) bool {
	midIdx := startIdx + (endIdx-startIdx)/2

	if target == matrix[row][midIdx] {
		return true
	}

	if startIdx >= endIdx {
		return false
	}

	if target > matrix[row][midIdx] {
		return colBinarySearch(matrix, target, row, midIdx+1, endIdx)
	} else if target < matrix[row][midIdx] {
		return colBinarySearch(matrix, target, row, startIdx, midIdx)
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}, 10))
}
