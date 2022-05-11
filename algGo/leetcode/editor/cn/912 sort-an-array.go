package main

import (
	"fmt"
)

//给你一个整数数组 nums，请你将该数组升序排列。
// 示例 1：
//输入：nums = [5,2,3,1]
//输出：[1,2,3,5]
// 
// 示例 2：
//输入：nums = [5,1,1,2,0,0]
//输出：[0,0,1,1,2,5]
// 
// Related Topics 数组 分治 桶排序 计数排序 基数排序 排序 堆（优先队列） 归并排序 👍 494 👎 0

//leetcode submit region begin(Prohibit modification and deletion)

func swap(arr []int, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

// 插入排序
//func sortArray(nums []int) []int {
//	if len(nums) < 2 {
//		return nums
//	}
//	for i := 1; i < len(nums); i++ {
//		for j := i; j > 0; j-- {
//			if nums[j] < nums[j-1] {
//				temp := nums[j]
//				nums[j] = nums[j-1]
//				nums[j-1] = temp
//			} else {
//				break
//			}
//		}
//	}
//	return nums
//}

// 选择排序
//func sortArray(nums []int) []int {
//	if len(nums) < 2 {
//		return nums
//	}
//	temp := 0
//	for i := 0; i < len(nums)-1; i++ {
//		for j := i; j < len(nums); j++ {
//			if nums[i] > nums[j] {
//				temp = nums[i]
//				nums[i] = nums[j]
//				nums[j] = temp
//			}
//		}
//	}
//	return nums
//}

// 归并排序
//func sortArray(nums []int) []int {
//	if len(nums) < 2 {
//		return nums
//	}
//
//	sort(nums, 0, len(nums)/2, len(nums)-1)
//
//	return nums
//}
//
//func sort(nums []int, startIdx, midIdx, endIdx int) {
//
//	if startIdx >= endIdx {
//		return
//	}
//
//	sort(nums, startIdx, startIdx+(midIdx-startIdx)/2, midIdx)
//
//	sort(nums, midIdx+1, midIdx+1+(endIdx-midIdx)/2, endIdx)
//
//	merge(nums, startIdx, midIdx, midIdx+1, endIdx)
//
//}
//
//func merge(nums []int, startIdx1, endIdx1, startIdx2, endIdx2 int) {
//
//	ret := make([]int, 0, endIdx1-startIdx1+1+endIdx2-startIdx2+1)
//
//	i, j := startIdx1, startIdx2
//
//	for i <= endIdx1 && j <= endIdx2 {
//		if nums[i] <= nums[j] {
//			ret = append(ret, nums[i])
//			i++
//		} else {
//			ret = append(ret, nums[j])
//			j++
//		}
//	}
//
//	for ; i <= endIdx1; i++ {
//		ret = append(ret, nums[i])
//	}
//
//	for ; j <= endIdx2; j++ {
//		ret = append(ret, nums[j])
//	}
//
//	for m, n := startIdx1, 0; m <= endIdx2 && n < len(ret); m++ {
//		nums[m] = ret[n]
//		n++
//	}
//}

// 堆排序
//func sortArray(nums []int) []int {
//	if len(nums) < 2 {
//		return nums
//	}
//
//	heapSize := len(nums)
//
//	for i := len(nums) - 1; i >= 0; i-- {
//		heapify(nums, i, heapSize)
//	}
//
//	//for i := 0; i < len(nums); i++ {
//	//	heapInsert(nums, i)
//	//}
//
//	heapSize--
//	swap(nums, 0, heapSize)
//	for heapSize > 0 {
//		heapify(nums, 0, heapSize)
//		heapSize--
//		swap(nums, 0, heapSize)
//	}
//
//	return nums
//}
//
//func heapInsert(arr []int, i int) {
//	if i == 0 {
//		return
//	}
//
//	for i > 0 && arr[(i-1)/2] < arr[i] {
//		swap(arr, (i-1)/2, i)
//		i = (i - 1) / 2
//	}
//
//}
//
//func heapify(arr []int, i, heapSize int) {
//	l := i*2 + 1
//	r := i*2 + 2
//	maxIdx := 0
//
//	for l < heapSize {
//		if r < heapSize && arr[l] < arr[r] {
//			maxIdx = r
//		} else {
//			maxIdx = l
//		}
//
//		if arr[maxIdx] > arr[i] {
//			swap(arr, maxIdx, i)
//			i = maxIdx
//			l = i*2 + 1
//			r = i*2 + 2
//		} else {
//			break
//		}
//	}
//}

// 快排
//func sortArray(nums []int) []int {
//	if len(nums) < 2 {
//		return nums
//	}
//
//	quickSort(nums, 0, len(nums)-1)
//
//	return nums
//}
//
//func quickSort(arr []int, start, end int) {
//
//	if start >= end || end >= len(arr) {
//		return
//	}
//
//	left, right := start, end
//	target := arr[start]
//
//	for left != right {
//		for left < right && arr[right] > target {
//			right--
//		}
//
//		for left < right && arr[left] <= target {
//			left++
//		}
//
//		if left < right {
//			arr[left], arr[right] = arr[right], arr[left]
//		}
//	}
//
//	arr[left], arr[start] = target, arr[left]
//
//	quickSort(arr, start, left-1)
//	quickSort(arr, left+1, end)
//}

// 基数排序（桶排序）
//func sortArray(nums []int) []int {
//	if len(nums) < 2 {
//		return nums
//	}
//
//	flag := true
//
//	for i := 1; flag; i = i * 10 {
//
//		flag = false
//
//		addBuckets := make([][]int, 10, 10)
//		subBuckets := make([][]int, 10, 10)
//
//		for _, num := range nums {
//			bucketIdx := num / i % 10
//			if bucketIdx != 0 {
//				flag = true
//			}
//
//			if bucketIdx >= 0 {
//				addBuckets[bucketIdx] = append(addBuckets[bucketIdx], num)
//
//			} else if bucketIdx < 0 {
//				subBuckets[bucketIdx*-1] = append(subBuckets[bucketIdx*-1], num)
//			}
//		}
//
//		nums = make([]int, 0, len(nums))
//
//		for j := len(subBuckets) - 1; j >= 0; j-- {
//			for _, num := range subBuckets[j] {
//				nums = append(nums, num)
//			}
//		}
//
//		for _, bucket := range addBuckets {
//			for _, num := range bucket {
//				nums = append(nums, num)
//			}
//		}
//	}
//
//	return nums
//}


// 基数排序优化
func sortArray(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}




	return nums
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	nums := []int{-2, 3, -5}
	fmt.Println(nums)
	fmt.Println(sortArray(nums))
}

//func init() {
//	Register(912, func() {
//		nums := []int{5, 1, 1, 2, 0, 0}
//		fmt.Println(nums)
//		fmt.Println(sortArray(nums))
//	})
//}
