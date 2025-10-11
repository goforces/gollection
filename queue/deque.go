package queue

// Deque is a double-ended queue implemented as a ring buffer.
// Zero value is ready to use.
type Deque[T any] struct {
	buf        []T
	head, tail int // head points to first element; tail points to slot after last
	size       int
}

// NewDeque creates an empty Deque.
func NewDeque[T any]() *Deque[T] { return &Deque[T]{} }

func (d *Deque[T]) growIfNeeded() {
	if len(d.buf) == 0 {
		d.buf = make([]T, 1)
		return
	}
	if d.size < len(d.buf) {
		return
	}
	newBuf := make([]T, d.size*2)
	for i := 0; i < d.size; i++ {
		newBuf[i] = d.buf[(d.head+i)%len(d.buf)]
	}
	d.buf = newBuf
	d.head = 0
	d.tail = d.size
}

// Len returns the number of elements in the deque.
func (d *Deque[T]) Len() int { return d.size }

// IsEmpty reports whether the deque has no elements.
func (d *Deque[T]) IsEmpty() bool { return d.size == 0 }

// Clear removes all elements.
func (d *Deque[T]) Clear() {
	d.buf = nil
	d.head, d.tail, d.size = 0, 0, 0
}

// PushFront adds an element at the front.
func (d *Deque[T]) PushFront(v T) {
	d.growIfNeeded()
	d.head = (d.head - 1 + len(d.buf)) % len(d.buf)
	d.buf[d.head] = v
	d.size++
}

// PushBack adds an element at the back.
func (d *Deque[T]) PushBack(v T) {
	d.growIfNeeded()
	d.buf[d.tail] = v
	d.tail = (d.tail + 1) % len(d.buf)
	d.size++
}

// PopFront removes and returns the front element.
func (d *Deque[T]) PopFront() (T, bool) {
	var zero T
	if d.size == 0 {
		return zero, false
	}
	v := d.buf[d.head]
	var zeroT T
	d.buf[d.head] = zeroT
	d.head = (d.head + 1) % len(d.buf)
	d.size--
	return v, true
}

// PopBack removes and returns the back element.
func (d *Deque[T]) PopBack() (T, bool) {
	var zero T
	if d.size == 0 {
		return zero, false
	}
	d.tail = (d.tail - 1 + len(d.buf)) % len(d.buf)
	v := d.buf[d.tail]
	var zeroT T
	d.buf[d.tail] = zeroT
	d.size--
	return v, true
}

// PeekFront returns the front element without removing it.
func (d *Deque[T]) PeekFront() (T, bool) {
	var zero T
	if d.size == 0 {
		return zero, false
	}
	return d.buf[d.head], true
}

// PeekBack returns the back element without removing it.
func (d *Deque[T]) PeekBack() (T, bool) {
	var zero T
	if d.size == 0 {
		return zero, false
	}
	idx := (d.tail - 1 + len(d.buf)) % len(d.buf)
	return d.buf[idx], true
}

// ToSlice returns elements from front to back.
func (d *Deque[T]) ToSlice() []T {
	out := make([]T, d.size)
	for i := 0; i < d.size; i++ {
		out[i] = d.buf[(d.head+i)%len(d.buf)]
	}
	return out
}
