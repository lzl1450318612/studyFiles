package main

import (
	"fmt"
	"math"
)

//给你一个非负整数数组 nums ，你最初位于数组的第一个位置。
// 数组中的每个元素代表你在该位置可以跳跃的最大长度。
// 你的目标是使用最少的跳跃次数到达数组的最后一个位置。
// 假设你总是可以到达数组的最后一个位置。
//
// 示例 1: 
//输入: nums = [2,3,1,1,4]
//输出: 2
//解释: 跳到最后一个位置的最小跳跃数是 2。
//     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
//
// 示例 2: 
//输入: nums = [2,3,0,1,4]
//输出: 2

//leetcode submit region begin(Prohibit modification and deletion)
func jump(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = math.MaxInt32
	}

	dp[0] = 0

	for i := 0; i < len(nums); i++ {
		num := nums[i]

		for j := i + 1; j-i <= num && j < len(dp); j++ {
			dp[j] = min(dp[j], dp[i]+1)
		}
	}
	return dp[len(nums)-1]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(jump([]int{2, 3, 0, 1, 4}))
}
