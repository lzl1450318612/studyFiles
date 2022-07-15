package main

import (
	"fmt"
	"strconv"
	"strings"
)

//给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，1 翻译成 “b”，……，11 翻译成 “l”，……，25 翻译成 “z”。一个数字可
//能有多个翻译。请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法。 
//
// 示例 1: 
// 输入: 12258
//输出: 5
//解释: 12258有5种不同的翻译，分别是"bccfi", "bwfi", "bczi", "mcfi"和"mzi" 

//leetcode submit region begin(Prohibit modification and deletion)
func translateNum(num int) int {

	numStr := strconv.Itoa(num)

	if len(numStr) < 2 {
		return len(numStr)
	}
	numStrArr := strings.Split(numStr, "")

	dp := make([]int, len(numStr))
	dp[0] = 1
	if numStrArr[0]+numStrArr[1] < "26" && numStrArr[0] != "0" {
		dp[1] = 2
	} else {
		dp[1] = 1
	}

	for i := 2; i < len(numStr); i++ {
		if numStrArr[i-1]+numStrArr[i] < "26" && numStrArr[i-1] != "0" {
			dp[i] = dp[i-1] + dp[i-2]
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[len(numStr)-1]
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(translateNum(26))
}
