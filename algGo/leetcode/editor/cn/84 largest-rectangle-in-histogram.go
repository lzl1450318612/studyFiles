package main

import (
	"fmt"
)

//给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
// 求在该柱状图中，能够勾勒出来的矩形的最大面积。
//
// 示例 1: 
//输入：heights = [2,1,5,6,2,3]
//输出：10
//解释：最大的矩形为图中红色区域，面积为 10
//
// 示例 2： 
//输入： heights = [2,4]
//输出： 4 

//leetcode submit region begin(Prohibit modification and deletion)
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
	fmt.Println(largestRectangleArea([]int{1, 1}))
}
