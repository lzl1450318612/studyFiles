package main

import (
	"fmt"
)

//给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
// 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
//
// 示例 1： 
//输入：nums = [1,2,3]
//输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
//
// 示例 2： 
//输入：nums = [0]
//输出：[[],[0]]

//leetcode submit region begin(Prohibit modification and deletion)
func subsets(nums []int) [][]int {
	res := make([][]int, 0)

	return dfsSubsets(nums, 0, []int{}, res)
}

func dfsSubsets(nums []int, curIdx int, curNums []int, res [][]int) [][]int {

	res = append(res, curNums)

	for i := curIdx; i < len(nums); i++ {
		temp := make([]int, len(curNums))
		copy(temp, curNums)
		temp = append(temp, nums[i])
		res = dfsSubsets(nums, i+1, temp, res)
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
