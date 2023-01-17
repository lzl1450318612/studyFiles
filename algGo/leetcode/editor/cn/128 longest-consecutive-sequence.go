package main

import (
	"fmt"
)

//给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
// 请你设计并实现时间复杂度为 O(n) 的算法解决此问题。
//
// 示例 1：
//输入：nums = [100,4,200,1,3,2]
//输出：4
//解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
//
// 示例 2：
//输入：nums = [0,3,7,2,5,8,4,6,0,1]
//输出：9

//leetcode submit region begin(Prohibit modification and deletion)
func longestConsecutive(nums []int) int {
	flag := make(map[int]struct{})
	for _, num := range nums {
		flag[num] = struct{}{}
	}

	maxLen := 0
	for num := range flag {
		if _, ok := flag[num-1]; ok {
			continue
		}
		numLen := 1
		for {
			if _, ok := flag[num+1]; ok {
				numLen++
				num++
			} else {
				if maxLen < numLen {
					maxLen = numLen
				}
				break
			}
		}
	}
	return maxLen
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
}