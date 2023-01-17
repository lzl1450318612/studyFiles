package main

import (
	"fmt"
	"strings"
)

//给定一个经过编码的字符串，返回它解码后的字符串。
// 编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
// 你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
// 此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。
//
// 示例 1：
//输入：s = "3[a]2[bc]"
//输出："aaabcbc"
//
// 示例 2：
//输入：s = "3[a2[c]]"
//输出："accaccacc"
//
// 示例 3：
//输入：s = "2[abc]3[cd]ef"
//输出："abcabccdcdcdef"
//
// 示例 4：
//输入：s = "abc3[cd]xyz"
//输出："abccdcdcdxyz"

//leetcode submit region begin(Prohibit modification and deletion)
func decodeString(s string) string {
	if len(s) == 0 {
		return ""
	}
	if !strings.Contains(s, "[") {
		return s
	}
	res, _ := decodeDfs(s, 0)
	return res
}

func decodeDfs(s string, startIdx int) (string, int) {
	res := ""
	leftKuohaoCount := 1
	count := 1
	for i := startIdx; i < len(s); i++ {
		if s[i] == '[' {
			leftKuohaoCount++
			temp := ""
			temp, i = decodeDfs(s, i+1)
			res += strings.Repeat(temp, count)
			count = 1
		} else if s[i] == ']' {
			leftKuohaoCount--
			if leftKuohaoCount == 0 {
				return res, i - 1
			}
		} else if s[i] >= '0' && s[i] <= '9' {
			count, i = getDigits(s, i)
		} else {
			res += strings.Repeat(string(s[i]), count)
			count = 1
		}
	}
	return res, len(s) - 1
}

func getDigits(s string, startIdx int) (int, int) {
	res := 0
	for startIdx <= len(s)-1 && s[startIdx] >= '0' && s[startIdx] <= '9' {
		res = res*10 + int(s[startIdx]-'0')
		startIdx++
	}
	return res, startIdx - 1
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(decodeString("3[z]2[2[y]pq4[2[jk]e1[f]]]ef"))
}
