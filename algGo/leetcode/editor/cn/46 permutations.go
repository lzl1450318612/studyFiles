package main

import (
	"fmt"
)

//给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
//
// 示例 1：
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
// 
// 示例 2：
//输入：nums = [0,1]
//输出：[[0,1],[1,0]]
//
// 示例 3： 
//输入：nums = [1]
//输出：[[1]]

//leetcode submit region begin(Prohibit modification and deletion)
func permute(nums []int) [][]int {
	res := make([][]int, 0)
	visited := make([]bool, len(nums))

	res = searchPermute(nums, visited, []int{}, res)
	return res
}

func searchPermute(nums []int, visited []bool, curNums []int, res [][]int) [][]int {
	if len(curNums) == len(nums) {
		res = append(res, curNums)
	}

	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue
		}

		visited[i] = true
		res = searchPermute(nums, visited, append(curNums, nums[i]), res)
		visited[i] = false
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
