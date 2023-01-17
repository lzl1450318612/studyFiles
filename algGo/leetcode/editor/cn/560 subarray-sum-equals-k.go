package main

import (
	"fmt"
)

//给你一个整数数组 nums 和一个整数 k ，请你返回 该数组中和为 k 的连续子数组的个数 。
//
// 示例 1：
//输入：nums = [1,1,1], k = 2
//输出：2
//
// 示例 2：
//输入：nums = [1,2,3], k = 3
//输出：2

//leetcode submit region begin(Prohibit modification and deletion)
var flag = false

func subarraySum(nums []int, k int) int {
	count, pre := 0, 0
	m := map[int]int{}
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if _, ok := m[pre-k]; ok {
			count += m[pre-k]
		}
		m[pre] += 1
	}
	return count
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(subarraySum([]int{1, -1, 0}, 0))
}
