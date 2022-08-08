package main

import (
	"fmt"
)

//有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
// 例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312"
//和 "192.168@1.1" 是 无效 IP 地址。 
// 给定一个只包含数字的字符串 s ，用以表示一个 IP 地址，返回所有可能的有效 IP 地址，这些地址可以通过在 s 中插入 '.' 来形成。你 不能 重新
//排序或删除 s 中的任何数字。你可以按 任何 顺序返回答案。 
//
// 示例 1： 
//输入：s = "25525511135"
//输出：["255.255.11.135","255.255.111.35"]
//
// 示例 2： 
//输入：s = "0000"
//输出：["0.0.0.0"]
//
// 示例 3： 
//输入：s = "101023"
//输出：["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]

//leetcode submit region begin(Prohibit modification and deletion)
func restoreIpAddresses(s string) []string {
	if len(s) > 12 || len(s) < 4 {
		return []string{}
	}

	return getIp(s, 0, 3, []string{}, []string{})

}

func getIp(s string, left int, dotCount int, curStr []string, res []string) []string {
	if dotCount == 0 {
		if !checkIp(s, left, len(s)-1) {
			return res
		}
		for i := left; i < len(s); i++ {
			curStr = append(curStr, string(s[i]))
		}

		temp := ""
		for _, s2 := range curStr {
			temp += s2
		}
		res = append(res, temp)
		return res
	}

	if checkIp(s, left, left) {
		curStr = append(curStr, string(s[left]))
		curStr = append(curStr, ".")
		res = getIp(s, left+1, dotCount-1, curStr, res)
	}

	if checkIp(s, left, left+1) {
		curStr = curStr[:len(curStr)-1]
		curStr = append(curStr, string(s[left+1]))
		curStr = append(curStr, ".")
		res = getIp(s, left+2, dotCount-1, curStr, res)
	}

	if checkIp(s, left, left+2) {
		curStr = curStr[:len(curStr)-1]
		curStr = append(curStr, string(s[left+2]))
		curStr = append(curStr, ".")
		res = getIp(s, left+3, dotCount-1, curStr, res)
	}

	return res
}

func checkIp(s string, left, right int) bool {
	if left > len(s)-1 || right > len(s)-1 {
		return false
	}

	if right-left != 0 && s[left] == '0' {
		return false
	}

	if right-left > 2 {
		return false
	}

	if right-left == 2 {
		if s[left] == '2' && (s[left+1] > '5' || (s[left+1] == '5' && s[left+2] > '5')) {
			return false
		}
		if s[left] > '2' {
			return false
		}
	}

	return true
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(restoreIpAddresses("1111"))
}
