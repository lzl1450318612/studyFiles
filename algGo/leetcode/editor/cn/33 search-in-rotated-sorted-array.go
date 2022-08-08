package main

import (
	"fmt"
)

//整数数组 nums 按升序排列，数组中的值 互不相同 。
// 在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k], nums[
//k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2
//,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。 
// 给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。
// 你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。
//
// 示例 1： 
//输入：nums = [4,5,6,7,0,1,2], target = 0
//输出：4
// 
// 示例 2：
//输入：nums = [4,5,6,7,0,1,2], target = 3
//输出：-1 
//
// 示例 3： 
//输入：nums = [1], target = 0
//输出：-1

//leetcode submit region begin(Prohibit modification and deletion)
func search(nums []int, target int) int {
	if len(nums) < 3 {
		for i, num := range nums {
			if num == target {
				return i
			}
		}
		return -1
	}

	return binarySearchReverseArr(nums, 0, len(nums)-1, target, false)
}

func binarySearchReverseArr(nums []int, startIdx, endIdx int, target int, isSort bool) int {
	midIdx := startIdx + (endIdx-startIdx)/2

	if target == nums[midIdx] {
		return midIdx
	}

	if startIdx == endIdx {
		return -1
	}

	if isSort {
		if target < nums[midIdx] {
			return binarySearchReverseArr(nums, startIdx, midIdx, target, true)
		} else {
			return binarySearchReverseArr(nums, midIdx+1, endIdx, target, true)
		}
	} else {
		// 左边有序
		if nums[startIdx] <= nums[midIdx] {
			// 如果在左边的范围
			if target <= nums[midIdx] && target >= nums[startIdx] {
				return binarySearchReverseArr(nums, startIdx, midIdx, target, true)
			} else {
				return binarySearchReverseArr(nums, midIdx+1, endIdx, target, false)
			}
		} else {
			// 右边有序
			// 如果在右边的范围
			if target <= nums[endIdx] && target >= nums[midIdx] {
				return binarySearchReverseArr(nums, midIdx+1, endIdx, target, true)
			} else {
				return binarySearchReverseArr(nums, startIdx, midIdx, target, false)
			}
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(search([]int{3, 1}, 1))
}
