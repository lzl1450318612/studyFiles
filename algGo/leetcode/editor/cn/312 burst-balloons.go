package main

import (
	"fmt"
)

//有 n 个气球，编号为0 到 n - 1，每个气球上都标有一个数字，这些数字存在数组 nums 中。
// 现在要求你戳破所有的气球。戳破第 i 个气球，你可以获得 nums[i - 1] * nums[i] * nums[i + 1] 枚硬币。 这里的 i -
// 1 和 i + 1 代表和 i 相邻的两个气球的序号。如果 i - 1或 i + 1 超出了数组的边界，那么就当它是一个数字为 1 的气球。
// 求所能获得硬币的最大数量。
//
//示例 1：
//输入：nums = [3,1,5,8]
//输出：167
//解释：
//nums = [3,1,5,8] --> [3,5,8] --> [3,8] --> [8] --> []
//coins =  3*1*5    +   3*5*8   +  1*3*8  + 1*8*1 = 167
//
// 示例 2：
//输入：nums = [1,5]
//输出：10

//leetcode submit region begin(Prohibit modification and deletion)
func maxCoins(nums []int) int {
	var tmpNums []int
	tmpNums = append(tmpNums, 1)
	tmpNums = append(tmpNums, nums...)
	tmpNums = append(tmpNums, 1)
	nums = tmpNums

	dp := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, len(nums))
		for j := 0; j < len(nums); j++ {
			dp[i][j] = -1
		}
	}

	return processBalloon(nums, 0, len(nums)-1, dp)
}

func processBalloon(nums []int, leftIdx, rightIdx int, dp [][]int) int {

	if leftIdx >= rightIdx-1 {
		return 0
	}

	if dp[leftIdx][rightIdx] != -1 {
		return dp[leftIdx][rightIdx]
	}

	for midIdx := leftIdx + 1; midIdx < rightIdx; midIdx++ {
		dp[leftIdx][rightIdx] = max(dp[leftIdx][rightIdx], nums[leftIdx]*nums[midIdx]*nums[rightIdx]+processBalloon(nums, leftIdx, midIdx, dp)+processBalloon(nums, midIdx, rightIdx, dp))
	}
	return dp[leftIdx][rightIdx]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(maxCoins([]int{3, 1, 5, 8}))
}
