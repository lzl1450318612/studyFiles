package main

import (
	"fmt"
)

//给你一个整数数组 nums ，请你找出数组中乘积最大的非空连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
// 测试用例的答案是一个 32-位 整数。
// 子数组 是数组的连续子序列。
//
// 示例 1:
//输入: nums = [2,3,-2,4]
//输出: 6
//解释: 子数组 [2,3] 有最大乘积 6。
//
// 示例 2:
//输入: nums = [-2,0,-1]
//输出: 0
//解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。

//leetcode submit region begin(Prohibit modification and deletion)
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	maxRes := nums[0]
	temp := 1
	for _, num := range nums {
		temp *= num
		maxRes = max(maxRes, temp)
		if num == 0 {
			temp = 1
		}
	}

	temp = 1
	for i := len(nums) - 1; i >= 0; i-- {
		temp *= nums[i]
		maxRes = max(maxRes, temp)
		if nums[i] == 0 {
			temp = 1
		}
	}

	return maxRes
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(maxProduct([]int{-2, 0, -1}))
}
