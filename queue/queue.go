package queue

// Queue is a generic FIFO queue implemented as a ring buffer.
// Zero value is ready to use.
type Queue[T any] struct {
	buf        []T
	head, tail int
	size       int
}

// NewQueue creates an empty Queue.
func NewQueue[T any]() *Queue[T] { return &Queue[T]{} }

// Len returns the number of elements in the queue.
func (q *Queue[T]) Len() int { return q.size }

// IsEmpty reports whether the queue has no elements.
func (q *Queue[T]) IsEmpty() bool { return q.size == 0 }

// Clear removes all elements from the queue.
func (q *Queue[T]) Clear() {
	q.buf = nil
	q.head, q.tail, q.size = 0, 0, 0
}

func (q *Queue[T]) growIfNeeded() {
	if len(q.buf) == 0 {
		q.buf = make([]T, 1)
		return
	}
	if q.size < len(q.buf) {
		return
	}
	newBuf := make([]T, q.size*2)
	// copy in order from head
	for i := 0; i < q.size; i++ {
		newBuf[i] = q.buf[(q.head+i)%len(q.buf)]
	}
	q.buf = newBuf
	q.head = 0
	q.tail = q.size
}

// Enqueue adds an element to the back of the queue.
func (q *Queue[T]) Enqueue(v T) {
	q.growIfNeeded()
	q.buf[q.tail] = v
	q.tail = (q.tail + 1) % len(q.buf)
	q.size++
}

// Dequeue removes and returns the element at the front of the queue.
// The boolean is false when the queue is empty.
func (q *Queue[T]) Dequeue() (T, bool) {
	var zero T
	if q.size == 0 {
		return zero, false
	}
	v := q.buf[q.head]
	var zeroT T
	q.buf[q.head] = zeroT // avoid memory leak for reference types
	q.head = (q.head + 1) % len(q.buf)
	q.size--
	return v, true
}

// Peek returns the front element without removing it.
func (q *Queue[T]) Peek() (T, bool) {
	var zero T
	if q.size == 0 {
		return zero, false
	}
	return q.buf[q.head], true
}

// ToSlice returns the elements in FIFO order.
func (q *Queue[T]) ToSlice() []T {
	out := make([]T, q.size)
	for i := 0; i < q.size; i++ {
		out[i] = q.buf[(q.head+i)%len(q.buf)]
	}
	return out
}
