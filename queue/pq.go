package queue

import "container/heap"

// PriorityQueue is a heap-backed priority queue where lower priority values
// come out first if less(a,b) returns true when a has higher priority than b.
type PriorityQueue[T any] struct {
	h *genericHeap[T]
}

// NewPriorityQueue creates a priority queue using the provided ordering.
// The function less must return true when a should come before b.
func NewPriorityQueue[T any](less func(a, b T) bool) *PriorityQueue[T] {
	gh := &genericHeap[T]{less: less}
	heap.Init(gh)
	return &PriorityQueue[T]{h: gh}
}

// Push inserts an element into the queue.
func (pq *PriorityQueue[T]) Push(v T) {
	heap.Push(pq.h, v)
}

// Pop removes and returns the top-priority element.
// The boolean is false when the queue is empty.
func (pq *PriorityQueue[T]) Pop() (T, bool) {
	var zero T
	if pq.h.Len() == 0 {
		return zero, false
	}
	v := heap.Pop(pq.h).(T)
	return v, true
}

// Peek returns the top-priority element without removing it.
func (pq *PriorityQueue[T]) Peek() (T, bool) {
	var zero T
	if pq.h.Len() == 0 {
		return zero, false
	}
	return pq.h.data[0], true
}

// Len returns the number of elements.
func (pq *PriorityQueue[T]) Len() int { return pq.h.Len() }

// IsEmpty reports whether the queue is empty.
func (pq *PriorityQueue[T]) IsEmpty() bool { return pq.h.Len() == 0 }

// Clear removes all elements.
func (pq *PriorityQueue[T]) Clear() {
	pq.h.data = nil
}

// PriorityQueueFromSlice creates a PriorityQueue from a slice using the provided ordering.
// The function less must return true when a should come before b.
func PriorityQueueFromSlice[T any](items []T, less func(a, b T) bool) *PriorityQueue[T] {
	pq := NewPriorityQueue[T](less)
	for _, item := range items {
		pq.Push(item)
	}
	return pq
}

// Clone returns a shallow copy of the priority queue.
func (pq *PriorityQueue[T]) Clone() *PriorityQueue[T] {
	clone := &PriorityQueue[T]{
		h: &genericHeap[T]{
			data: make([]T, len(pq.h.data)),
			less: pq.h.less,
		},
	}
	copy(clone.h.data, pq.h.data)
	return clone
}

// internal generic heap wrapper
type genericHeap[T any] struct {
	data []T
	less func(a, b T) bool
}

func (h *genericHeap[T]) Len() int           { return len(h.data) }
func (h *genericHeap[T]) Less(i, j int) bool { return h.less(h.data[i], h.data[j]) }
func (h *genericHeap[T]) Swap(i, j int)      { h.data[i], h.data[j] = h.data[j], h.data[i] }
func (h *genericHeap[T]) Push(x any)         { h.data = append(h.data, x.(T)) }
func (h *genericHeap[T]) Pop() any {
	n := len(h.data)
	v := h.data[n-1]
	h.data = h.data[:n-1]
	return v
}
