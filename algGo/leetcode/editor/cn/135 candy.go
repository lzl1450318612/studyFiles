package main

import (
	"fmt"
)

//n 个孩子站成一排。给你一个整数数组 ratings 表示每个孩子的评分。
// 你需要按照以下要求，给这些孩子分发糖果：
// 每个孩子至少分配到 1 个糖果。
// 相邻两个孩子评分更高的孩子会获得更多的糖果。 
// 请你给每个孩子分发糖果，计算并返回需要准备的 最少糖果数目 。
//
// 示例 1： 
//输入：ratings = [1,0,2]
//输出：5
//解释：你可以分别给第一个、第二个、第三个孩子分发 2、1、2 颗糖果。
//
// 示例 2： 
//输入：ratings = [1,2,2]
//输出：4
//解释：你可以分别给第一个、第二个、第三个孩子分发 1、2、1 颗糖果。
//     第三个孩子只得到 1 颗糖果，这满足题面中的两个条件。 

//leetcode submit region begin(Prohibit modification and deletion)
func candy(ratings []int) int {
	candies := make([]int, len(ratings))

	for i := 0; i < len(candies); i++ {
		candies[i] = 1
	}

	for i := 1; i < len(ratings); i++ {
		if ratings[i-1] < ratings[i] {
			candies[i] = candies[i-1] + 1
		}
	}

	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i+1] < ratings[i] && candies[i] <= candies[i+1] {
			candies[i] = candies[i+1] + 1
		}
	}

	count := 0
	for _, i := range candies {
		count += i
	}

	return count
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(candy([]int{1, 0, 2}))
}
