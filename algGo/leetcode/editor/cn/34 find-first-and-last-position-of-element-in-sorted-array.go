package main

import (
	"fmt"
)

//给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。
// 如果数组中不存在目标值 target，返回 [-1, -1]。
// 你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。
//
// 示例 1： 
//输入：nums = [5,7,7,8,8,10], target = 8
//输出：[3,4] 
//
// 示例 2： 
//输入：nums = [5,7,7,8,8,10], target = 6
//输出：[-1,-1] 
//
// 示例 3： 
//输入：nums = [], target = 0
//输出：[-1,-1] 
// nums 是一个非递减数组

//leetcode submit region begin(Prohibit modification and deletion)
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return []int{0, 0}
		}
		return []int{-1, -1}
	}

	return binarySearchRange(nums, 0, len(nums)-1, target)
}

func binarySearchRange(nums []int, startIdx, endIdx int, target int) []int {

	midIdx := startIdx + (endIdx-startIdx)/2

	if startIdx == endIdx && nums[midIdx] != target {
		if nums[midIdx] != target {
			return []int{-1, -1}
		}
		return []int{midIdx, midIdx}
	}

	if target < nums[midIdx] {
		return binarySearchRange(nums, startIdx, midIdx, target)
	} else if target > nums[midIdx] {
		return binarySearchRange(nums, midIdx+1, endIdx, target)
	} else if target == nums[midIdx] {
		left, right := midIdx, midIdx
		for i := midIdx; i >= 0; i-- {
			if nums[i] != target {
				break
			}
			left = i
		}

		for i := midIdx; i < len(nums); i++ {
			if nums[i] != target {
				break
			}
			right = i
		}

		return []int{left, right}
	}
	return []int{-1, -1}
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(searchRange([]int{1, 4}, 4))
}
