package main

type Node struct {
	value    int
	children []Node
}

func BFS(tree map[int]Node, key int) (Node, bool) {
	// Create the queue
	queue := make([]Node, 0)

	// Add the first level of keys to search
	for _, v := range tree {
		queue = append(queue, v)
	}

	// Start iterating over each node
	for len(queue) > 0 {

		// Shift the first element
		node := queue[0]
		queue = queue[1:]

		if node.value == key {
			return node, true
		}

		for _, child := range node.children {
			queue = append(queue, child)
		}
	}

	return Node{}, false
}
