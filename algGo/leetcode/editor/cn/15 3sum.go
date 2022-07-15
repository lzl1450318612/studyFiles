package main

import (
	"fmt"
	"math"
	"sort"
)

//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重
//复的三元组。 
// 注意：答案中不可以包含重复的三元组。
//
// 示例 1： 
//输入：nums = [-1,0,1,2,-1,-4]
//输出：[[-1,-1,2],[-1,0,1]]
//
// 示例 2： 
//输入：nums = []
//输出：[]
//
// 示例 3： 
//输入：nums = [0]
//输出：[]

//leetcode submit region begin(Prohibit modification and deletion)
func threeSum(nums []int) [][]int {
	res := make([][]int, 0)

	if len(nums) < 3 {
		return res
	}

	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if i != 0 && nums[i] == nums[i-1] {
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
			if sum > 0 {
				right--
			} else if sum < 0 {
				left++
			} else {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				lastLeftNum, lastRightNum = nums[left], nums[right]
				right--
			}
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
