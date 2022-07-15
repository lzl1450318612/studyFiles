package main

import (
	"fmt"
	"math"
	"sort"
)

//给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组 [nums[a], nums[
//b], nums[c], nums[d]] （若两个四元组元素一一对应，则认为两个四元组重复）： 
// 0 <= a, b, c, d < n
// a、b、c 和 d 互不相同 
// nums[a] + nums[b] + nums[c] + nums[d] == target 
// 你可以按 任意顺序 返回答案 。
//
// 示例 1： 
//输入：nums = [1,0,-1,0,-2,2], target = 0
//输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
//
// 示例 2： 
//输入：nums = [2,2,2,2,2], target = 8
//输出：[[2,2,2,2]]

//leetcode submit region begin(Prohibit modification and deletion)
func fourSum(nums []int, target int) [][]int {
	res := make([][]int, 0)

	sort.Ints(nums)

	if len(nums) < 4 {
		return res
	}

	for i := 0; i < len(nums)-3; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		res = findThree(nums, target-nums[i], i+1, nums[i], res)
	}
	return res
}

func findThree(nums []int, target, startIdx, curValue int, res [][]int) [][]int {
	for i := startIdx; i < len(nums)-2; i++ {
		if i != startIdx && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1

		lastLeftNum := math.MinInt32
		lastRightNum := math.MinInt32

		for left < right {
			if nums[left] == lastLeftNum {
				left++
				continue
			} else if nums[right] == lastRightNum {
				right--
				continue
			}

			sum := nums[i] + nums[left] + nums[right]
			if sum > target {
				right--
			} else if sum < target {
				left++
			} else {
				res = append(res, []int{curValue, nums[i], nums[left], nums[right]})
				lastLeftNum, lastRightNum = nums[left], nums[right]
				right--
			}
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(fourSum([]int{2, 2, 2, 2, 2}, 8))
}
