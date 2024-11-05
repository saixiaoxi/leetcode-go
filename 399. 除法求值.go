package main

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// 构建图
	graph := make(map[string]map[string]float64)
	for i, equation := range equations {
		a, b := equation[0], equation[1]
		if graph[a] == nil {
			graph[a] = make(map[string]float64)
		}
		if graph[b] == nil {
			graph[b] = make(map[string]float64)
		}
		graph[a][b] = values[i]
		graph[b][a] = 1 / values[i]
	}

	// 处理查询
	var dfs func(string, string, map[string]bool) float64
	dfs = func(start, end string, visited map[string]bool) float64 {
		if _, ok := graph[start]; !ok {
			return -1.0
		}
		if _, ok := graph[start][end]; ok {
			return graph[start][end]
		}
		visited[start] = true
		for neighbor, value := range graph[start] {
			if !visited[neighbor] {
				product := dfs(neighbor, end, visited)
				if product != -1.0 {
					return product * value
				}
			}
		}
		return -1.0
	}

	results := make([]float64, len(queries))
	for i, query := range queries {
		start, end := query[0], query[1]
		if start == end {
			if _, ok := graph[start]; ok {
				results[i] = 1.0
			} else {
				results[i] = -1.0
			}
		} else {
			results[i] = dfs(start, end, make(map[string]bool))
		}
	}

	return results
}
