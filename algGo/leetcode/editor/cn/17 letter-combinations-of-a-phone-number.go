package main

import (
	"fmt"
	"strings"
)

//给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
//
// 示例 1： 
//输入：digits = "23"
//输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
//
// 示例 2： 
//输入：digits = ""
//输出：[]
//
// 示例 3： 
//输入：digits = "2"
//输出：["a","b","c"]

//leetcode submit region begin(Prohibit modification and deletion)
var letterMap = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	res := make([]string, 0)

	if len(digits) == 0 {
		return res
	}

	strArr := strings.Split(digits, "")
	if len(strArr) == 1 {
		return letterMap[strArr[0]]
	}

	params := make([][]string, 0)
	for i := 0; i < len(strArr); i++ {
		params = append(params, letterMap[strArr[i]])
	}
	return letterProcess(params, 0, "", res)
}

func letterProcess(params [][]string, curLine int, curText string, res []string) []string {
	if curLine > len(params)-1 {
		res = append(res, curText)
		return res
	}

	lineTextArr := params[curLine]
	for i := 0; i < len(lineTextArr); i++ {
		res = letterProcess(params, curLine+1, curText+lineTextArr[i], res)
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(letterCombinations("234"))
}
