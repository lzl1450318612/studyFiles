package main

import (
	"fmt"
)

//给定一个字符串 s 和一些 长度相同 的单词 words 。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。
// 注意子串要与 words 中的单词完全匹配，中间不能有其他字符 ，但不需要考虑 words 中单词串联的顺序。
//
// 示例 1： 
//输入：s = "barfoothefoobarman", words = ["foo","bar"]
//输出：[0,9]
//解释：
//从索引 0 和 9 开始的子串分别是 "barfoo" 和 "foobar" 。
//输出的顺序不重要, [9,0] 也是有效答案。
//
// 示例 2： 
//输入：s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]
//输出：[]
//
// 示例 3： 
//输入：s = "barfoofoobarthefoobarman", words = ["bar","foo","the"]
//输出：[6,9,12]

//leetcode submit region begin(Prohibit modification and deletion)
func findSubstring(s string, words []string) []int {

	wordLen := len(words[0])
	wordsLen := wordLen * len(words)

	if len(s) < wordsLen {
		return []int{}
	}

	res := make([]int, 0)

	wordTargetMap := make(map[string]int, 0)
	for _, word := range words {
		if _, ok := wordTargetMap[word]; !ok {
			wordTargetMap[word] = 0
		}
		wordTargetMap[word]++
	}

	endIdx := wordsLen - 1
	for startIdx := 0; endIdx < len(s); startIdx++ {
		curStr := s[startIdx : endIdx+1]
		if checkSubString(curStr, wordTargetMap, wordLen) {
			res = append(res, startIdx)
		}

		endIdx++

	}

	return res
}

func checkSubString(s string, wordMap map[string]int, wordLen int) bool {
	target := make(map[string]int, len(wordMap))

	for str, val := range wordMap {
		target[str] = val
	}

	startIdx := 0
	endIdx := wordLen - 1

	for endIdx < len(s) {

		subSubStr := s[startIdx : endIdx+1]

		if val, ok := target[subSubStr]; !ok || val <= 0 {
			return false
		} else {
			target[subSubStr]--

			startIdx += wordLen
			endIdx += wordLen
		}

	}

	for _, val := range target {
		if val > 0 {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word","good","best","word"}))
}
