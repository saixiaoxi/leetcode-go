package main

func findOrder(numCourses int, prerequisites [][]int) []int {

	result := []int{}
	//构建图和入度数组
	graph := make(map[int][]int)
	indegree := make([]int, numCourses)

	for _, pair := range prerequisites {
		u := pair[1]
		v := pair[0]
		graph[u] = append(graph[u], v)
		indegree[v]++
	}

	//拓扑排序
	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		result = append(result, u)
		for _, v := range graph[u] {
			indegree[v]--
			if indegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
	if len(result) != numCourses {
		return nil
	}
	return result
}
