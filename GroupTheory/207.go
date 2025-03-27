package GroupTheory

/*
   题目描述：
	Middle 课程表
	你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。
	在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi] ，表示如果要学习课程 ai 则 必须 先学习课程  bi 。
	例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
	请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。
*/

/*
   解题思路：
	1.DFS判断有无环的存在
	2.拓扑排序看总共能上多少节课
	这里实现第二种算法，用到的数据结构为入度数组和邻接表
*/

func canFinish(numCourses int, prerequisites [][]int) bool {
	// 1. 构建入度表 & 邻接表
	inDegree := make([]int, numCourses)  // 入度数组
	adjList := make([][]int, numCourses) // 邻接表

	for _, pre := range prerequisites {
		a, b := pre[0], pre[1]
		inDegree[a]++                      // 课程 a 的入度 +1
		adjList[b] = append(adjList[b], a) // 课程 b 指向 a（b 是 a 的前置）
	}

	// 2. 将所有入度为 0 的课程入队
	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	// 3. BFS 拓扑排序
	count := 0 // 记录已学习的课程数
	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]
		count++ // 记录已完成的课程数

		// 遍历 course 的后续课程
		for _, next := range adjList[course] {
			inDegree[next]-- // 依赖减少
			if inDegree[next] == 0 {
				queue = append(queue, next) // 如果后续课程入度为 0，则可学习
			}
		}
	}

	// 如果所有课程都能学完，则返回 true，否则返回 false
	return count == numCourses
}
