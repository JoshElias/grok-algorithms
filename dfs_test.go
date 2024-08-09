package main

import (
	"testing"
)

func TestDFS(t *testing.T) {
	// Test case 1: Node is found
	tree1 := map[int]Node{
		1: {value: 1, children: []Node{
			{value: 2, children: []Node{
				{value: 4, children: nil},
				{value: 5, children: nil},
			}},
			{value: 3, children: []Node{
				{value: 6, children: nil},
				{value: 7, children: nil},
			}},
		}},
	}
	node, found := DFS(tree1, 5)
	if !found || node.value != 5 {
		t.Errorf("Expected to find node with value 5, but got %v", node.value)
	}

	// Test case 2: Node is not found
	tree2 := map[int]Node{
		1: {value: 1, children: []Node{
			{value: 2, children: []Node{
				{value: 4, children: nil},
				{value: 5, children: nil},
			}},
			{value: 3, children: []Node{
				{value: 6, children: nil},
				{value: 7, children: nil},
			}},
		}},
	}
	_, found = DFS(tree2, 8)
	if found {
		t.Errorf("Expected not to find node with value 8, but found it")
	}

	// Test case 3: Empty tree
	tree3 := map[int]Node{}
	_, found = DFS(tree3, 1)
	if found {
		t.Errorf("Expected not to find node in an empty tree, but found it")
	}
}
