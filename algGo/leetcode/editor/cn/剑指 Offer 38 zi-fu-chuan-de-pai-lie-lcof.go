package main

import (
	"fmt"
	"sort"
	"strings"
)

//输入一个字符串，打印出该字符串中字符的所有排列。
// 你可以以任意顺序返回这个字符串数组，但里面不能有重复元素。
//
// 示例: 
// 输入：s = "abc"
//输出：["abc","acb","bac","bca","cab","cba"]

//leetcode submit region begin(Prohibit modification and deletion)
func permutation(s string) []string {
	if len(s) < 2 {
		return []string{s}
	}

	chars := strings.Split(s, "")

	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})

	return findAllString(chars, make([]bool, len(s)), "", []string{})

}

func findAllString(chars []string, visitedArr []bool, curStr string, res []string) []string {
	if len(curStr) == len(chars) {
		res = append(res, curStr)
		return res
	}

	for i, char := range chars {
		if visitedArr[i] || (i != 0 && chars[i-1] == char && !visitedArr[i-1]) {
			continue
		}
		visitedArr[i] = true
		res = findAllString(chars, visitedArr, curStr+char, res)
		visitedArr[i] = false
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(permutation("aab"))
}
