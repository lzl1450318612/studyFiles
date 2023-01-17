package main

//用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的
//功能。(若队列中没有元素，deleteHead 操作返回 -1 )
//
// 示例 1：
//输入：
//["CQueue","appendTail","deleteHead","deleteHead","deleteHead"]
//[[],[3],[],[],[]]
//输出：[null,null,3,-1,-1]
//
// 示例 2：
//输入：
//["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]
//[[],[],[5],[2],[],[]]
//输出：[null,-1,null,null,5,2]

//leetcode submit region begin(Prohibit modification and deletion)
type CQueue struct {
	InputStack  []int
	OutputStack []int
}

func Constructor() CQueue {
	return CQueue{
		InputStack:  []int{},
		OutputStack: []int{},
	}

}

func (this *CQueue) AppendTail(value int) {
	this.InputStack = append(this.InputStack, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.OutputStack) != 0 {
		n := len(this.OutputStack)
		temp := this.OutputStack[n-1]
		this.OutputStack = this.OutputStack[:n-1]
		return temp
	}
	for len(this.InputStack) != 0 {
		this.OutputStack = append(this.OutputStack, this.InputStack[len(this.InputStack)-1])
		this.InputStack = this.InputStack[:len(this.InputStack)-1]
	}
	n := len(this.OutputStack)
	if n == 0 {
		return -1
	}
	temp := this.OutputStack[n-1]
	this.OutputStack = this.OutputStack[:n-1]
	return temp
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
//leetcode submit region end(Prohibit modification and deletion)
