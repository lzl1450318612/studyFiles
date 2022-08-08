package main

import (
	"fmt"
	"strings"
)

//给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数 。
// 你可以对一个单词进行如下三种操作：
// 插入一个字符
// 删除一个字符 
// 替换一个字符 
// 
// 示例 1：
//输入：word1 = "horse", word2 = "ros"
//输出：3
//解释：
//horse -> rorse (将 'h' 替换为 'r')
//rorse -> rose (删除 'r')
//rose -> ros (删除 'e')
//
// 示例 2： 
//输入：word1 = "intention", word2 = "execution"
//输出：5
//解释：
//intention -> inention (删除 't')
//inention -> enention (将 'i' 替换为 'e')
//enention -> exention (将 'n' 替换为 'x')
//exention -> exection (将 'n' 替换为 'c')
//exection -> execution (插入 'u')

//leetcode submit region begin(Prohibit modification and deletion)
func minDistance(word1 string, word2 string) int {

	word1Arr := strings.Split(word1, "")
	word2Arr := strings.Split(word2, "")

	dp := make([]int, len(word1))

	flag := false
	for _, s := range word2Arr {
		if s == word1Arr[0] {
			flag = true
			break
		}
	}

	if flag {
		dp[0] = len(word2) - 1
	} else {
		dp[0] = len(word2)
	}

	for i := 1; i < len(word1); i++ {
		if i <= len(word2)-1 && word1[i] == word2[i] {
			dp[i] = dp[i-1] - 1
		} else if i <= len(word2)-1 {
			dp[i] = dp[i-1]
		} else {
			dp[i] = dp[i-1] + 1
		}
	}

	return dp[len(word1)-1]
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(minDistance("horse", "ros"))
}
