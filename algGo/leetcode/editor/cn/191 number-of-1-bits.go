package main

//编写一个函数，输入是一个无符号整数（以二进制串的形式），返回其二进制表达式中数字位数为 '1' 的个数（也被称为汉明重量）。
//
// 示例 1：
//输入：00000000000000000000000000001011
//输出：3
//解释：输入的二进制串 00000000000000000000000000001011 中，共有三位为 '1'。
//
// 示例 2：
//输入：00000000000000000000000010000000
//输出：1
//解释：输入的二进制串 00000000000000000000000010000000 中，共有一位为 '1'。
//
// 示例 3：
//输入：11111111111111111111111111111101
//输出：31
//
//
// 进阶：
// 如果多次调用这个函数，你将如何优化你的算法？

//leetcode submit region begin(Prohibit modification and deletion)
func hammingWeight(num uint32) int {
	count := 0
	for num > 0 {
		num = num & (num - 1)
		count++
	}
	return count
}

//leetcode submit region end(Prohibit modification and deletion)
