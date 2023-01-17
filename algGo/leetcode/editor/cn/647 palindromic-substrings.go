package main

//给你一个字符串 s ，请你统计并返回这个字符串中 回文子串 的数目。
// 回文字符串 是正着读和倒过来读一样的字符串。
// 子字符串 是字符串中的由连续字符组成的一个序列。
// 具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。
//
// 示例 1：
//输入：s = "abc"
//输出：3
//解释：三个回文子串: "a", "b", "c"
//
// 示例 2：
//输入：s = "aaa"
//输出：6
//解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"

//leetcode submit region begin(Prohibit modification and deletion)
func countSubstrings(s string) int {

	dp := make([][]bool, len(s))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(s))
	}
	count := 0

	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] != s[j] {
				dp[i][j] = false
				continue
			}
			if j == i || j == i+1 {
				dp[i][j] = true
			} else {
				dp[i][j] = dp[i+1][j-1]
			}

			if dp[i][j] {
				count++
			}

		}

	}
	return count
}

//leetcode submit region end(Prohibit modification and deletion)
