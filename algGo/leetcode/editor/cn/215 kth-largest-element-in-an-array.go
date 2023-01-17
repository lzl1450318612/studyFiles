package main

import (
	"container/heap"
	"fmt"
)

//给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
// 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
// 你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。
//
// 示例 1:
//输入: [3,2,1,5,6,4], k = 2
//输出: 5
//
// 示例 2:
//输入: [3,2,3,1,2,4,5,5,6], k = 4
//输出: 4

//leetcode submit region begin(Prohibit modification and deletion)
func findKthLargest(nums []int, k int) int {
	h := new(smallHeap)
	heap.Init(h)

	for i := 0; i < len(nums); i++ {
		if h.Len() == k {
			num := heap.Pop(h).(int)
			if nums[i] > num {
				heap.Push(h, nums[i])
			} else {
				heap.Push(h, num)
			}
		} else {
			heap.Push(h, nums[i])
		}
	}

	return heap.Pop(h).(int)
}

type smallHeap []int

func (h smallHeap) Len() int {
	return len(h)
}

func (h smallHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h smallHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *smallHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *smallHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

//
//func smallHeapInsert(nums []int, num int, size int) []int {
//	if len(nums) < size {
//		nums = append(nums, num)
//		for i := len(nums) / 2; i >= 0; i-- {
//			nums = smallHeapify(nums, i, len(nums))
//		}
//		return nums
//	}
//	if len(nums) > 0 && num > nums[0] {
//		nums[0] = num
//		nums = smallHeapify(nums, 0, size)
//	}
//	return nums
//}
//
//func smallHeapify(nums []int, i, size int) []int {
//	l := i*2 + 1
//	r := i*2 + 2
//	minIdx := i
//
//	if l <= size-1 && nums[minIdx] > nums[l] {
//		minIdx = l
//	}
//	if r <= size-1 && nums[minIdx] > nums[r] {
//		minIdx = r
//	}
//
//	if minIdx != i {
//		swap(nums, i, minIdx)
//		smallHeapify(nums, minIdx, size)
//	}
//
//	return nums
//}
//
//func swap(arr []int, i, j int) {
//	temp := arr[i]
//	arr[i] = arr[j]
//	arr[j] = temp
//}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(findKthLargest([]int{5, 2, 4, 1, 3, 6, 0}, 4))
}
