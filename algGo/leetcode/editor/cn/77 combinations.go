package main

import (
	"fmt"
)

//给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
// 你可以按 任何顺序 返回答案。
//
// 示例 1： 
//输入：n = 4, k = 2
//输出：
//[
//  [2,4],
//  [3,4],
//  [2,3],
//  [1,2],
//  [1,3],
//  [1,4],
//] 
//
// 示例 2： 
//输入：n = 1, k = 1
//输出：[[1]] 

//leetcode submit region begin(Prohibit modification and deletion)
func combine(n int, k int) [][]int {

	nums := make([]int, 0, n)

	for i := 1; i <= n; i++ {
		nums = append(nums, i)
	}

	res := make([][]int, 0)

	return dfsCombine(nums, k, 0, []int{}, res)
}

func dfsCombine(nums []int, count int, curIdx int, curNums []int, res [][]int) [][]int {
	if len(nums)-curIdx < count {
		return res
	}

	if count == 0 {
		res = append(res, curNums)
		return res
	}

	for i := curIdx; i < len(nums); i++ {
		count--
		temp := make([]int, len(curNums))
		copy(temp, curNums)
		temp = append(temp, nums[i])
		res = dfsCombine(nums, count, i+1, temp, res)
		count++
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(combine(4, 2))
}
