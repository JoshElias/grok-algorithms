package main

func DFS(tree map[int]Node, key int) (Node, bool) {
	var nodeList []Node
	for _, v := range tree {
		nodeList = append(nodeList, v)
	}
	return DFSRecursive(nodeList, key)
}

func DFSRecursive(nodes []Node, key int) (Node, bool) {
	for _, node := range nodes {
		if node.value == key {
			return node, true
		}
		if len(node.children) > 0 {
			return DFSRecursive(node.children, key)
		}
	}
	return Node{}, false
}
