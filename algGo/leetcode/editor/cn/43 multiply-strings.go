package main

import (
	"fmt"
	"strconv"
)

//给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。
// 注意：不能使用任何内置的 BigInteger 库或直接将输入转换为整数。
//
// 示例 1: 
//输入: num1 = "2", num2 = "3"
//输出: "6" 
//
// 示例 2: 
//输入: num1 = "123", num2 = "456"
//输出: "56088" 

//leetcode submit region begin(Prohibit modification and deletion)
func multiply(num1 string, num2 string) string {
	if len(num1) == 0 || len(num2) == 0 {
		return ""
	}

	var upNumArr []rune
	var downNumArr []rune
	if len(num1) >= len(num2) {
		upNumArr = []rune(num1)
		downNumArr = []rune(num2)
	} else {
		upNumArr = []rune(num2)
		downNumArr = []rune(num1)
	}

	tempArr := make([]int, 0)
	for downIdx := len(downNumArr) - 1; downIdx >= 0; downIdx-- {
		jinwei := 0
		temp := 0
		times := 1
		for upIdx := len(upNumArr) - 1; upIdx >= 0; upIdx-- {
			upNum := int(upNumArr[upIdx] - '0')
			downNum := int(downNumArr[downIdx] - '0')

			res := upNum*downNum + jinwei
			jinwei = res / 10
			temp += res % 10 * times
			times *= 10
		}
		temp += jinwei * times
		tempArr = append(tempArr, temp)
	}

	solution := 0
	times := 1
	for i := 0; i < len(tempArr); i++ {
		solution += tempArr[i] * times
		times *= 10
	}

	return strconv.Itoa(solution)
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(multiply("9", "9"))
}
