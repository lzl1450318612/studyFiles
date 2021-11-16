package main

import (
	"fmt"
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
//
// 0 <= s.length <= 5 * 10⁴ 
// s 由英文字母、数字、符号和空格组成 

//leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) int {

	if len(s) < 2 {
		return len(s)
	}

	maxLen := 0

	m := make(map[byte]int)

	left := 0
	right := 0

	flag := false

	for right < len(s) {
		if _, ok := m[s[right]]; ok {
			if a := right - left; a > maxLen {
				maxLen = a
			}
			left = m[s[right]] + 1
			flag = true
		}
		m[s[right]] = right
		right++

	}

	if a := right - left; !flag && a > maxLen {
		maxLen = a
	}

	return maxLen

}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(lengthOfLongestSubstring("aab"))
}
