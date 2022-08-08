package main

import (
	"fmt"
)

//给你一个 无重叠的 ，按照区间起始端点排序的区间列表。
// 在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。
//
// 示例 1： 
//输入：intervals = [[1,3],[6,9]], newInterval = [2,5]
//输出：[[1,5],[6,9]]
//
// 示例 2： 
//输入：intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
//输出：[[1,2],[3,10],[12,16]]
//解释：这是因为新的区间 [4,8] 与 [3,5],[6,7],[8,10] 重叠。 
//
// 示例 3： 
//输入：intervals = [], newInterval = [5,7]
//输出：[[5,7]]
// 
// 示例 4：
//输入：intervals = [[1,5]], newInterval = [2,3]
//输出：[[1,5]]
// 
// 示例 5：
//输入：intervals = [[1,5]], newInterval = [2,7]
//输出：[[1,7]]

//leetcode submit region begin(Prohibit modification and deletion)
func insert(intervals [][]int, newInterval []int) [][]int {
	firstBiggerIdx := -1

	for i := 0; i < len(intervals); i++ {
		if intervals[i][0] > newInterval[0] {
			firstBiggerIdx = i
			break
		}
	}

	if firstBiggerIdx != -1 {
		leftInterval := make([][]int, firstBiggerIdx)
		rightInterval := make([][]int, len(intervals)-firstBiggerIdx)

		copy(leftInterval, intervals[:firstBiggerIdx])
		copy(rightInterval, intervals[firstBiggerIdx:])

		leftInterval = append(leftInterval, newInterval)
		leftInterval = append(leftInterval, rightInterval...)
		intervals = leftInterval
	} else {
		intervals = append(intervals, newInterval)
	}

	res := make([][]int, 0)

	l := intervals[0][0]
	r := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > r {
			res = append(res, []int{l, r})
			l = intervals[i][0]
			r = intervals[i][1]
		} else if intervals[i][1] > r {
			r = intervals[i][1]
		}
	}

	res = append(res, []int{l, r})
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(insert([][]int{{1, 3}, {6, 9}}, []int{2, 5}))
}
