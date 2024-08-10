package main

import (
	"fmt"
)

type LRU struct {
	values  []Node
	maxSize int
}

type Node struct {
	key   string
	value string
}

// The node read becomes the newest node
func (cache *LRU) Read(key string) (string, bool) {
	for idx, node := range cache.values {
		if node.key == key {
			freshNode := cache.values[idx]
			cache.values = append(cache.values[:idx], cache.values[idx+1:]...)
			cache.values = append(cache.values, freshNode)
			return node.value, true
		}
	}
	return "", false
}

func (cache *LRU) Count() int {
	return len(cache.values)
}

// Check if we have space (max 100)
// Delete if necessary
// If overriding an existing value, it's the newest
func (cache *LRU) Write(key string, value string) {
	idx := -1
	for i, n := range cache.values {
		if key == n.key {
			idx = i
		}
	}

	node := Node{
		key:   key,
		value: value,
	}

	if idx != -1 {
		cache.values[idx] = node
		return
	}

	if len(cache.values) == cache.maxSize {
		cache.values = cache.values[1:]
	}

	cache.values = append(cache.values, node)
}

func (cache *LRU) Delete(key string) bool {
	for i, node := range cache.values {
		if node.key == key {
			cache.values = append(cache.values[:i], cache.values[i+1:]...)
			return true
		}
	}
	return false
}

func (cache *LRU) Clear() {
	cache.values = make([]Node, 0)
}

func (cache *LRU) Print() {
	for _, n := range cache.values {
		fmt.Println("Key: ", n.key, " Value: ", n.value)
	}
}

// func main() {
// 	lru := LRU{
// 		values:  make([]Node, 0, 100),
// 		maxSize: 3,
// 	}
//
// 	lru.Write("1", "1")
// 	v, _ := lru.Read("1")
// 	fmt.Println("retrieved value: ", v)
//
// 	lru.Write("2", "2")
// 	lru.Write("3", "3")
// 	lru.Write("4", "4")
//
// 	lru.Print()
// }
