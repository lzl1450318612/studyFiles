package main

//从若干副扑克牌中随机抽 5 张牌，判断是不是一个顺子，即这5张牌是不是连续的。2～10为数字本身，A为1，J为11，Q为12，K为13，而大、小王为 0 ，
//可以看成任意数字。A 不能视为 14。 
//
// 示例 1: 
//输入: [1,2,3,4,5]
//输出: True 
//
// 示例 2: 
//输入: [0,0,1,2,5]
//输出: True 
//
// 限制： 
// 数组长度为 5
// 数组的数取值为 [0, 13] .
// Related Topics 数组 排序 👍 249 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func isStraight(nums []int) bool {
	if len(nums) != 5 {
		return false
	}

	nums, zeroCount, flag := sortArr(nums)
	if !flag {
		return false
	}

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == 0 {
			continue
		}
		if nums[i]+1 != nums[i+1] {
			if nums[i+1]-nums[i] > zeroCount+1 {
				return false
			}
			zeroCount--
		}
	}
	return true
}

func sortArr(nums []int) ([]int, int, bool) {
	arr := make([]int, 14)
	zeroCount := 0
	for _, num := range nums {
		if num == 0 {
			zeroCount++
			continue
		}
		if arr[num]!=0 {
			return nums, zeroCount, false
		}
		arr[num] = num
	}

	cursor := 0
	for ; cursor < zeroCount; cursor++ {
		nums[cursor] = 0
	}

	for _, a := range arr {
		if a != 0 {
			nums[cursor] = a
			cursor++
		}
	}
	return nums, zeroCount, true
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	isStraight([]int{10, 11, 0, 12, 6})
}
