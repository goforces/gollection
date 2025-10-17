# Go Collections Framework (GCF)

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A Go-native Collections Framework inspired by Java/Kotlin `java.util.*`, built with generics (Go 1.18+).

## Vision
- Fill the missing gap in Go: a standardized collections library.
- Provide data structures (stacks, queues, sets, graphs, trees, etc.).
- Provide algorithms (sorting, searching, graph traversal).
- Be idiomatic: follow Go’s simplicity, avoid heavy OOP.
- Be modern: leverage Go’s generics, benchmarks, and functional utilities (Map, Filter, Reduce).

## Packages
- `stack`: `Stack[T]`
- `queue`: `Queue[T]`, `Deque[T]`, `PriorityQueue[T]`
- `set`: `HashSet[T]`
- `collections`: `Dictionary[K,V]`, functional utilities `Map`, `Filter`, `Reduce`
- `algorithms`: sorting and searching utilities
- `graph`: adjacency-list graphs (directed/undirected)
- `tree`: basic BST for ordered keys

## Install

```bash
go get github.com/goforces/gollection
```

## Quick Start

```go
package main

import (
	"fmt"
	"github.com/goforces/gollection/stack"
	"github.com/goforces/gollection/queue"
)

func main() {
	s := stack.New[int]()
	s.Push(1)
	s.Push(2)
	v, _ := s.Pop()
	fmt.Println(v) // 2

	q := queue.NewQueue[int]()
	q.Enqueue(10)
	q.Enqueue(20)
	x, _ := q.Dequeue()
	fmt.Println(x) // 10
}
```

## Roadmap
- Phase 1: Stack, Queue/Deque, Set, PriorityQueue, Dictionary
- Phase 2: Graph, Tree, UnionFind
- Phase 3: Algorithms (BFS/DFS/Dijkstra), Sorting, Searching

Contributions and ideas welcome!
