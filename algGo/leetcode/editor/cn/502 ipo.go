package main

//假设 力扣（LeetCode）即将开始 IPO 。为了以更高的价格将股票卖给风险投资公司，力扣 希望在 IPO 之前开展一些项目以增加其资本。 由于资源有限
//，它只能在 IPO 之前完成最多 k 个不同的项目。帮助 力扣 设计完成最多 k 个不同项目后得到最大总资本的方式。 
// 给你 n 个项目。对于每个项目 i ，它都有一个纯利润 profits[i] ，和启动该项目需要的最小资本 capital[i] 。
// 最初，你的资本为 w 。当你完成一个项目时，你将获得纯利润，且利润将被添加到你的总资本中。
// 总而言之，从给定项目中选择 最多 k 个不同项目的列表，以 最大化最终资本 ，并输出最终可获得的最多资本。
// 答案保证在 32 位有符号整数范围内。
//
// 示例 1： 
//输入：k = 2, w = 0, profits = [1,2,3], capital = [0,1,1]
//输出：4
//解释：
//由于你的初始资本为 0，你仅可以从 0 号项目开始。
//在完成后，你将获得 1 的利润，你的总资本将变为 1。
//此时你可以选择开始 1 号或 2 号项目。
//由于你最多可以选择两个项目，所以你需要完成 2 号项目以获得最大的资本。
//因此，输出最后最大化的资本，为 0 + 1 + 3 = 4。
//
// 示例 2： 
//输入：k = 3, w = 0, profits = [1,2,3], capital = [0,1,2]
//输出：6

//leetcode submit region begin(Prohibit modification and deletion)
import (
	"container/heap"
	"fmt"
	"sort"
)

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	if k == 0 || len(profits) == 0 || len(capital) == 0 {
		return 0
	}

	profitHeap := &ProfitHeap{}
	heap.Init(profitHeap)

	ipoNodes := make([]*IpoNode, 0, len(capital))
	for i := 0; i < len(capital); i++ {
		ipoNodes = append(ipoNodes, &IpoNode{
			Cost:   capital[i],
			Profit: profits[i],
		})
	}

	sort.Slice(ipoNodes, func(i, j int) bool {
		return ipoNodes[i].Cost < ipoNodes[j].Cost
	})

	cur := 0
	for k > 0 {
		for i := cur; i < len(ipoNodes); i++ {
			if ipoNodes[i].Cost > w {
				cur = i
				break
			}
			heap.Push(profitHeap, ipoNodes[i])
			cur++
		}

		if profitHeap.Len() == 0 {
			return w
		}
		w += heap.Pop(profitHeap).(*IpoNode).Profit
		k--
	}
	return w
}

type IpoNode struct {
	Cost   int
	Profit int
}

type ProfitHeap []*IpoNode

func (h ProfitHeap) Len() int {
	return len(h)
}

func (h ProfitHeap) Less(i, j int) bool {
	return h[i].Profit > h[j].Profit
}

func (h ProfitHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *ProfitHeap) Push(x interface{}) {
	*h = append(*h, x.(*IpoNode))
}

func (h *ProfitHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(findMaximizedCapital(10, 0, []int{1, 2, 3}, []int{0, 1, 2}))
}
