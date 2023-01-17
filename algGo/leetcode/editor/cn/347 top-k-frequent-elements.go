package main

import (
	"container/heap"
	"fmt"
)

//给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。
//
// 示例 1:
//输入: nums = [1,1,1,2,2,3], k = 2
//输出: [1,2]
//
// 示例 2:
//输入: nums = [1], k = 1
//输出: [1]
//
// 进阶：你所设计算法的时间复杂度 必须 优于 O(n log n) ，其中 n 是数组大小。

//leetcode submit region begin(Prohibit modification and deletion)
func topKFrequent(nums []int, k int) []int {
	res := make([]int, 0, k)
	countMap := make(map[int]int, 0)

	for _, num := range nums {
		countMap[num]++
	}

	countHeap := &CountHeap{
		CountArr: []int{},
		CountMap: countMap,
	}

	heap.Init(countHeap)

	for num, count := range countMap {
		if countHeap.Len() < k {
			heap.Push(countHeap, num)
			continue
		} else if countHeap.Len() == k {
			heapNum := heap.Pop(countHeap).(int)
			if countMap[heapNum] >= count {
				heap.Push(countHeap, heapNum)
				continue
			}
			heap.Push(countHeap, num)
		}
	}

	for countHeap.Len() > 0 {
		res = append(res, heap.Pop(countHeap).(int))
	}
	return res
}

type CountHeap struct {
	CountArr []int
	CountMap map[int]int
}

func (h CountHeap) Len() int {
	return len(h.CountArr)
}

func (h CountHeap) Less(i, j int) bool {
	return h.CountMap[h.CountArr[i]] < h.CountMap[h.CountArr[j]]
}

func (h *CountHeap) Swap(i, j int) {
	h.CountArr[i], h.CountArr[j] = h.CountArr[j], h.CountArr[i]
}

func (h *CountHeap) Push(x interface{}) {
	h.CountArr = append(h.CountArr, x.(int))
}

func (h *CountHeap) Pop() interface{} {
	old := (*h).CountArr
	n := len(old)
	x := old[n-1]
	(*h).CountArr = old[0 : n-1]
	return x
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}
