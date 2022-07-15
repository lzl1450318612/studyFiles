package main

import (
	"fmt"
)

//给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
// 算法的时间复杂度应该为 O(log (m+n)) 。
//
// 示例 1： 
//输入：nums1 = [1,3], nums2 = [2]
//输出：2.00000
//解释：合并数组 = [1,2,3] ，中位数 2
//
// 示例 2： 
//输入：nums1 = [1,2], nums2 = [3,4]
//输出：2.50000
//解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5

//leetcode submit region begin(Prohibit modification and deletion)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 && len(nums2) == 0 {
		return 0
	} else if len(nums1) == 0 {
		if len(nums2)%2 == 0 {
			idx := len(nums2) / 2
			return float64(nums2[idx]+nums2[idx-1]) / 2
		}
		return float64(nums2[len(nums2)/2])
	} else if len(nums2) == 0 {
		if len(nums1)%2 == 0 {
			idx := len(nums1) / 2
			return float64(nums1[idx]+nums1[idx-1]) / 2
		}
		return float64(nums1[len(nums1)/2])
	}

	totalLen := len(nums1) + len(nums2)
	isSingle := true

	if totalLen%2 == 0 {
		isSingle = false
	}
	midCount := totalLen / 2
	if isSingle {
		midCount++
	}

	idx1 := 0
	idx2 := 0
	count := 1

	for idx1 < len(nums1) && idx2 < len(nums2) {
		count++
		if nums1[idx1] <= nums2[idx2] {
			idx1++
			if count == midCount {
				if isSingle {
					return float64(nums2[idx2])
				}
				return float64(nums1[idx1]+nums2[idx2]) / 2

			}
		} else {
			idx2++
			if count == midCount {
				if isSingle {
					return float64(nums1[idx1])
				}
				return float64(nums1[idx1]+nums2[idx2]) / 2
			}
		}

	}
	return float64(nums1[idx1]+nums2[idx2]) / 2
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(findMedianSortedArrays([]int{2,3,5}, []int{}))
}
