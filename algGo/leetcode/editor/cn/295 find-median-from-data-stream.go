package main

//中位数是有序列表中间的数。如果列表长度是偶数，中位数则是中间两个数的平均值。
// 例如，[2,3,4] 的中位数是 3
// [2,3] 的中位数是 (2 + 3) / 2 = 2.5
// 设计一个支持以下两种操作的数据结构：
// void addNum(int num) - 从数据流中添加一个整数到数据结构中。
// double findMedian() - 返回目前所有元素的中位数。 
//
// 示例： 
// addNum(1)
//addNum(2)
//findMedian() -> 1.5
//addNum(3) 
//findMedian() -> 2 
//
// 进阶: 
// 如果数据流中所有整数都在 0 到 100 范围内，你将如何优化你的算法？
// 如果数据流中 99% 的整数都在 0 到 100 范围内，你将如何优化你的算法？ 

//leetcode submit region begin(Prohibit modification and deletion)
import (
	"container/heap"
	"fmt"
)

type MedianFinder struct {
	SmallHeap *IntSmallHeap
	BigHeap   *IntBigHeap
}

func Constructor() MedianFinder {
	smallHeap := &IntSmallHeap{}
	bigHeap := &IntBigHeap{}

	heap.Init(smallHeap)
	heap.Init(bigHeap)

	return MedianFinder{
		SmallHeap: smallHeap,
		BigHeap:   bigHeap,
	}
}

func (this *MedianFinder) AddNum(num int) {
	smallHeap := this.SmallHeap
	bigHeap := this.BigHeap

	if smallHeap.Len() == 0 && bigHeap.Len() == 0 {
		heap.Push(smallHeap, num)
		return
	}

	smallVal := heap.Pop(smallHeap) 
	heap.Push(smallHeap, smallVal)

	if num > smallVal.(int) {
		heap.Push(smallHeap, num)
	} else {
		heap.Push(bigHeap, num)
	}

	if smallHeap.Len()-bigHeap.Len() >= 2 {
		heap.Push(bigHeap, heap.Pop(smallHeap) )
		return
	} else if bigHeap.Len()-smallHeap.Len() >= 2 {
		heap.Push(smallHeap, heap.Pop(bigHeap) )
		return
	}

}

func (this *MedianFinder) FindMedian() float64 {
	smallHeap := this.SmallHeap
	bigHeap := this.BigHeap

	if smallHeap.Len() > bigHeap.Len() {
		temp := heap.Pop(smallHeap) 
		heap.Push(smallHeap, temp)
		return float64(temp.(int))
	} else if bigHeap.Len() > smallHeap.Len() {
		temp := heap.Pop(bigHeap) 
		heap.Push(bigHeap, temp)
		return float64(temp.(int))
	} else {
		temp1 := heap.Pop(smallHeap) 
		heap.Push(smallHeap, temp1)

		temp2 := heap.Pop(bigHeap) 
		heap.Push(bigHeap, temp2)

		return float64(temp1.(int)+temp2.(int)) / 2
	}
}

type IntBigHeap []int

func (h IntBigHeap) Len() int {
	return len(h)
}

func (h IntBigHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntBigHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntBigHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntBigHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type IntSmallHeap []int

func (h IntSmallHeap) Len() int {
	return len(h)
}

func (h IntSmallHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntSmallHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntSmallHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntSmallHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	//obj1 := Constructor()
	//heap.Push(obj1.SmallHeap, -3)
	//heap.Push(obj1.SmallHeap, -5)
	//heap.Push(obj1.SmallHeap, -2)
	//heap.Push(obj1.SmallHeap, -4)
	//heap.Push(obj1.SmallHeap, -1)
	//heap.Push(obj1.SmallHeap, -8)
	//
	//fmt.Println(heap.Pop(obj1.SmallHeap))
	//
	//obj2 := Constructor()
	//heap.Push(obj2.SmallHeap, 3)
	//heap.Push(obj2.SmallHeap, 5)
	//heap.Push(obj2.SmallHeap, 2)
	//heap.Push(obj2.SmallHeap, 4)
	//heap.Push(obj2.SmallHeap, 1)
	//heap.Push(obj2.SmallHeap, 8)
	//fmt.Println(heap.Pop(obj2.SmallHeap))

	obj := Constructor()
	obj.AddNum(-1)
	obj.AddNum(-2)
	obj.AddNum(-3)
	obj.AddNum(-4)
	fmt.Println(obj.FindMedian())
	obj.AddNum(-5)
	fmt.Println(obj.FindMedian())
}
