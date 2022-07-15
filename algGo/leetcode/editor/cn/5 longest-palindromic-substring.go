package main

import (
	"fmt"
	"strings"
)

//给你一个字符串 s，找到 s 中最长的回文子串。
// 示例 1：
//输入：s = "babad"
//输出："bab"
//解释："aba" 同样是符合题意的答案。
//
// 示例 2： 
//输入：s = "cbbd"
//输出："bb"
//
// 示例 3： 
//输入：s = "a"
//输出："a"
//
// 示例 4： 
//输入：s = "ac"
//输出："a"

//leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	strArr := strings.Split(s, "")

	dp := make([][]bool, 0)
	for i := 0; i < len(s); i++ {
		dp = append(dp, make([]bool, len(s)))
	}
	for i := 0; i < len(s); i++ {
		dp[i][i] = true
	}

	left := 0
	right := 0
	maxLen := 0

	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			if strArr[i] != strArr[j] {
				dp[i][j] = false
			} else {
				if j == i+1 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}

			if dp[i][j] && j-i > maxLen {
				maxLen = j - i
				left = i
				right = j
			}
		}
	}

	return s[left : right+1]
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(longestPalindrome("aaaa"))
}
