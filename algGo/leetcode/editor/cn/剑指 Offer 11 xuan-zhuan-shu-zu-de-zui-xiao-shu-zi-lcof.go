package main

import (
	"fmt"
)

//把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
// 给你一个可能存在 重复 元素值的数组 numbers ，它原来是一个升序排列的数组，并按上述情形进行了一次旋转。请返回旋转数组的最小元素。例如，数组
//[3,4,5,1,2] 为 [1,2,3,4,5] 的一次旋转，该数组的最小值为 1。
//
// 注意，数组 [a[0], a[1], a[2], ..., a[n-1]] 旋转一次 的结果为数组 [a[n-1], a[0], a[1], a[2], ..., a[n-2]] 。
//
// 示例 1：
//输入：numbers = [3,4,5,1,2]
//输出：1
//
// 示例 2：
//输入：numbers = [2,2,2,0,1]
//输出：0

//leetcode submit region begin(Prohibit modification and deletion)
func minArray(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	startIdx, endIdx := 0, len(numbers)-1

	for endIdx > startIdx {
		midIdx := startIdx + (endIdx-startIdx)/2

		if numbers[midIdx] > numbers[midIdx+1] {
			return numbers[midIdx+1]
		}

		if numbers[midIdx] > numbers[endIdx] {
			startIdx = midIdx
			continue
		}
		if numbers[midIdx] < numbers[startIdx] {
			endIdx = midIdx
			continue
		} else if numbers[startIdx] == numbers[endIdx] {
			endIdx--
		} else {
			return numbers[startIdx]
		}
	}

	return min(numbers[startIdx], numbers[endIdx])
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(minArray([]int{1, 3, 5}))
}
