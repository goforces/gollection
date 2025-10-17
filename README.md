# Go Collections Framework (GCF)

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/badge/Go-1.18+-00ADD8?logo=go)](https://go.dev/)
[![Test Coverage](https://img.shields.io/badge/coverage-96%2B%25-brightgreen)](https://github.com/goforces/gollection)

A Go-native Collections Framework inspired by Java/Kotlin `java.util.*`, built with generics (Go 1.18+).

## ✨ Features
- 🎯 **Type-Safe Generics** - Leverage Go 1.18+ generics for compile-time type safety
- 📦 **Comprehensive Data Structures** - Stack, Queue, Deque, Priority Queue, Set, Dictionary, Tree, Graph
- ⚡ **High Performance** - Optimized implementations with benchmarks
- 🧪 **Well Tested** - 96%+ test coverage with 200+ test cases
- 🔧 **Functional Utilities** - Map, Filter, Reduce for idiomatic Go
- 📖 **Zero Dependencies** - Only uses Go standard library

## 📦 Packages

### Data Structures
- **`stack`** - LIFO stack with `Stack[T]`
- **`queue`** - `Queue[T]` (FIFO), `Deque[T]` (double-ended), `PriorityQueue[T]` (heap-based)
- **`set`** - `HashSet[T]` with set operations (Union, Intersection, Difference)
- **`collections`** - `Dictionary[K,V]` map wrapper with helpful methods
- **`tree`** - `BinaryTree[K,V]` for ordered key-value pairs
- **`graph`** - `Graph[T]` adjacency-list implementation (directed/undirected, weighted)

### Algorithms & Utilities
- **`algorithms`** - `BinarySearch`, `QuickSort` with custom comparators
- **`collections`** - Functional utilities: `Map`, `Filter`, `Reduce`

## 🎯 Test Coverage
All packages are thoroughly tested with comprehensive test suites:

| Package | Coverage | Test Cases |
|---------|----------|------------|
| algorithms | 100.0% | Binary search, QuickSort, edge cases |
| stack | 100.0% | All operations, generic types, stress tests |
| graph | 100.0% | Directed/undirected, weighted edges, removal |
| queue | 98.6% | Queue, Deque, Priority Queue with benchmarks |
| tree | 98.0% | BST operations, deletion, custom comparators |
| collections | 96.4% | Dictionary, Map/Filter/Reduce |
| set | 96.3% | Set operations, generic types, large sets |

## Install

```bash
go get github.com/goforces/gollection
```

## 🚀 Quick Start

### Stack
```go
import "github.com/goforces/gollection/stack"

s := stack.New[int]()
s.Push(1)
s.Push(2)
s.Push(3)
val, ok := s.Pop()  // 3, true
fmt.Println(val)    // 3
```

### Queue & Deque
```go
import "github.com/goforces/gollection/queue"

// FIFO Queue
q := queue.NewQueue[string]()
q.Enqueue("first")
q.Enqueue("second")
val, _ := q.Dequeue()  // "first"

// Double-ended Queue
dq := queue.NewDeque[int]()
dq.PushFront(1)
dq.PushBack(2)
front, _ := dq.PopFront()  // 1
back, _ := dq.PopBack()    // 2

// Priority Queue (min-heap)
pq := queue.NewPriorityQueue(func(a, b int) bool { return a < b })
pq.Push(5)
pq.Push(2)
pq.Push(8)
val, _ := pq.Pop()  // 2
```

### Set Operations
```go
import "github.com/goforces/gollection/set"

s1 := set.FromSlice([]int{1, 2, 3, 4})
s2 := set.FromSlice([]int{3, 4, 5, 6})

union := set.Union(s1, s2)         // {1, 2, 3, 4, 5, 6}
intersection := set.Intersection(s1, s2)  // {3, 4}
difference := set.Difference(s1, s2)      // {1, 2}

s1.Add(10)
s1.Remove(1)
fmt.Println(s1.Contains(10))  // true
```

### Dictionary
```go
import "github.com/goforces/gollection/collections"

dict := collections.New[string, int]()
dict.Put("age", 30)
dict.Put("score", 95)

age, ok := dict.Get("age")  // 30, true
dict.Delete("score")

keys := dict.Keys()      // ["age"]
values := dict.Values()  // [30]
```

### Tree (BST)
```go
import "github.com/goforces/gollection/tree"

cmp := func(a, b int) int {
    if a < b { return -1 }
    if a > b { return 1 }
    return 0
}

bst := tree.New[int, string](cmp)
bst.Put(5, "five")
bst.Put(3, "three")
bst.Put(7, "seven")

val, ok := bst.Get(3)     // "three", true
bst.Delete(5)
exists := bst.Contains(5) // false
```

### Graph
```go
import "github.com/goforces/gollection/graph"

// Directed graph
g := graph.New[string](true)
g.AddEdge("A", "B", 1.0)
g.AddEdge("B", "C", 2.0)
g.AddEdge("A", "C", 4.0)

neighbors := g.Neighbors("A")  // map[B:1.0 C:4.0]
vertices := g.Vertices()       // [A, B, C]

// Undirected graph
ug := graph.New[int](false)
ug.AddEdge(1, 2, 5.0)  // Creates edges 1->2 and 2->1
```

### Functional Utilities
```go
import "github.com/goforces/gollection/collections"

// Map
nums := []int{1, 2, 3, 4, 5}
doubled := collections.Map(nums, func(x int) int { return x * 2 })
// [2, 4, 6, 8, 10]

// Filter
evens := collections.Filter(nums, func(x int) bool { return x%2 == 0 })
// [2, 4]

// Reduce
sum := collections.Reduce(nums, 0, func(acc, x int) int { return acc + x })
// 15

// Chaining
result := collections.Reduce(
    collections.Filter(
        collections.Map(nums, func(x int) int { return x * 2 }),
        func(x int) bool { return x > 5 },
    ),
    0,
    func(acc, x int) int { return acc + x },
)
// 30 (doubled, filtered > 5, summed)
```

### Algorithms
```go
import "github.com/goforces/gollection/algorithms"

// Binary Search
arr := []int{1, 3, 5, 7, 9, 11, 13}
cmp := func(a, b int) int {
    if a < b { return -1 }
    if a > b { return 1 }
    return 0
}
idx := algorithms.BinarySearch(arr, 7, cmp)  // 3

// QuickSort
data := []int{5, 2, 8, 1, 9, 3}
algorithms.QuickSort(data, func(a, b int) bool { return a < b })
// data is now [1, 2, 3, 5, 8, 9]
```

## 🧪 Testing

Run tests for all packages:
```bash
go test ./...
```

Run tests with coverage:
```bash
go test ./... -cover
```

Run benchmarks:
```bash
go test ./... -bench=. -benchmem
```

## ⚠️ Thread Safety

**Note**: All data structures in this library are **NOT thread-safe** by design. This keeps the implementations simple and performant. If you need concurrent access, wrap the structures with `sync.Mutex` or `sync.RWMutex`:

```go
type SafeStack[T any] struct {
    mu sync.Mutex
    s  *stack.Stack[T]
}

func (ss *SafeStack[T]) Push(v T) {
    ss.mu.Lock()
    defer ss.mu.Unlock()
    ss.s.Push(v)
}
```

## 📋 Roadmap

### ✅ Completed
- ✅ Phase 1: Stack, Queue/Deque, Set, PriorityQueue, Dictionary
- ✅ Phase 2: Graph, Tree (BST)
- ✅ Phase 3: Algorithms (BinarySearch, QuickSort)
- ✅ Comprehensive test coverage (96%+)
- ✅ Benchmarks for all data structures
- ✅ Functional utilities (Map, Filter, Reduce)

### 🚧 Planned
- [ ] Graph algorithms (BFS, DFS, Dijkstra, Kruskal, Prim)
- [ ] Union-Find (Disjoint Set)
- [ ] Balanced trees (AVL, Red-Black)
- [ ] Trie data structure
- [ ] More sorting algorithms (MergeSort, HeapSort)
- [ ] Iterator patterns
- [ ] JSON serialization support

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to:
- Add tests for new features
- Update documentation
- Follow existing code style
- Ensure all tests pass

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

Inspired by:
- Java Collections Framework (`java.util.*`)
- Kotlin Standard Library
- Go's philosophy of simplicity and clarity
