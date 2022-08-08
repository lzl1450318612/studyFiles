package main

import (
	"fmt"
)

//一个数组的 最小乘积 定义为这个数组中 最小值 乘以 数组的 和 。
// 比方说，数组 [3,2,5] （最小值是 2）的最小乘积为 2 * (3+2+5) = 2 * 10 = 20 。
// 给你一个正整数数组 nums ，请你返回 nums 任意 非空子数组 的最小乘积 的 最大值 。由于答案可能很大，请你返回答案对 10⁹ + 7 取余 的
//结果。 
// 请注意，最小乘积的最大值考虑的是取余操作 之前 的结果。题目保证最小乘积的最大值在 不取余 的情况下可以用 64 位有符号整数 保存。
// 子数组 定义为一个数组的 连续 部分。
//
// 示例 1： 
//输入：nums = [1,2,3,2]
//输出：14
//解释：最小乘积的最大值由子数组 [2,3,2] （最小值是 2）得到。
//2 * (2+3+2) = 2 * 7 = 14 。
//
// 示例 2： 
//输入：nums = [2,3,3,1,2]
//输出：18
//解释：最小乘积的最大值由子数组 [3,3] （最小值是 3）得到。
//3 * (3+3) = 3 * 6 = 18 。
//
// 示例 3： 
//输入：nums = [3,1,5,6,4,2]
//输出：60
//解释：最小乘积的最大值由子数组 [5,6,4] （最小值是 4）得到。
//4 * (5+6+4) = 4 * 15 = 60 。

//leetcode submit region begin(Prohibit modification and deletion)
func maxSumMinProduct(nums []int) int {
	stack := make([][]int, 0, len(nums))

	leftMax := make([]int, len(nums))
	rightMax := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if len(stack) == 0 {
			stack = append(stack, []int{i})
			continue
		}
		for len(stack) > 0 && nums[stack[len(stack)-1][0]] > nums[i] {
			nodeList := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			for _, idx := range nodeList {
				if len(stack) != 0 {
					leftMax[idx] = stack[len(stack)-1][len(stack[len(stack)-1])-1]
				}
				rightMax[idx] = i
			}
		}
		if len(stack) == 0 {
			stack = append(stack, []int{i})
		} else if nums[stack[len(stack)-1][0]] == nums[i] {
			stack[len(stack)-1] = append(stack[len(stack)-1], i)
			continue
		} else {
			stack = append(stack, []int{i})
		}
	}

	for len(stack) > 0 {
		ints := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for _, idx := range ints {
			if len(stack) != 0 {
				leftMax[idx] = stack[len(stack)-1][len(stack[len(stack)-1])-1]
			}
			rightMax[idx] = len(nums) - 1
		}
	}

	maxRes := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := leftMax[i]; j <= rightMax[i]; j++ {
			sum += nums[j]
		}
		if maxRes < sum*nums[i] {
			maxRes = sum * nums[i]
		}
	}

	return maxRes
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(maxSumMinProduct([]int{1, 2, 3, 2}))
}
