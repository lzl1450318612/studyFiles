package main

import (
	"fmt"
	"sort"
)

//给定一个会议时间安排的数组 intervals ，每个会议时间都会包括开始和结束的时间 intervals[i] = [starti, endi] ，请你判断一个人是否能够参加这里面的全部会议。
//
//示例 1：
//输入：intervals = [[0,30],[5,10],[15,20]]
//输出：false
//
//示例 2：
//输入：intervals = [[7,10],[2,4]]
//输出：true

//leetcode submit region begin(Prohibit modification and deletion)
func canAttendMeetings(intervals [][]int) bool {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	flag := 0
	for _, interval := range intervals {
		if interval[1] < flag {
			return false
		}
		flag = interval[1]
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(canAttendMeetings([][]int{
		{7, 10},
		{2, 4},
	}))
}
