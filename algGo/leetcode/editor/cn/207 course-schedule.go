package main

// 你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。
//
// 在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi] ，表
//示如果要学习课程 ai 则 必须 先学习课程 bi 。 
//
// 例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
// 
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
//
// 提示： 
// 1 <= numCourses <= 10⁵
// 0 <= prerequisites.length <= 5000 
// prerequisites[i].length == 2 
// 0 <= ai, bi < numCourses 
// prerequisites[i] 中的所有课程对 互不相同 
// 
// Related Topics 深度优先搜索 广度优先搜索 图 拓扑排序 👍 1140 👎 0

//leetcode submit region begin(Prohibit modification and deletion)

func canFinish(numCourses int, prerequisites [][]int) bool {
	if numCourses < 1 {
		return false
	}
	if numCourses == 1 {
		return true
	}

	if len(prerequisites) < 1 {
		return false
	}

	visitedMap := make(map[int]int, 0)
	relationMap := make(map[int]int, 0)
	resultMap := make(map[int]bool, 0)

	for _, relation := range prerequisites {
		relationMap[relation[0]] = relation[1]
	}

	for k := range relationMap {
		res := search(k, relationMap, visitedMap, resultMap)
		if !res {
			return res
		}
	}
	return true
}

func search(head int, relationMap map[int]int, visitedMap map[int]int, resultMap map[int]bool) bool {
	if visitedMap[head] == 1 {
		return false
	}

	if _, ok := relationMap[head]; !ok {
		return true
	}

	var result bool
	visitedMap[head] = 1
	if resultMap[head] {
		result = true
	} else {
		result = search(relationMap[head], relationMap, visitedMap, resultMap)
		resultMap[head] = result
	}
	visitedMap[head] = 0
	return result
}

//leetcode submit region end(Prohibit modification and deletion)

func init() {
	Register(207, func() {
		relation := [][]int{{0, 1}, {1, 2}}
		print(canFinish(2, relation))
	})
}
