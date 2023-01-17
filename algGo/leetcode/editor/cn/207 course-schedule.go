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

type GraphNode struct {
	inCount int
	nexts   []int
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	graphMap := buildGraph(prerequisites)
	return judgeCanFinish(graphMap)
}

func buildGraph(prerequisites [][]int) map[int]GraphNode {
	graph := make(map[int]GraphNode, 0)

	for _, prerequisite := range prerequisites {
		fromVal := prerequisite[1]
		toVal := prerequisite[0]

		if _, ok := graph[fromVal]; !ok {
			graph[fromVal] = GraphNode{nexts: []int{}}
		}

		if _, ok := graph[toVal]; !ok {
			graph[toVal] = GraphNode{nexts: []int{}}
		}

		toNode := graph[toVal]
		toNode.inCount++
		graph[toVal] = toNode

		fromNode := graph[fromVal]
		fromNode.nexts = append(fromNode.nexts, toVal)
		graph[fromVal] = fromNode
	}

	return graph
}

func judgeCanFinish(graphMap map[int]GraphNode) bool {

	flag := false
	for {
		for val, node := range graphMap {
			if node.inCount == 0 {
				flag = true
				for _, nextVal := range node.nexts {
					nextNode := graphMap[nextVal]
					nextNode.inCount--
					graphMap[nextVal] = nextNode
				}
				delete(graphMap, val)
				break
			}
		}
		if !flag {
			break
		} else {
			flag = false
		}
	}
	if len(graphMap) == 0 {
		return true
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	relation := [][]int{{1, 4}, {2, 4}, {3, 1}, {3, 2}}
	print(canFinish(5, relation))
}
