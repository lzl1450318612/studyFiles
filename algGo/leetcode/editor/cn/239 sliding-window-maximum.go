package main

import (
	"fmt"
)

//给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位
// 返回 滑动窗口中的最大值 。
//
// 示例 1：
//输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
//输出：[3,3,5,5,6,7]
//解释：
//滑动窗口的位置                最大值
//---------------               -----
//[1  3  -1] -3  5  3  6  7       3
// 1 [3  -1  -3] 5  3  6  7       3
// 1  3 [-1  -3  5] 3  6  7       5
// 1  3  -1 [-3  5  3] 6  7       5
// 1  3  -1  -3 [5  3  6] 7       6
// 1  3  -1  -3  5 [3  6  7]      7
//
// 示例 2：
//输入：nums = [1], k = 1
//输出：[1]

//leetcode submit region begin(Prohibit modification and deletion)
func maxSlidingWindow(nums []int, k int) []int {
	twoDirQueue := make([]int, 0, len(nums)-k+1)
	res := make([]int, 0, 1)

	for i := 0; i < k; i++ {
		for len(twoDirQueue) > 0 && nums[i] >= nums[twoDirQueue[len(twoDirQueue)-1]] {
			twoDirQueue = twoDirQueue[:len(twoDirQueue)-1]
		}
		twoDirQueue = append(twoDirQueue, i)
	}
	res = append(res, nums[twoDirQueue[0]])

	for i := k; i < len(nums); i++ {

		for len(twoDirQueue) > 0 && nums[i] >= nums[twoDirQueue[len(twoDirQueue)-1]] {
			twoDirQueue = twoDirQueue[:len(twoDirQueue)-1]
		}
		twoDirQueue = append(twoDirQueue, i)

		for len(twoDirQueue) != 0 && twoDirQueue[0] <= i-k {
			twoDirQueue = twoDirQueue[1:]
		}

		res = append(res, nums[twoDirQueue[0]])
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, 1, 2, 0, 5}, 3))
}
