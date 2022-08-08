package main

import (
	"fmt"
)

//给定一个仅包含 0 和 1 、大小为 rows x cols 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。 
//
// 示例 1： 
//输入：matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"]
//,["1","0","0","1","0"]]
//输出：6
//解释：最大矩形如上图所示。
// 
// 示例 2：
//输入：matrix = []
//输出：0
//
// 示例 3： 
//输入：matrix = [["0"]]
//输出：0
//
// 示例 4： 
//输入：matrix = [["1"]]
//输出：1
//
// 示例 5： 
//输入：matrix = [["0","0"]]
//输出：0

//leetcode submit region begin(Prohibit modification and deletion)
func maximalRectangle(matrix [][]byte) int {

	maxHeight := 0
	heights := make([]int, len(matrix[0]))

	for _, row := range matrix {
		for i, numByte := range row {
			// 对于每一行，把元素累加上，把二维数组转换为84题的矩形高度一维数组，这样就能套用84题的方法
			// 如果是0，就直接赋值0；如果是1则向height数组中累加
			if numByte-'0' == 1 {
				heights[i]++
			} else {
				heights[i] = 0
			}
		}

		maxHeight = max(largestRectangleArea(heights), maxHeight)
	}

	return maxHeight
}

func largestRectangleArea(heights []int) int {

	if len(heights) == 0 {
		return 0
	}

	if len(heights) == 1 {
		return heights[0]
	}

	// 在前后都添加一个0，处理边界case，使得所有元素都能出栈
	temp := make([]int, len(heights))
	copy(temp, heights)
	temp = append(temp, 0)
	heights = []int{0}
	heights = append(heights, temp...)

	stack := make([]int, 0, len(heights))
	leftHeight := make([]int, len(heights))
	rightHeight := make([]int, len(heights))

	// 单调栈 找出一个元素左右第一个比它小的元素
	for i := 0; i < len(heights); i++ {
		if len(stack) == 0 {
			stack = append(stack, i)
			continue
		}

		for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
			// 如果小于栈顶元素，就出栈并更新出栈元素左边第一个最小元素为新的栈顶元素，右边第一个最小元素为将要入栈的元素
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			rightHeight[idx] = i
			if len(stack) != 0 {
				leftHeight[idx] = stack[len(stack)-1]
			}
		}
		stack = append(stack, i)
	}

	maxRectangleArea := 0

	for i := 0; i < len(heights); i++ {
		// 依次计算以每个元素为最小高度的矩形面积，选出一个最大的即为答案
		maxRectangleArea = max(maxRectangleArea, (rightHeight[i]-leftHeight[i]-1)*heights[i])
	}

	return maxRectangleArea
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(maximalRectangle([][]byte{
		{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'},
	}))
}
