package main

import (
	"fmt"
)

//给你一个非负整数 x ，计算并返回 x 的 算术平方根 。
// 由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。
// 注意：不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5 。
//
// 示例 1： 
//输入：x = 4
//输出：2
//
// 示例 2： 
//输入：x = 8
//输出：2
//解释：8 的算术平方根是 2.82842..., 由于返回类型是整数，小数部分将被舍去。

//leetcode submit region begin(Prohibit modification and deletion)
func mySqrt(x int) int {
	if x == 0 {
		return 0
	}

	if x < 4 {
		return 1
	}

	num := 2
	small, big := num, 0

	for small != big {
		if small==big-1 {
			return small
		}

		value := num * num

		if value < x {
			small = num
			num *= 2
		} else if value > x {
			big = num
			num = small + (num-small)/2
		} else {
			return num
		}
	}
	return num
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(mySqrt(8))
}
