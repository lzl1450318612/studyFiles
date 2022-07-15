package main

import (
	"fmt"
	"strings"
)

//给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
//
//输入: s = "abcabcbb"
//输出: 3 
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//
//输入: s = "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//
//输入: s = "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
//
//输入: s = ""
//输出: 0

//leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	maxLen := 0
	strArr := strings.Split(s, "")
	l, r := 0, 0

	for l < len(strArr) && r < len(strArr) && maxLen < len(strArr)-l {
		r = l
		record := make(map[string]int, len(strArr))
		for r < len(strArr) {
			if _, ok := record[strArr[r]]; ok {
				if r-l > maxLen {
					maxLen = r - l
				}
				l++
				break
			}
			record[strArr[r]] = r
			r++
		}
	}

	if r-l > maxLen {
		return r - l
	}

	return maxLen
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(lengthOfLongestSubstring("dvdf"))
}
