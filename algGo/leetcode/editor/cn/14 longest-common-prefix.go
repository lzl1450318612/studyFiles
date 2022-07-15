package main

import (
	"fmt"
	"sort"
	"strings"
)

//编写一个函数来查找字符串数组中的最长公共前缀。
// 如果不存在公共前缀，返回空字符串 ""。
//
// 示例 1： 
//输入：strs = ["flower","flow","flight"]
//输出："fl"
//
// 示例 2： 
//输入：strs = ["dog","racecar","car"]
//输出：""
//解释：输入不存在公共前缀。 

//leetcode submit region begin(Prohibit modification and deletion)
func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	n := len(strs)
	sort.Strings(strs)

	startStr := strs[0]
	endStr := strs[n-1]

	if startStr == "" {
		return ""
	}

	startStrArr := strings.Split(startStr, "")
	endStrArr := strings.Split(endStr, "")

	res := ""

	for i := 0; i < len(startStr); i++ {
		if startStrArr[i] != endStrArr[i] {
			break
		}
		res += startStrArr[i]
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(longestCommonPrefix([]string{"ab", "a"}))
}
