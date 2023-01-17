package main

import (
	"fmt"
	"math"
)

//给你一个整数数组 nums ，你需要找出一个 连续子数组 ，如果对这个子数组进行升序排序，那么整个数组都会变为升序排序。
// 请你找出符合题意的 最短 子数组，并输出它的长度。
//
// 示例 1：
//输入：nums = [2,6,4,8,10,9,15]
//输出：5
//解释：你只需要对 [6, 4, 8, 10, 9] 进行升序排序，那么整个表都会变为升序排序。
//
// 示例 2：
//输入：nums = [1,2,3,4]
//输出：0
//
// 示例 3：
//输入：nums = [1]
//输出：0
//
// 进阶：你可以设计一个时间复杂度为 O(n) 的解决方案吗？

//leetcode submit region begin(Prohibit modification and deletion)
func findUnsortedSubarray(nums []int) int {

	if len(nums) < 2 {
		return 0
	}

	minNum, maxNum := math.MaxInt32, math.MinInt32

	lMin, rMax := len(nums)-1, 0

	for i := 0; i < len(nums); i++ {
		if nums[i] < maxNum {
			rMax = i
		}
		maxNum = max(maxNum, nums[i])
	}

	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] > minNum {
			lMin = i
		}
		minNum = min(minNum, nums[i])
	}

	if lMin > rMax {
		return 0
	}

	return rMax - lMin + 1
}

//func findUnsortedSubarray(nums []int) int {
//
//	lMin, rMax := math.MaxInt32, math.MinInt32
//
//	// 找到每个元素右边第一个比他大的元素
//	stack := make([]int, 0, len(nums))
//	for i := 0; i < len(nums); i++ {
//		if len(stack) == 0 {
//			stack = append(stack, i)
//			continue
//		}
//
//		for len(stack) != 0 && nums[i] < nums[stack[len(stack)-1]] {
//			lMin = min(nums[stack[len(stack)-1]], lMin)
//			stack = stack[:len(stack)-1]
//		}
//		stack = append(stack, i)
//	}
//
//	//for len(stack) != 0 {
//	//	rightMoreArr[stack[len(stack)-1]] = len(nums)
//	//	stack = stack[:len(stack)-1]
//	//}
//
//	// 找到每个元素左边第一个比他小的元素
//	stack = make([]int, 0, len(nums))
//	for i := len(nums) - 1; i >= 0; i-- {
//		if len(stack) == 0 {
//			stack = append(stack, i)
//			continue
//		}
//
//		for len(stack) != 0 && nums[i] > nums[stack[len(stack)-1]] {
//			rMax = max(nums[stack[len(stack)-1]], rMax)
//			stack = stack[:len(stack)-1]
//		}
//		stack = append(stack, i)
//	}
//
//	if lMin == math.MaxInt32 {
//		return 0
//	}
//
//	return rMax - lMin + 1
//}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(findUnsortedSubarray([]int{1, 2, 3, 4}))
}
