package main

import (
	"strings"
)

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
// 有效字符串需满足：
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。 
//
// 示例 1： 
//输入：s = "()"
//输出：true
//
// 示例 2： 
//输入：s = "()[]{}"
//输出：true
//
// 示例 3： 
//输入：s = "(]"
//输出：false
//
// 示例 4： 
//输入：s = "([)]"
//输出：false
//
// 示例 5： 
//输入：s = "{[]}"
//输出：true 

//leetcode submit region begin(Prohibit modification and deletion)
func isValid(s string) bool {
	charArr := strings.Split(s, "")

	stack := make([]string, 0, len(s))

	for _, c := range charArr {
		if c == "{" || c == "(" || c == "[" {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				return false
			}

			flag := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if (c == "}" && flag != "{") || (c == "]" && flag != "[") || (c == ")" && flag != "(") {
				return false
			}
		}
	}
	if len(stack) != 0 {
		return false
	}

	return true
}

//leetcode submit region end(Prohibit modification and deletion)
