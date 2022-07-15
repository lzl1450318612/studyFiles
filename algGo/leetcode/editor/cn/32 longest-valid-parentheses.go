package main

import (
	"fmt"
	"strings"
)

//给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
//
// 示例 1： 
//输入：s = "(()"
//输出：2
//解释：最长有效括号子串是 "()"
//
// 示例 2： 
//输入：s = ")()())"
//输出：4
//解释：最长有效括号子串是 "()()"
//
// 示例 3： 
//输入：s = ""
//输出：0

//leetcode submit region begin(Prohibit modification and deletion)
func longestValidParentheses(s string) int {
	if len(s) < 2 {
		return 0
	}

	// 用来存储左括号的idx
	stack := make([]int, 0, len(s))

	flagArr := make([]int, len(s))

	strArr := strings.Split(s, "")
	for i := 0; i < len(s); i++ {
		if strArr[i] == "(" {
			stack = append(stack, i)
		} else if strArr[i] == ")" {
			if len(stack) == 0 {
				flagArr[i] = 0
			} else {
				n := len(stack) - 1
				leftIdx := stack[n]
				stack = stack[:n]

				flagArr[leftIdx] = 1
				flagArr[i] = 1
			}
		}
	}

	maxLen := 0
	count := 0
	for i := 0; i < len(flagArr); i++ {
		if flagArr[i] == 1 {
			count++
		} else {
			if count > maxLen {
				maxLen = count
			}
			count = 0
		}
	}
	if count > maxLen {
		return count
	}

	return maxLen
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(longestValidParentheses("(()"))
}
