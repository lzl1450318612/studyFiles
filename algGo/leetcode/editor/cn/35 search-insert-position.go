package main

import (
	"fmt"
)

//给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
// 请必须使用时间复杂度为 O(log n) 的算法。
//
// 示例 1: 
//输入: nums = [1,3,5,6], target = 5
//输出: 2
//
// 示例 2: 
//输入: nums = [1,3,5,6], target = 2
//输出: 1
//
// 示例 3: 
//输入: nums = [1,3,5,6], target = 7
//输出: 4
// nums 为 无重复元素 的 升序 排列数组

//leetcode submit region begin(Prohibit modification and deletion)
func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		}
		if nums[0] > target {
			return 0
		}
		return 1
	}

	return binarySearchInsert(nums, 0, len(nums)-1, target)
}

func binarySearchInsert(nums []int, startIdx, endIdx int, target int) int {

	midIdx := startIdx + (endIdx-startIdx)/2

	midVal := nums[midIdx]

	if target == midVal {
		return midIdx
	}

	if startIdx == endIdx {
		if target>nums[startIdx] {
			return startIdx+1
		}
		return startIdx
	}

	if midVal > target {
		return binarySearchInsert(nums, startIdx, midIdx, target)
	} else if midVal < target {
		return binarySearchInsert(nums, midIdx+1, endIdx, target)
	}

	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
}
