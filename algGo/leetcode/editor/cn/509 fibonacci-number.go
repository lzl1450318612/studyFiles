package main

import (
	"fmt"
)

//斐波那契数 （通常用 F(n) 表示）形成的序列称为 斐波那契数列 。该数列由 0 和 1 开始，后面的每一项数字都是前面两项数字的和。也就是：
//F(0) = 0，F(1) = 1
//F(n) = F(n - 1) + F(n - 2)，其中 n > 1
// 给定 n ，请计算 F(n) 。
//
// 示例 1：
//输入：n = 2
//输出：1
//解释：F(2) = F(1) + F(0) = 1 + 0 = 1
// 
// 示例 2：
//输入：n = 3
//输出：2
//解释：F(3) = F(2) + F(1) = 1 + 1 = 2
//
// 示例 3： 
//输入：n = 4
//输出：3
//解释：F(4) = F(3) + F(2) = 2 + 1 = 3

//leetcode submit region begin(Prohibit modification and deletion)
func fib(n int) int {
	if n < 2 {
		return n
	}
	p1, p2, sum := 0, 0, 1
	for i := 2; i <= n; i++ {
		p1 = p2
		p2 = sum
		sum = p1 + p2
	}

	return sum
}

//func fib(n int) int {
//	record := make(map[int]int, 0)
//
//	return feibo(n, record)
//}
//
//func feibo(n int, recordMap map[int]int) int {
//	if n == 1 || n == 0 {
//		return n
//	}
//
//	f1, ok1 := recordMap[n-1]
//	if !ok1 {
//		f1 = feibo(n-1, recordMap)
//	}
//	f2, ok2 := recordMap[n-2]
//	if !ok2 {
//		f2 = feibo(n-2, recordMap)
//	}
//
//	temp := f1 + f2
//
//	return temp
//}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(fib(4))
}
