package main

import (
	"strings"
)

//实现 strStr() 函数。
// 给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。如
//果不存在，则返回 -1 。 
//
// 说明： 
// 当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。
// 对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与 C 语言的 strstr() 以及 Java 的 indexOf() 定义相符。
//
// 示例 1： 
//输入：haystack = "hello", needle = "ll"
//输出：2
//
// 示例 2： 
//输入：haystack = "aaaaa", needle = "bba"
//输出：-1

//leetcode submit region begin(Prohibit modification and deletion)
func strStr(haystack string, needle string) int {

}

func GetNextArr(str string) []int {
	if len(str) == 0 {
		return []int{}
	}

	if len(str) == 1 {
		return []int{-1}
	}

	strArr := strings.Split(str, "")

	nextArr := make([]int, len(str))
	nextArr[0] = -1
	nextArr[1] = 0

	count := 0

	for i := 2; i < len(str); {
		if strArr[i-1] == strArr[count] {
			nextArr[i] = count + 1
			count++
			i++
		} else if count > 0 {
			count = nextArr[count]
		} else {
			nextArr[i] = 0
			i++
		}
	}

	return nextArr
}

//leetcode submit region end(Prohibit modification and deletion)
