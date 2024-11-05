package main

func minMutation(start string, end string, bank []string) int {
	// 将基因库转换为集合
	bankSet := make(map[string]bool)
	for _, gene := range bank {
		bankSet[gene] = true
	}

	// 如果目标基因序列不在基因库中，直接返回 -1
	if !bankSet[end] {
		return -1
	}

	// 初始化队列和访问集合
	queue := []string{start}
	visited := make(map[string]bool)
	visited[start] = true

	// 定义基因字符集
	genes := []byte{'A', 'C', 'G', 'T'}
	steps := 0

	// 广度优先搜索
	for len(queue) > 0 {
		steps++
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			current := queue[0]
			queue = queue[1:]

			// 尝试将每个字符变为 'A'、'C'、'G' 和 'T'
			for j := 0; j < len(current); j++ {
				for _, gene := range genes {
					if current[j] == gene {
						continue
					}
					next := current[:j] + string(gene) + current[j+1:]
					if next == end {
						return steps
					}
					if bankSet[next] && !visited[next] {
						queue = append(queue, next)
						visited[next] = true
					}
				}
			}
		}
	}

	return -1
}
