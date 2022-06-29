package main

import (
	"fmt"
	"sort"
)

//给你输入若干形如 [begin, end] 的区间，代表若干会议的开始时间和结束时间，请你计算至少需要申请多少间会议室。
//比如给你输入 meetings = [[0,30],[5,10],[15,20]]，算法应该返回 2，因为后两个会议和第一个会议时间是冲突的，至少申请两个会议室才能让所有会议顺利进行。

//leetcode submit region begin(Prohibit modification and deletion)
func minMeetingRooms(meetings [][]int) int {
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	count := 1

	flag := 0
	for _, meeting := range meetings {
		if meeting[1] < flag {
			count++
		}
		flag = meeting[1]
	}
	return count
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(minMeetingRooms([][]int{
		{0, 30},
		{5, 10},
		{15, 20},
	}))
}
