// Package stack provides a generic LIFO stack.
//
// ⚠️  NOT THREAD-SAFE
// This implementation is not safe for concurrent access.
// Wrap with external synchronization (sync.Mutex) if needed.
package stack

// Stack is a generic LIFO stack backed by a slice.
// Zero value is ready to use.
type Stack[T any] struct {
	data []T
}

// New returns an empty Stack.
func New[T any]() *Stack[T] {
	return &Stack[T]{data: make([]T, 0)}
}

// FromSlice builds a stack with elements from the provided slice.
// The last element of the slice becomes the top of the stack.
func FromSlice[T any](items []T) *Stack[T] {
	dup := make([]T, len(items))
	copy(dup, items)
	return &Stack[T]{data: dup}
}

// Push adds an element to the top of the stack.
func (s *Stack[T]) Push(v T) {
	s.data = append(s.data, v)
}

// Pop removes and returns the top-most element.
// The boolean is false when the stack is empty.
func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.data) == 0 {
		return zero, false
	}
	idx := len(s.data) - 1
	val := s.data[idx]
	s.data = s.data[:idx]
	return val, true
}

// Peek returns the top-most element without removing it.
// The boolean is false when the stack is empty.
func (s *Stack[T]) Peek() (T, bool) {
	var zero T
	if len(s.data) == 0 {
		return zero, false
	}
	return s.data[len(s.data)-1], true
}

// Len returns the number of elements in the stack.
func (s *Stack[T]) Len() int { return len(s.data) }

// IsEmpty reports whether the stack is empty.
func (s *Stack[T]) IsEmpty() bool { return len(s.data) == 0 }

// Clear removes all elements from the stack.
func (s *Stack[T]) Clear() {
	s.data = s.data[:0]
}

// ToSlice returns a copy of the underlying slice in stack order
// (first element is the bottom, last element is the top).
func (s *Stack[T]) ToSlice() []T {
	out := make([]T, len(s.data))
	copy(out, s.data)
	return out
}
