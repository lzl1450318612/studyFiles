package main

import (
	"fmt"
	"sort"
)

//给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
//
// 示例 1： 
//输入：nums = [1,1,2]
//输出：
//[[1,1,2],
// [1,2,1],
// [2,1,1]]
//
// 示例 2： 
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

//leetcode submit region begin(Prohibit modification and deletion)
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	perm := []int{}
	ans := make([][]int, 0)
	vis := make([]bool, n)
	var backtrack func(int)
	backtrack = func(idx int) {
		if idx == n {
			ans = append(ans, append([]int(nil), perm...))
			return
		}
		for i, v := range nums {
			if vis[i] || i > 0 && !vis[i-1] && v == nums[i-1] {
				continue
			}
			perm = append(perm, v)
			vis[i] = true
			backtrack(idx + 1)
			vis[i] = false
			perm = perm[:len(perm)-1]
		}
	}
	backtrack(0)
	return ans
}

func findAllNum(nums, curNums []int, visited []bool, res [][]int, idx int) [][]int {
	if idx == len(nums) {
		res = append(res, curNums)
		return res
	}

	for i, num := range nums {
		// 剪枝，如果和前一个字符相等，并且前一个字符在这次递归中没有访问过，说明可以用第一次字符出现的那次等效于这次递归，所以可省略
		if visited[i] || (i != 0 && !visited[i-1] && num == nums[i-1]) {
			continue
		}

		visited[i] = true
		curNums = append(curNums, num)
		res = findAllNum(nums, curNums, visited, res, idx+1)
		curNums = curNums[:len(curNums)-1]
		visited[i] = false
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(permuteUnique([]int{-1, 2, -1, 2, 1, -1, 2, 1}))
}
