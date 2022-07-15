package main

import (
	"fmt"
)

//给你一个整数数组 nums 。玩家 1 和玩家 2 基于这个数组设计了一个游戏。
// 玩家 1 和玩家 2 轮流进行自己的回合，玩家 1 先手。开始时，两个玩家的初始分值都是 0 。每一回合，玩家从数组的任意一端取一个数字（即，nums[0
//] 或 nums[nums.length - 1]），取到的数字将会从数组中移除（数组长度减 1 ）。玩家选中的数字将会加到他的得分上。当数组中没有剩余数字可取
//时，游戏结束。 
// 如果玩家 1 能成为赢家，返回 true 。如果两个玩家得分相等，同样认为玩家 1 是游戏的赢家，也返回 true 。你可以假设每个玩家的玩法都会使他的分
//数最大化。 
//
// 示例 1：
//输入：nums = [1,5,2]
//输出：false
//解释：一开始，玩家 1 可以从 1 和 2 中进行选择。
//如果他选择 2（或者 1 ），那么玩家 2 可以从 1（或者 2 ）和 5 中进行选择。如果玩家 2 选择了 5 ，那么玩家 1 则只剩下 1（或者 2 ）
//可选。 
//所以，玩家 1 的最终分数为 1 + 2 = 3，而玩家 2 为 5 。
//因此，玩家 1 永远不会成为赢家，返回 false 。 
//
// 示例 2： 
//输入：nums = [1,5,233,7]
//输出：true
//解释：玩家 1 一开始选择 1 。然后玩家 2 必须从 5 和 7 中进行选择。无论玩家 2 选择了哪个，玩家 1 都可以选择 233 。
//最终，玩家 1（234 分）比玩家 2（12 分）获得更多的分数，所以返回 true，表示玩家 1 可以成为赢家。 

//leetcode submit region begin(Prohibit modification and deletion)
func PredictTheWinner(nums []int) bool {
	if len(nums) <= 2 || len(nums)%2 == 0 {
		return true
	}

	sum := 0
	for _, num := range nums {
		sum += num
	}

	dp := make([][]int, 0, len(nums))

	for i := 0; i < len(nums); i++ {
		dp = append(dp, make([]int, len(nums)))
	}

	for i := 0; i < len(nums); i++ {
		dp[i][i] = nums[i]
	}

	for i := 1; i < len(nums); i++ {
		dp[i-1][i] = max(nums[i-1], nums[i])
	}

	// 倒着遍历是因为i需要用到i+1的结果
	for i := len(nums) - 2; i >= 0; i-- {
		// 从i+2开始遍历j是因为相邻的不用动态规划，直接提前填充好base case
		for j := i + 2; j < len(nums); j++ {
			dp[i][j] = max(nums[i]+min(dp[i+2][j], dp[i+1][j-1]), nums[j]+min(dp[i+1][j-1], dp[i][j-2]))
		}
	}
	return dp[0][len(nums)-1] >= sum-dp[0][len(nums)-1]
}

//func PredictTheWinner(nums []int) bool {
//	if len(nums)%2 == 0 {
//		return true
//	}
//	a, b := findMaxScore(nums, 0, len(nums)-1, 0, 0, true)
//	return a >= b
//}
//
//func findMaxScore(nums []int, startIdx, endIdx int, curAScore, curBScore int, isATurn bool) (int, int) {
//	if startIdx > endIdx {
//		return curAScore, curBScore
//	}
//
//	if isATurn {
//		a1, b1 := findMaxScore(nums, startIdx+1, endIdx, curAScore+nums[startIdx], curBScore, false)
//		a2, b2 := findMaxScore(nums, startIdx, endIdx-1, curAScore+nums[endIdx], curBScore, false)
//		if a1 >= a2 {
//			return a1, b1
//		} else {
//			return a2, b2
//		}
//	} else {
//		a1, b1 := findMaxScore(nums, startIdx+1, endIdx, curAScore, curBScore+nums[startIdx], true)
//		a2, b2 := findMaxScore(nums, startIdx, endIdx-1, curAScore, curBScore+nums[endIdx], true)
//		if b1 >= b2 {
//			return a1, b1
//		} else {
//			return a2, b2
//		}
//	}
//}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(PredictTheWinner([]int{1, 1, 1}))
}
