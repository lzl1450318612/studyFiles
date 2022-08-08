package main

import (
	"fmt"
	"strconv"
)

//给出集合 [1,2,3,...,n]，其所有元素共有 n! 种排列。
// 按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：
// "123"
// "132" 
// "213" 
// "231" 
// "312" 
// "321" 
// 给定 n 和 k，返回第 k 个排列。
//
// 示例 1：
//输入：n = 3, k = 3
//输出："213"
// 
// 示例 2：
//输入：n = 4, k = 9
//输出："2314"
//
// 示例 3： 
//输入：n = 3, k = 1
//输出："123"

//leetcode submit region begin(Prohibit modification and deletion)
func getPermutation(n int, k int) string {
	res := ""

	for n > 0 {
		if n == 1 {
			res+= strconv.Itoa(k)
			break
		}
		jiecheng := getJieCheng(n - 1)
		idx := k / jiecheng

		res += strconv.Itoa(idx + 1)

		n--
		k = k - jiecheng + len(res)
	}
	return res
}

func getJieCheng(n int) int {
	if n <= 2 {
		return n
	}

	res := 1

	for i := n; i > 1; i-- {
		res *= i
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(getPermutation(4, 9))
}
