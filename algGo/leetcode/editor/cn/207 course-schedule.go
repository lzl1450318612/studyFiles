package main

// 你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。
// 在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi] ，表
//示如果要学习课程 ai 则 必须 先学习课程 bi 。 
// 例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
// 请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。
//
// 示例 1： 
//输入：numCourses = 2, prerequisites = [[1,0]]
//输出：true
//解释：总共有 2 门课程。学习课程 1 之前，你需要完成课程 0 。这是可能的。 
//
// 示例 2： 
//输入：numCourses = 2, prerequisites = [[1,0],[0,1]]
//输出：false
//解释：总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0 ；并且学习课程 0 之前，你还应先完成课程 1 。这是不可能的。 

//leetcode submit region begin(Prohibit modification and deletion)

type graph struct {
	nodes map[int]*graphNode
	edges []*graphEdge
}

type graphNode struct {
	val   int
	In    int
	Out   int
	nexts []*graphNode
	edges []*graphEdge
}

type graphEdge struct {
	weight int
	from   *graphNode
	to     *graphNode
}

func buildGraph(edgeArr [][]int) *graph {
	nodes := make(map[int]*graphNode, len(edgeArr))
	edges := make([]*graphEdge, len(edgeArr))

	for _, edgeDesc := range edgeArr {
		var fromNode *graphNode
		var toNode *graphNode
		if node, ok := nodes[edgeDesc[1]]; !ok {
			fromNode = &graphNode{
				val:   edgeDesc[1],
				In:    0,
				Out:   0,
				nexts: []*graphNode{},
				edges: []*graphEdge{},
			}
		} else {
			fromNode = node
		}

		if node, ok := nodes[edgeDesc[0]]; !ok {
			toNode = &graphNode{
				val:   edgeDesc[0],
				In:    0,
				Out:   0,
				nexts: []*graphNode{},
				edges: []*graphEdge{},
			}
		} else {
			toNode = node
		}
		toNode.In++
		toNode.val = edgeDesc[0]
		toNode.edges = append(toNode.edges, &graphEdge{
			from: fromNode,
			to:   toNode,
		})

		fromNode.Out++
		fromNode.val = edgeDesc[1]
		fromNode.nexts = append(fromNode.nexts, toNode)
		fromNode.edges = append(fromNode.edges, &graphEdge{
			from: fromNode,
			to:   toNode,
		})

		nodes[fromNode.val] = fromNode
		nodes[toNode.val] = toNode

		edges = append(edges, &graphEdge{
			from: fromNode,
			to:   toNode,
		})
	}

	return &graph{
		nodes: nodes,
		edges: edges,
	}
}

func topologicalSort(graph *graph) []*graphNode {
	nodes := graph.nodes

	orderList := make([]*graphNode, 0, len(nodes))

	for {
		var zeroInNode *graphNode
		for _, node := range nodes {
			if node.In == 0 {
				zeroInNode = node
				break
			}
		}

		if zeroInNode == nil {
			if len(nodes) != 0 {
				return nil
			}
			return orderList
		}

		orderList = append(orderList, zeroInNode)
		for _, nextNode := range zeroInNode.nexts {
			nextNode.In--
		}
		delete(nodes, zeroInNode.val)
	}
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	if len(prerequisites) == 0 && numCourses != 0 {
		return true
	}

	graph := buildGraph(prerequisites)
	orderList := topologicalSort(graph)
	if orderList == nil || len(orderList) == 0 {
		return false
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	relation := [][]int{}
	print(canFinish(1, relation))
}
