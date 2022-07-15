package main

import (
	"container/heap"
	"fmt"
)

//给定一个单词列表 words 和一个整数 k ，返回前 k 个出现次数最多的单词。
// 返回的答案应该按单词出现频率由高到低排序。如果不同的单词有相同出现频率， 按字典顺序 排序。
//
// 示例 1： 
//输入: words = ["i", "love", "leetcode", "i", "love", "coding"], k = 2
//输出: ["i", "love"]
//解析: "i" 和 "love" 为出现次数最多的两个单词，均为2次。
//    注意，按字母顺序 "i" 在 "love" 之前。
//
// 示例 2： 
//输入: ["the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"],
//k = 4
//输出: ["the", "is", "sunny", "day"]
//解析: "the", "is", "sunny" 和 "day" 是出现次数最多的四个单词，
//    出现次数依次为 4, 3, 2 和 1 次。
//
// 进阶：尝试以 O(n log k) 时间复杂度和 O(n) 空间复杂度解决。 

//leetcode submit region begin(Prohibit modification and deletion)
func topKFrequent(words []string, k int) []string {

	h := &wordHeap{}
	countMap := make(map[string]int, 0)

	for _, word := range words {
		if _, ok := countMap[word]; !ok {
			countMap[word] = 0
		} else {
			countMap[word]++
		}
	}

	for word, cnt := range countMap {
		heap.Push(h, &pair{
			word:  word,
			count: cnt,
		})
	}

	res := make([]string, 0, k)

	for len(res) < k && h.Len() > 0 {
		p := heap.Pop(h).(*pair)
		res = append(res, p.word)
	}

	return res
}

type pair struct {
	word  string
	count int
}

type wordHeap []*pair

func (h wordHeap) Len() int {
	return len(h)
}

func (h wordHeap) Less(i, j int) bool {
	a, b := h[i], h[j]
	return a.count > b.count || (a.count == b.count && a.word < b.word)
}

func (h wordHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *wordHeap) Push(x interface{}) {
	*h = append(*h, x.(*pair))
}

func (h *wordHeap) Pop() interface{} {
	old := *h
	v := old[len(old)-1]
	*h = old[:len(old)-1]
	return v
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(topKFrequent([]string{"i", "love", "leetcode", "i", "love", "coding"}, 2))
}
