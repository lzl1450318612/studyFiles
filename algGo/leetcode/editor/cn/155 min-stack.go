package main

import (
	"fmt"
	"math"
)

//设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
// 实现 MinStack 类:
//
// MinStack() 初始化堆栈对象。
// void push(int val) 将元素val推入堆栈。
// void pop() 删除堆栈顶部的元素。
// int top() 获取堆栈顶部的元素。
// int getMin() 获取堆栈中的最小元素。
//
// 示例 1:
//输入：
//["MinStack","push","push","push","getMin","pop","top","getMin"]
//[[],[-2],[0],[-3],[],[],[],[]]
//
//输出：
//[null,null,null,null,-3,null,0,-2]
//
//解释：
//MinStack minStack = new MinStack();
//minStack.push(-2);
//minStack.push(0);
//minStack.push(-3);
//minStack.getMin();   --> 返回 -3.
//minStack.pop();
//minStack.top();      --> 返回 0.
//minStack.getMin();   --> 返回 -2.

//leetcode submit region begin(Prohibit modification and deletion)
type MinStack struct {
	Stack  []int
	MinNum int
}

func Constructor() MinStack {
	return MinStack{Stack: make([]int, 0), MinNum: math.MinInt32}
}

func (this *MinStack) Push(val int) {
	if val < this.MinNum || len(this.Stack) == 0 {
		this.MinNum = val
	}
	this.Stack = append(this.Stack, val)
	this.Stack = append(this.Stack, this.MinNum)
}

func (this *MinStack) Pop() {
	if len(this.Stack) == 0 {
		return
	}
	if this.Stack[len(this.Stack)-2] == this.MinNum && len(this.Stack) > 2 {
		this.MinNum = this.Stack[len(this.Stack)-3]
	}
	this.Stack = this.Stack[:len(this.Stack)-2]
}

func (this *MinStack) Top() int {
	if len(this.Stack) == 0 {
		return 0
	}
	return this.Stack[len(this.Stack)-2]
}

func (this *MinStack) GetMin() int {
	return this.Stack[len(this.Stack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
//leetcode submit region end(Prohibit modification and deletion)
func main() {
	minStack := Constructor()
	minStack.Push(-10)
	minStack.Push(14)
	fmt.Println(minStack.GetMin())
	fmt.Println(minStack.GetMin())
	minStack.Push(-20)
	fmt.Println(minStack.GetMin())
	fmt.Println(minStack.GetMin())
	fmt.Println(minStack.Top())
	fmt.Println(minStack.GetMin())
	minStack.Pop()
	minStack.Push(10)
	minStack.Push(-7)
	fmt.Println(minStack.GetMin())
	minStack.Push(-7)
	minStack.Pop()
	fmt.Println(minStack.Top())
	fmt.Println(minStack.GetMin())
	minStack.Pop()
}
