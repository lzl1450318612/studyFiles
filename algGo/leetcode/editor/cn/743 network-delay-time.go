package main

import (
	"fmt"
)

//有 n 个网络节点，标记为 1 到 n。
// 给你一个列表 times，表示信号经过 有向 边的传递时间。 times[i] = (ui, vi, wi)，其中 ui 是源节点，vi 是目标节点，
//wi 是一个信号从源节点传递到目标节点的时间。 
// 现在，从某个节点 K 发出一个信号。需要多久才能使所有节点都收到信号？如果不能使所有节点收到信号，返回 -1 。
//
// 示例 1： 
//输入：times = [[2,1,1],[2,3,1],[3,4,1]], n = 4, k = 2
//输出：2
//
// 示例 2： 
//输入：times = [[1,2,1]], n = 2, k = 1
//输出：1
//
// 示例 3： 
//输入：times = [[1,2,1]], n = 2, k = 2
//输出：-1

//leetcode submit region begin(Prohibit modification and deletion)

func networkDelayTime(times [][]int, n int, k int) int {
	queue := make([]int, 0, n)
	queueMap := make(map[int]struct{}, n)

	graph := buildNetworkGraph(times, n+1)

	queue = append(queue, k)
	queueMap[k] = struct{}{}
	distanceMap := make(map[int]int, n)
	distanceMap[k] = 0

	for len(queue) != 0 {
		curNode := queue[0]
		queue = queue[1:]

		if curNode == 5 {
			fmt.Println("as")
		}

		for i := 0; i < len(graph[curNode]); i++ {
			weight := graph[curNode][i]
			if weight != 0 {
				if distance, ok := distanceMap[i]; !ok {
					distanceMap[i] = distanceMap[curNode] + weight
					if _, ok1 := queueMap[i]; !ok1 {
						queue = append(queue, i)
						queueMap[i] = struct{}{}
					}
				} else if weight+distanceMap[curNode] < distance {
					distanceMap[i] = weight + distanceMap[curNode]
					if _, ok1 := queueMap[i]; !ok1 {
						queue = append(queue, i)
						queueMap[i] = struct{}{}
					}
				}
			}
		}
	}

	if len(distanceMap) != n {
		return -1
	}

	maxDistance := 0
	for _, distance := range distanceMap {
		if distance > maxDistance {
			maxDistance = distance
		}
	}
	return maxDistance
}

// 构造邻接矩阵
func buildNetworkGraph(graphInfos [][]int, n int) [][]int {
	res := make([][]int, 0)

	for i := 0; i <= len(graphInfos)+1; i++ {
		res = append(res, make([]int, n))
	}

	for _, graphInfo := range graphInfos {
		if res[graphInfo[0]] == nil {
			res[graphInfo[0]] = []int{}
		}
		res[graphInfo[0]][graphInfo[1]] = graphInfo[2]
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(networkDelayTime([][]int{{1, 2, 15}, {1, 3, 2}, {1, 5, 1}, {3, 5, 3}, {5, 4, 4}, {4, 2, 5}, {2, 4, 6}}, 5, 1))
}
