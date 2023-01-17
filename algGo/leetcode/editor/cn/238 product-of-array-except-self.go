package main

//给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。
// 题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在 32 位 整数范围内。
// 请不要使用除法，且在 O(n) 时间复杂度内完成此题。
//
// 示例 1:
//输入: nums = [1,2,3,4]
//输出: [24,12,8,6]
//
// 示例 2:
//输入: nums = [-1,1,0,-3,3]
//输出: [0,0,9,0,0]
//
// 进阶：你可以在 O(1) 的额外空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组不被视为额外空间。）

//leetcode submit region begin(Prohibit modification and deletion)
func productExceptSelf(nums []int) []int {
	leftChengArr := make([]int, len(nums))
	rightChengArr := make([]int, len(nums))

	product := 1

	for i := 0; i < len(nums); i++ {
		leftChengArr[i] = product
		product *= nums[i]
	}

	product = 1

	for i := len(nums) - 1; i >= 0; i-- {
		rightChengArr[i] = product
		product *= nums[i]
	}

	res := make([]int, len(nums))

	for i := 0; i < len(res); i++ {
		res[i] = leftChengArr[i] * rightChengArr[i]
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
