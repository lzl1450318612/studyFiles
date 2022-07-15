package main

import (
	"fmt"
	"math"
)

//给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
// 计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
// 你可以认为每种硬币的数量是无限的。
//
// 示例 1： 
//输入：coins = [1, 2, 5], amount = 11
//输出：3 
//解释：11 = 5 + 5 + 1 
//
// 示例 2： 
//输入：coins = [2], amount = 3
//输出：-1 
//
// 示例 3： 
//输入：coins = [1], amount = 0
//输出：0

//leetcode submit region begin(Prohibit modification and deletion)
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)

	dp[0] = 0

	for x := 1; x < amount+1; x++ {
		minCount := math.MaxInt32
		for _, coin := range coins {
			if coin > x {
				continue
			}
			temp := dp[x-coin] + 1
			if temp < minCount {
				minCount = temp
			}
		}

		dp[x] = minCount
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}

	return dp[amount]
}

//func coinChange(coins []int, amount int) int {
//	sort.Ints(coins)
//	resMap := make(map[int]int, 0)
//	temp := coinProcess(coins, amount, resMap)
//	if temp == math.MaxInt32 {
//		return -1
//	}
//	return temp
//}
//
//func coinProcess(coins []int, target int, recordMap map[int]int) int {
//	if target == 0 {
//		return 0
//	}
//	if target < 0 {
//		return -1
//	}
//
//	minCount := math.MaxInt32
//	for i := 0; i < len(coins); i++ {
//		count := 0
//		if c, ok := recordMap[target-coins[i]]; !ok {
//			count = coinProcess(coins, target-coins[i], recordMap)
//		} else {
//			count = c
//		}
//		recordMap[target-coins[i]] = count
//
//		if count+1 < minCount && count != -1 {
//			minCount = count + 1
//		}
//	}
//
//	return minCount
//}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(coinChange([]int{2}, 3))
}
