package main

import (
	"fmt"
)

//给定一个包含 n + 1 个整数的数组 nums ，其数字都在 [1, n] 范围内（包括 1 和 n），可知至少存在一个重复的整数。
// 假设 nums 只有 一个重复的整数 ，返回 这个重复的数 。
// 你设计的解决方案必须 不修改 数组 nums 且只用常量级 O(1) 的额外空间。
//
// 示例 1：
//输入：nums = [1,3,4,2,2]
//输出：2
//
// 示例 2：
//输入：nums = [3,1,3,4,2]
//输出：3
//
// 进阶：
// 如何证明 nums 中至少存在一个重复的数字?
// 你可以设计一个线性级时间复杂度 O(n) 的解决方案吗？

//leetcode submit region begin(Prohibit modification and deletion)
func findDuplicate(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	slow, fast := 0, 0
	firstFlag := true

	for slow != fast || firstFlag {
		firstFlag = false
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	fast = nums[0]
	slow = nums[slow]

	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return fast
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 2}))
}
