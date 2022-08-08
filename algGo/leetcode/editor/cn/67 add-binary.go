package main

import (
	"fmt"
	"strings"
)

//给你两个二进制字符串，返回它们的和（用二进制表示）。
// 输入为 非空 字符串且只包含数字 1 和 0。
//
// 示例 1: 
// 输入: a = "11", b = "1"
//输出: "100" 
//
// 示例 2: 
// 输入: a = "1010", b = "1011"
//输出: "10101" 

//leetcode submit region begin(Prohibit modification and deletion)
func addBinary(a string, b string) string {
	aArr := strings.Split(a, "")
	bArr := strings.Split(b, "")

	temp := ""
	flag := false

	for len(aArr) != 0 && len(bArr) != 0 {
		aStr := aArr[len(aArr)-1]
		aArr = aArr[:len(aArr)-1]

		bStr := bArr[len(bArr)-1]
		bArr = bArr[:len(bArr)-1]

		if aStr == "0" && bStr == "0" {
			if flag {
				temp = "1" + temp
				flag = false
			} else {
				temp = "0" + temp
			}
		} else if aStr == "1" && bStr == "1" {
			if flag {
				temp = "1" + temp
			} else {
				temp = "0" + temp
			}
			flag = true
		} else {
			if flag {
				temp = "0" + temp
			} else {
				temp = "1" + temp
			}
		}
	}

	for len(aArr) != 0 {
		s := aArr[len(aArr)-1]
		aArr = aArr[:len(aArr)-1]

		if flag {
			if s == "0" {
				temp = "1" + temp
				flag = false
			} else {
				temp = "0" + temp
				flag = true
			}
		} else {
			temp = s + temp
		}
	}

	for len(bArr) != 0 {
		s := bArr[len(bArr)-1]
		bArr = bArr[:len(bArr)-1]

		if flag {
			if s == "0" {
				temp = "1" + temp
				flag = false
			} else {
				temp = "0" + temp
				flag = true
			}
		} else {
			temp = s + temp
		}
	}

	if flag {
		temp = "1" + temp
	}

	return temp
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(addBinary("1010", "1011"))
}
