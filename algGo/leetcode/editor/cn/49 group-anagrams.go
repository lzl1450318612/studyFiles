package main

import (
	"sort"
)

//给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
// 字母异位词 是由重新排列源单词的字母得到的一个新单词，所有源单词中的字母通常恰好只用一次。
//
// 示例 1: 
//输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
//输出: [["bat"],["nat","tan"],["ate","eat","tea"]] 
//
// 示例 2: 
//输入: strs = [""]
//输出: [[""]]
// 
// 示例 3:
//输入: strs = ["a"]
//输出: [["a"]] 

//leetcode submit region begin(Prohibit modification and deletion)
func groupAnagrams(strs []string) [][]string {
	flagMap := make(map[string][]string, 0)

	res := make([][]string, 0)

	for i := 0; i < len(strs); i++ {
		temp := []rune(strs[i])
		sort.Slice(temp, func(i, j int) bool {
			return temp[i] < temp[j]
		})

		str := string(temp)
		if _, ok := flagMap[str]; !ok {
			flagMap[str] = make([]string, 0, 1)
		}

		flagMap[str] = append(flagMap[str], strs[i])
	}

	for _, ss := range flagMap {
		temp := make([]string, 0, len(ss))
		for _, s := range ss {
			temp = append(temp, s)
		}
		res = append(res, temp)
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
