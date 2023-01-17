package main

import (
	"fmt"
)

//给你一个字符串 s 和一个字符串列表 wordDict 作为字典。请你判断是否可以利用字典中出现的单词拼接出 s 。
// 注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。
//
// 示例 1：
//输入: s = "leetcode", wordDict = ["leet", "code"]
//输出: true
//解释: 返回 true 因为 "leetcode" 可以由 "leet" 和 "code" 拼接成。
//
// 示例 2：
//输入: s = "applepenapple", wordDict = ["apple", "pen"]
//输出: true
//解释: 返回 true 因为 "applepenapple" 可以由 "apple" "pen" "apple" 拼接成。
//
// 示例 3：
//输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
//输出: false

//leetcode submit region begin(Prohibit modification and deletion)

func wordBreak(s string, wordDict []string) bool {
	if s == "" || len(wordDict) == 0 {
		return false
	}
	dp := make([]bool, len(s))
	wordMap := make(map[string]struct{})
	for _, word := range wordDict {
		wordMap[word] = struct{}{}
	}

	for word := range wordMap {
		if len(word) <= len(s) && word == s[0:len(word)] {
			dp[len(word)-1] = true
		}
	}

	for i := 0; i < len(dp); i++ {
		for word := range wordMap {
			if i >= len(word) && dp[i-len(word)] && s[i-len(word)+1:i+1] == word {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)-1]
}

//func wordBreak(s string, wordDict []string) bool {
//	if s == "" || len(wordDict) == 0 {
//		return false
//	}
//	return processWord(wordDict, s, "", 0)
//}
//
//func processWord(wordDict []string, target string, curStr string, curIdx int) bool {
//	if curStr == target {
//		return true
//	}
//	if len(curStr) >= len(target) {
//		return false
//	}
//
//	for i := 0; i < len(wordDict); i++ {
//		if curIdx+len(wordDict[i]) < len(target)-1 && wordDict[i] != target[curIdx:curIdx+len(wordDict[i])] {
//			continue
//		}
//		if processWord(wordDict, target, curStr+wordDict[i], curIdx+len(wordDict[i])) {
//			return true
//		}
//	}
//	return false
//}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(wordBreak("bb", []string{"a", "b", "bbb", "bbbb"}))
}
