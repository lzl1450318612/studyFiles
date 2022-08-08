package main

import (
	"fmt"
	"sort"
)

//给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。
// 解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。
//
// 示例 1： 
//输入：nums = [1,2,2]
//输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]
//
// 示例 2： 
//输入：nums = [0]
//输出：[[],[0]]

//leetcode submit region begin(Prohibit modification and deletion)
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	return getZiJi(nums, []int{}, 0, [][]int{})
}

func getZiJi(nums []int, curNums []int, curIdx int, res [][]int) [][]int {
	res = append(res, append([]int(nil), curNums...))

	for i := curIdx; i < len(nums); i++ {
		if i > curIdx && nums[i] == nums[i-1] {
			continue
		}

		temp := make([]int, len(curNums))
		copy(temp, curNums)
		temp = append(temp, nums[i])
		res = getZiJi(nums, temp, i+1, res)
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
}
