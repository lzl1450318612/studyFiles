package main

import (
	"strings"
)

//Trie（发音类似 "try"）或者说 前缀树 是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。这一数据结构有相当多的应用情景，例如自动补完和拼
//写检查。
// 请你实现 Trie 类：
// Trie() 初始化前缀树对象。
// void insert(String word) 向前缀树中插入字符串 word 。
// boolean search(String word) 如果字符串 word 在前缀树中，返回 true（即，在检索之前已经插入）；否则，返回
//false 。
// boolean startsWith(String prefix) 如果之前已经插入的字符串 word 的前缀之一为 prefix ，返回 true ；否
//则，返回 false 。
//
// 示例：
//输入
//["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
//[[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
//输出
//[null, null, true, false, true, null, true]
//
//解释
//Trie trie = new Trie();
//trie.insert("apple");
//trie.search("apple");   // 返回 True
//trie.search("app");     // 返回 False
//trie.startsWith("app"); // 返回 True
//trie.insert("app");
//trie.search("app");     // 返回 True

//leetcode submit region begin(Prohibit modification and deletion)
type Trie struct {
	letter  string
	isEnd   int
	nextMap map[string]*Trie
}

func Constructor() Trie {
	return Trie{
		letter:  "",
		isEnd:   0,
		nextMap: map[string]*Trie{},
	}
}

func (this *Trie) Insert(word string) {
	if word == "" {
		return
	}
	wordArr := strings.Split(word, "")
	node := this

	for _, letter := range wordArr {
		if _, ok := node.nextMap[letter]; !ok {
			node.nextMap[letter] = &Trie{
				letter:  letter,
				isEnd:   0,
				nextMap: map[string]*Trie{},
			}
		}
		node = node.nextMap[letter]
	}
	node.isEnd++
}

func (this *Trie) Search(word string) bool {
	if word == "" {
		return true
	}
	wordArr := strings.Split(word, "")
	node := this
	for _, letter := range wordArr {
		if _, ok := node.nextMap[letter]; !ok {
			return false
		}
		node = node.nextMap[letter]
	}
	return node.isEnd > 0
}

func (this *Trie) StartsWith(prefix string) bool {
	if prefix == "" {
		return true
	}
	prefixArr := strings.Split(prefix, "")
	node := this
	for _, letter := range prefixArr {
		if _, ok := node.nextMap[letter]; !ok {
			return false
		}
		node = node.nextMap[letter]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
//leetcode submit region end(Prohibit modification and deletion)
