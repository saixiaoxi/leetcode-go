package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	// 构建图和入度数组
	graph := make(map[int][]int)
	indegree := make([]int, numCourses)

	for _, pair := range prerequisites {
		u, v := pair[1], pair[0]
		graph[u] = append(graph[u], v)
		indegree[v]++
	}

	// 使用队列进行拓扑排序
	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	count := 0
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		count++

		for _, v := range graph[u] {
			indegree[v]--
			if indegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	return count == numCourses
}
