package main

import (
	"fmt"
)

//整数数组的一个 排列 就是将其所有成员以序列或线性顺序排列。
// 例如，arr = [1,2,3] ，以下这些都可以视作 arr 的排列：[1,2,3]、[1,3,2]、[3,1,2]、[2,3,1] 。
// 整数数组的 下一个排列 是指其整数的下一个字典序更大的排列。更正式地，如果数组的所有排列根据其字典顺序从小到大排列在一个容器中，那么数组的 下一个排列 就
//是在这个有序容器中排在它后面的那个排列。如果不存在下一个更大的排列，那么这个数组必须重排为字典序最小的排列（即，其元素按升序排列）。 
// 例如，arr = [1,2,3] 的下一个排列是 [1,3,2] 。
// 类似地，arr = [2,3,1] 的下一个排列是 [3,1,2] 。 
// 而 arr = [3,2,1] 的下一个排列是 [1,2,3] ，因为 [3,2,1] 不存在一个字典序更大的排列。 
// 给你一个整数数组 nums ，找出 nums 的下一个排列。
// 必须 原地 修改，只允许使用额外常数空间。
//
// 示例 1： 
//输入：nums = [1,2,3]
//输出：[1,3,2]
//
// 示例 2： 
//输入：nums = [3,2,1]
//输出：[1,2,3]
// 
// 示例 3：
//输入：nums = [1,1,5]
//输出：[1,5,1]

//leetcode submit region begin(Prohibit modification and deletion)
func nextPermutation(nums []int) {
	numLen := len(nums)
	flag := false

	diffNumIdx := 0
	diffNum := 0
	for i := numLen - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			diffNumIdx = i - 1
			diffNum = nums[diffNumIdx]
			flag = true
			break
		}
	}

	if !flag {
		reverseArr(nums, 0, len(nums)-1)
		return
	}

	for i := numLen - 1; i > 0; i-- {
		if nums[i] > diffNum {
			nums[i], nums[diffNumIdx] = nums[diffNumIdx], nums[i]
			reverseArr(nums, diffNumIdx+1, len(nums)-1)
			return
		}

	}

}

func reverseArr(nums []int, i, j int) {
	if len(nums) < 2 {
		return
	}

	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	arr := []int{2, 3, 1}
	nextPermutation(arr)
	fmt.Println(arr)
}
