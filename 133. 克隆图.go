package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */
type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	// 哈希表记录已克隆的节点
	visited := make(map[*Node]*Node)

	// 深度优先搜索（DFS）函数
	var dfs func(*Node) *Node
	dfs = func(n *Node) *Node {
		if n == nil {
			return nil
		}
		// 如果该节点已经被克隆，直接返回克隆的节点
		if clonedNode, found := visited[n]; found {
			return clonedNode
		}
		// 克隆节点
		clone := &Node{Val: n.Val}
		visited[n] = clone
		// 克隆邻居节点
		for _, neighbor := range n.Neighbors {
			clone.Neighbors = append(clone.Neighbors, dfs(neighbor))
		}
		return clone
	}

	return dfs(node)
}
