package main

import (
	"fmt"
	"sort"
)

//给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
// candidates 中的每个数字在每个组合中只能使用 一次 。
// 注意：解集不能包含重复的组合。
//
//示例 1:
//输入: candidates = [10,1,2,7,6,1,5], target = 8,
//输出:
//[
//[1,1,6],
//[1,2,5],
//[1,7],
//[2,6]
//] 
//
// 示例 2: 
//输入: candidates = [2,5,2,1,2], target = 5,
//输出:
//[
//[1,2,2],
//[5]
//] 

//leetcode submit region begin(Prohibit modification and deletion)
func combinationSum2(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}

	sort.Ints(candidates)

	res := make([][]int, 0)
	res = searchCombine2(candidates, target, make([]bool, len(candidates)), 0, []int{}, 0, res)
	return res
}

func searchCombine2(nums []int, target int, visited []bool, curVal int, curNums []int, startIdx int, res [][]int) [][]int {

	if curVal > target {
		return res
	}

	if curVal == target {
		temp := make([]int, 0, len(curNums))
		for _, num := range curNums {
			temp = append(temp, num)
		}

		res = append(res, temp)
		return res
	}

	for i := startIdx; i < len(nums); i++ {
		if visited[i] {
			continue
		}
		if i > startIdx && nums[i] == nums[i-1] {
			continue
		}

		visited[i] = true
		res = searchCombine2(nums, target, visited, curVal+nums[i], append(curNums, nums[i]), i+1, res)
		visited[i] = false
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(combinationSum2([]int{2, 5, 2, 1, 2}, 5))
}
