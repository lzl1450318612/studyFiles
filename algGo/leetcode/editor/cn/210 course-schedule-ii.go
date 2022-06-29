package main

import (
	"fmt"
)

//现在你总共有 numCourses 门课需要选，记为 0 到 numCourses - 1。给你一个数组 prerequisites ，其中
//prerequisites[i] = [ai, bi] ，表示在选修课程 ai 前 必须 先选修 bi 。 
// 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示：[0,1] 。
// 返回你为了学完所有课程所安排的学习顺序。可能会有多个正确的顺序，你只要返回 任意一种 就可以了。如果不可能完成所有课程，返回 一个空数组 。
//
// 示例 1： 
//输入：numCourses = 2, prerequisites = [[1,0]]
//输出：[0,1]
//解释：总共有 2 门课程。要学习课程 1，你需要先完成课程 0。因此，正确的课程顺序为 [0,1] 。
//
// 示例 2： 
//输入：numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
//输出：[0,2,1,3]
//解释：总共有 4 门课程。要学习课程 3，你应该先完成课程 1 和课程 2。并且课程 1 和课程 2 都应该排在课程 0 之后。
//因此，一个正确的课程顺序是 [0,1,2,3] 。另一个正确的排序是 [0,2,1,3] 。 
//
// 示例 3： 
//输入：numCourses = 1, prerequisites = []
//输出：[0]

//leetcode submit region begin(Prohibit modification and deletion)

//type Graph struct {
//	nodes map[int]*GraphNode
//}
//
//type GraphNode struct {
//	val   int
//	In    int
//	nexts []*GraphNode
//}
//
//func BuildGraph(nodeCount int, edgeArr [][]int) *Graph {
//	nodes := make(map[int]*GraphNode, len(edgeArr))
//
//	for _, edgeDesc := range edgeArr {
//		var fromNode *GraphNode
//		var toNode *GraphNode
//		if node, ok := nodes[edgeDesc[1]]; !ok {
//			fromNode = &GraphNode{
//				val:   edgeDesc[1],
//				In:    0,
//				nexts: []*GraphNode{},
//			}
//		} else {
//			fromNode = node
//		}
//
//		if node, ok := nodes[edgeDesc[0]]; !ok {
//			toNode = &GraphNode{
//				val:   edgeDesc[0],
//				In:    0,
//				nexts: []*GraphNode{},
//			}
//		} else {
//			toNode = node
//		}
//		toNode.In++
//		toNode.val = edgeDesc[0]
//
//		fromNode.val = edgeDesc[1]
//		fromNode.nexts = append(fromNode.nexts, toNode)
//
//		nodes[fromNode.val] = fromNode
//		nodes[toNode.val] = toNode
//	}
//
//	for i := 0; i < nodeCount; i++ {
//		if _, ok := nodes[i]; !ok {
//			nodes[i] = &GraphNode{
//				val:   i,
//				In:    0,
//				nexts: []*GraphNode{},
//			}
//		}
//	}
//
//	return &Graph{
//		nodes: nodes,
//	}
//}
//
//func TopologicalSort(graph *Graph) []int {
//	nodes := graph.nodes
//
//	orderList := make([]int, 0, len(nodes))
//
//	for {
//		var zeroInNode *GraphNode
//		for _, node := range nodes {
//			if node.In == 0 {
//				zeroInNode = node
//				break
//			}
//		}
//
//		if zeroInNode == nil {
//			if len(nodes) != 0 {
//				return nil
//			}
//			return orderList
//		}
//
//		orderList = append(orderList, zeroInNode.val)
//		for _, nextNode := range zeroInNode.nexts {
//			nextNode.In--
//		}
//		delete(nodes, zeroInNode.val)
//	}
//}

func findOrder(numCourses int, prerequisites [][]int) []int {
	edgeMap := make(map[int][]int)
	inCountArr := make([]int, numCourses)
	res := make([]int, 0, numCourses)

	for _, edge := range prerequisites {
		if _, ok := edgeMap[edge[1]]; !ok {
			edgeMap[edge[1]] = []int{}
		}
		edgeMap[edge[1]] = append(edgeMap[edge[1]], edge[0])
		inCountArr[edge[0]]++
	}

	for i := 0; i < len(inCountArr); i++ {
		if inCountArr[i] == 0 {
			res = append(res, i)
			for _, nodeVal := range edgeMap[i] {
				inCountArr[nodeVal]--
			}
			inCountArr[i] = -1
			i = -1
		}
	}

	if len(res)!=numCourses {
		return []int{}
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	relation := [][]int{{1, 0}, {1, 2}, {0, 1}}
	fmt.Println(findOrder(3, relation))
}
