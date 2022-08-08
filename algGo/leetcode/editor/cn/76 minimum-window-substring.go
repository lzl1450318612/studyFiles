package main

import (
	"fmt"
	"math"
	"strings"
)

//给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
// 注意：
// 对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
// 如果 s 中存在这样的子串，我们保证它是唯一的答案。 
//
// 示例 1： 
//输入：s = "ADOBECODEBANC", t = "ABC"
//输出："BANC"
//
// 示例 2： 
//输入：s = "a", t = "a"
//输出："a"
//
// 示例 3: 
//输入: s = "a", t = "aa"
//输出: ""
//解释: t 中两个字符 'a' 均应包含在 s 的子串中，
//因此没有符合条件的子字符串，返回空字符串。
//
//进阶：你能设计一个在 o(n) 时间内解决此问题的算法吗？

//leetcode submit region begin(Prohibit modification and deletion)
func minWindow(s string, t string) string {

	res := ""
	sMap := make(map[string]int, len(t))
	count := 0

	sArr := strings.Split(s, "")
	tArr := strings.Split(t, "")
	for _, s2 := range tArr {
		sMap[s2]++
		count++
	}

	if count != len(t) {
		return ""
	}

	l, r := 0, 0

	if sMap[sArr[l]] > 0 {
		sMap[sArr[l]]--
		count--
	}

	minResLen := math.MaxInt32
	resL := 0
	resR := len(sArr) - 1

	for r < len(s) {
		for count > 0 && r < len(s) {
			if sMap[sArr[r]] > 0 {
				sMap[sArr[r]]--
				count--
			}
			r++
		}

		for count == 0 && l <= r {
			if _, ok := sMap[sArr[l]]; ok {
				sMap[sArr[l]]++
				count++
			}
			l++
		}

		if r-l < minResLen {
			minResLen = r - l
			resL = l-1
			resR = r-1
		}
		r++
	}

	for i := resL; i <= resR; i++ {
		res += sArr[i]
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}
