package main

import (
	"fmt"
)

//给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否
//则返回 -1。 
//
//示例 1: 
// 输入: nums = [-1,0,3,5,9,12], target = 9
//输出: 4
//解释: 9 出现在 nums 中并且下标为 4
//
// 示例 2: 
// 输入: nums = [-1,0,3,5,9,12], target = 2
//输出: -1
//解释: 2 不存在 nums 中因此返回 -1

//leetcode submit region begin(Prohibit modification and deletion)
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	return findTwoFen(nums, target, 0, len(nums)-1)
}

func findTwoFen(nums []int, target int, startIdx, endIdx int) int {
	midIdx := startIdx + (endIdx-startIdx)/2
	if nums[midIdx] == target {
		return midIdx
	} else if startIdx == endIdx {
		return -1
	} else if nums[midIdx] > target {
		return findTwoFen(nums, target, startIdx, midIdx)
	} else {
		return findTwoFen(nums, target, midIdx+1, endIdx)
	}
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 2))
}
