package main

import (
	"fmt"
)

//数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
//
// 示例 1： 
//输入：n = 3
//输出：["((()))","(()())","(())()","()(())","()()()"]
// 
// 示例 2：
//输入：n = 1
//输出：["()"]

//leetcode submit region begin(Prohibit modification and deletion)
func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	if n == 1 {
		return []string{"()"}
	}

	res := make([]string, 0)

	res = findKuoHao(n, n, "", []string{}, res)
	return res
}

func findKuoHao(left, right int, curStr string, stack []string, res []string) []string {
	if left == 0 && right == 0 && len(stack) == 0 {
		res = append(res, curStr)
	}

	if left != 0 {
		stack = append(stack, "(")
		res = findKuoHao(left-1, right, curStr+"(", stack, res)
		stack = stack[:len(stack)-1]
	}
	if right != 0 {
		if len(stack) != 0 {
			stack = stack[:len(stack)-1]
			res = findKuoHao(left, right-1, curStr+")", stack, res)
		}
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(generateParenthesis(4))
}
