package main

import (
	"fmt"
	"strings"
)

//给定一个字符串 s 和一个字符串 t ，计算在 s 的子序列中 t 出现的个数。
// 字符串的一个 子序列 是指，通过删除一些（也可以不删除）字符且不干扰剩余字符相对位置所组成的新字符串。（例如，"ACE" 是 "ABCDE" 的一个子序列
//，而 "AEC" 不是） 
// 题目数据保证答案符合 32 位带符号整数范围。
//
// 示例 1： 
//输入：s = "rabbbit", t = "rabbit"
//输出：3
//解释：
//如下图所示, 有 3 种可以从 s 中得到 "rabbit" 的方案。
//rabbbit
//rabbbit
//rabbbit 
//
// 示例 2： 
//输入：s = "babgbag", t = "bag"
//输出：5
//解释：
//如下图所示, 有 5 种可以从 s 中得到 "bag" 的方案。 
//babgbag
//babgbag
//babgbag
//babgbag
//babgbag

//leetcode submit region begin(Prohibit modification and deletion)
func numDistinct(s string, t string) int {
	if len(s) < len(t) {
		return 0
	}
	return findSubChar(strings.Split(s, ""), 0, "", t, strings.Split(t, ""), 0)
}

func findSubChar(useChars []string, index int, currentStr, target string, targetChars []string, count int) int {

	if len(currentStr) == len(target) || index > len(useChars)-1 {
		if currentStr == target {
			count++
		}
		return count
	}

	curChar := useChars[index]

	count = findSubChar(useChars, index+1, currentStr, target, targetChars, count)

	if curChar != targetChars[len(currentStr)] {
		return count
	}
	count = findSubChar(useChars, index+1, currentStr+curChar, target, targetChars, count)

	return count
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(numDistinct("daacaedaceacabbaabdccdaaeaebacddadcaeaacadbceaecddecdeedcebcdacdaebccdeebcbdeaccabcecbeeaadbccbaeccbbdaeadecabbbedceaddcdeabbcdaeadcddedddcececbeeabcbecaeadddeddccbdbcdcbceabcacddbbcedebbcaccac", "ceadbaa"))
}
