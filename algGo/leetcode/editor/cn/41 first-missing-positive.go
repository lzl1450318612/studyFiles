package main

import (
	"fmt"
)

//给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
//请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。
//
// 示例 1： 
//输入：nums = [1,2,0]
//输出：3
// 
// 示例 2：
//输入：nums = [3,4,-1,1]
//输出：2
//
// 示例 3： 
//输入：nums = [7,8,9,11,12]
//输出：1

//leetcode submit region begin(Prohibit modification and deletion)
//func firstMissingPositive(nums []int) int {
//	flagMap := make(map[int]struct{}, len(nums))
//
//	minPositive := math.MaxInt32
//
//	for _, num := range nums {
//		if num < minPositive && num > 0 {
//			minPositive = num
//		}
//
//		flagMap[num] = struct{}{}
//	}
//
//	if _, ok := flagMap[1]; !ok {
//		minPositive = 1
//	}
//
//	for i := minPositive; ; i++ {
//		if _, ok := flagMap[i]; !ok {
//			return i
//		}
//	}
//}

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] != i + 1 {
			return i + 1
		}
	}
	return n + 1
}
//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(firstMissingPositive([]int{3, 4, -1, 1}))
}
