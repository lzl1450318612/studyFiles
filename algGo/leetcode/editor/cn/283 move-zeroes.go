package main

import (
	"fmt"
)

//给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
// 请注意 ，必须在不复制数组的情况下原地对数组进行操作。
//
// 示例 1:
//输入: nums = [0,1,0,3,12]
//输出: [1,3,12,0,0]
//
// 示例 2:
//输入: nums = [0]
//输出: [0]
//
// 进阶：你能尽量减少完成的操作次数吗？

//leetcode submit region begin(Prohibit modification and deletion)
func moveZeroes(nums []int) {
	zeroPtr := -1
	for i, num := range nums {
		if num == 0 {
			zeroPtr = i
			break
		}
	}
	if zeroPtr == -1 {
		return
	}

	numPtr := zeroPtr + 1
	for numPtr <= len(nums)-1 {
		if nums[numPtr] == 0 {
			numPtr++
			continue
		}
		nums[zeroPtr], nums[numPtr] = nums[numPtr], nums[zeroPtr]
		zeroPtr++
		numPtr++
	}
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println()
}
