package main

import (
	"strings"
)

//请实现一个函数，把字符串 s 中的每个空格替换成"%20"。
//
// 示例 1：
// 输入：s = "We are happy."
//输出："We%20are%20happy."

//leetcode submit region begin(Prohibit modification and deletion)
func replaceSpace(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
}

//leetcode submit region end(Prohibit modification and deletion)
