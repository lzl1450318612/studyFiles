package main

import (
	"fmt"
)

//给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
//
// 示例 1：
//输入：nums = [1,5,11,5]
//输出：true
//解释：数组可以分割成 [1, 5, 5] 和 [11] 。
//
// 示例 2：
//输入：nums = [1,2,3,5]
//输出：false
//解释：数组不能分割成两个元素和相等的子集。

//leetcode submit region begin(Prohibit modification and deletion)

func canPartition(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	sum := 0

	for _, num := range nums {
		sum += num
	}

	if sum%2 != 0 {
		return false
	}

	target := sum / 2
	dp := make([]int, target+1)
	dp[0] = 1

	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}

	return dp[target] != 0
}

//
//func canPartition(nums []int) bool {
//	if len(nums) < 2 {
//		return false
//	}
//	return chooseBag(nums, 0, 0, 0)
//}
//
//func chooseBag(nums []int, idx int, lBag, rBag int) bool {
//	if idx >= len(nums) {
//		return lBag == rBag
//	}
//
//	return chooseBag(nums, idx+1, lBag+nums[idx], rBag) || chooseBag(nums, idx+1, lBag, rBag+nums[idx])
//}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(canPartition([]int{2, 2, 1, 1}))
}
