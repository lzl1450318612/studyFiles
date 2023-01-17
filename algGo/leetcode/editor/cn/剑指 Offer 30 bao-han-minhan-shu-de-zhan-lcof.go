package main

//定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，调用 min、push 及 pop 的时间复杂度都是 O(1)。
//
// 示例:
// MinStack minStack = new MinStack();
//minStack.push(-2);
//minStack.push(0);
//minStack.push(-3);
//minStack.min();   --> 返回 -3.
//minStack.pop();
//minStack.top();      --> 返回 0.
//minStack.min();   --> 返回 -2.
//
// 注意：本题与主站 155 题相同：https://leetcode-cn.com/problems/min-stack/

//leetcode submit region begin(Prohibit modification and deletion)
type MinStack struct {
	Stack []StackNode
}

type StackNode struct {
	Val    int
	CurMin int
}

func Constructor() MinStack {
	return MinStack{Stack: []StackNode{}}
}

func (this *MinStack) Push(x int) {
	if len(this.Stack) == 0 {
		this.Stack = append(this.Stack, StackNode{Val: x, CurMin: x})
		return
	}
	preNode := this.Stack[len(this.Stack)-1]
	if preNode.CurMin < x {
		this.Stack = append(this.Stack, StackNode{Val: x, CurMin: preNode.CurMin})
	} else {
		this.Stack = append(this.Stack, StackNode{Val: x, CurMin: x})
	}
}

func (this *MinStack) Pop() {
	if len(this.Stack) != 0 {
		this.Stack = this.Stack[:len(this.Stack)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.Stack) == 0 {
		return 0
	}
	return this.Stack[len(this.Stack)-1].Val
}

func (this *MinStack) Min() int {
	if len(this.Stack) == 0 {
		return 0
	}
	return this.Stack[len(this.Stack)-1].CurMin
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
//leetcode submit region end(Prohibit modification and deletion)
