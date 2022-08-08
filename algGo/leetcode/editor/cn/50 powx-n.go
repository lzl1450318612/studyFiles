package main

import (
	"fmt"
)

//实现 pow(x, n) ，即计算 x 的整数 n 次幂函数（即，xⁿ ）。
//
// 示例 1： 
//输入：x = 2.00000, n = 10
//输出：1024.00000
// 
// 示例 2：
//输入：x = 2.10000, n = 3
//输出：9.26100
//
// 示例 3： 
//输入：x = 2.00000, n = -2
//输出：0.25000
//解释：2-2 = 1/22 = 1/4 = 0.25

//leetcode submit region begin(Prohibit modification and deletion)
func myPow(x float64, n int) float64 {
	res := float64(1)

	for i := n; i != 0; i /= 2 {
		if i%2 != 0 {
			res *= x
		}
		x *= x
	}

	if n < 0 {
		return 1 / res
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(myPow(2.0, 10))
}
