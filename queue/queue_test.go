package queue

import (
	"testing"
)

func TestNewQueue(t *testing.T) {
	q := NewQueue[int]()
	if q == nil {
		t.Fatal("NewQueue() returned nil")
	}
	if !q.IsEmpty() {
		t.Error("New queue should be empty")
	}
	if q.Len() != 0 {
		t.Errorf("New queue length = %d, want 0", q.Len())
	}
}

func TestQueueFromSlice(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{"empty slice", []int{}, []int{}},
		{"single element", []int{1}, []int{1}},
		{"multiple elements", []int{1, 2, 3}, []int{1, 2, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := QueueFromSlice(tt.input)
			got := q.ToSlice()
			
			if len(got) != len(tt.want) {
				t.Errorf("len = %d, want %d", len(got), len(tt.want))
				return
			}
			
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("element[%d] = %d, want %d", i, got[i], tt.want[i])
				}
			}
		})
	}
}

func TestEnqueueDequeue(t *testing.T) {
	q := NewQueue[int]()
	
	// Enqueue elements
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	
	if q.Len() != 3 {
		t.Errorf("after 3 enqueues, len = %d, want 3", q.Len())
	}
	
	// Dequeue elements (FIFO order)
	tests := []struct {
		wantVal int
		wantOk  bool
	}{
		{1, true},
		{2, true},
		{3, true},
		{0, false}, // empty
	}
	
	for i, tt := range tests {
		val, ok := q.Dequeue()
		if ok != tt.wantOk {
			t.Errorf("dequeue[%d] ok = %v, want %v", i, ok, tt.wantOk)
		}
		if ok && val != tt.wantVal {
			t.Errorf("dequeue[%d] val = %d, want %d", i, val, tt.wantVal)
		}
	}
}

func TestDequeueEmpty(t *testing.T) {
	q := NewQueue[string]()
	val, ok := q.Dequeue()
	if ok {
		t.Error("Dequeue on empty queue should return false")
	}
	if val != "" {
		t.Errorf("Dequeue on empty queue returned %q, want zero value", val)
	}
}

func TestQueuePeek(t *testing.T) {
	q := NewQueue[int]()
	
	// Peek on empty
	_, ok := q.Peek()
	if ok {
		t.Error("Peek on empty queue should return false")
	}
	
	// Peek with elements
	q.Enqueue(1)
	q.Enqueue(2)
	
	val, ok := q.Peek()
	if !ok {
		t.Error("Peek should return true on non-empty queue")
	}
	if val != 1 {
		t.Errorf("Peek = %d, want 1", val)
	}
	
	// Verify peek doesn't remove element
	if q.Len() != 2 {
		t.Error("Peek should not modify queue size")
	}
}

func TestQueueLen(t *testing.T) {
	q := NewQueue[int]()
	
	lengths := []int{0, 1, 2, 3, 2, 1, 0}
	ops := []func(){
		func() {},
		func() { q.Enqueue(1) },
		func() { q.Enqueue(2) },
		func() { q.Enqueue(3) },
		func() { q.Dequeue() },
		func() { q.Dequeue() },
		func() { q.Dequeue() },
	}
	
	for i, op := range ops {
		op()
		if q.Len() != lengths[i] {
			t.Errorf("after op[%d], len = %d, want %d", i, q.Len(), lengths[i])
		}
	}
}

func TestQueueIsEmpty(t *testing.T) {
	q := NewQueue[int]()
	
	if !q.IsEmpty() {
		t.Error("new queue should be empty")
	}
	
	q.Enqueue(1)
	if q.IsEmpty() {
		t.Error("queue with element should not be empty")
	}
	
	q.Dequeue()
	if !q.IsEmpty() {
		t.Error("queue after dequeuing all should be empty")
	}
}

func TestQueueClear(t *testing.T) {
	q := QueueFromSlice([]int{1, 2, 3, 4, 5})
	q.Clear()
	
	if !q.IsEmpty() {
		t.Error("Clear() should make queue empty")
	}
	if q.Len() != 0 {
		t.Errorf("after Clear(), len = %d, want 0", q.Len())
	}
	
	// Should be able to use after clear
	q.Enqueue(10)
	val, ok := q.Dequeue()
	if !ok || val != 10 {
		t.Error("queue should be usable after Clear()")
	}
}

func TestQueueToSlice(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	q := QueueFromSlice(input)
	
	got := q.ToSlice()
	if len(got) != len(input) {
		t.Fatalf("ToSlice len = %d, want %d", len(got), len(input))
	}
	
	for i := range got {
		if got[i] != input[i] {
			t.Errorf("element[%d] = %d, want %d", i, got[i], input[i])
		}
	}
}

func TestQueueClone(t *testing.T) {
	original := QueueFromSlice([]int{1, 2, 3})
	clone := original.Clone()
	
	// Verify contents match
	origSlice := original.ToSlice()
	cloneSlice := clone.ToSlice()
	
	if len(origSlice) != len(cloneSlice) {
		t.Fatalf("clone len = %d, want %d", len(cloneSlice), len(origSlice))
	}
	
	for i := range origSlice {
		if origSlice[i] != cloneSlice[i] {
			t.Errorf("element[%d]: original = %d, clone = %d", i, origSlice[i], cloneSlice[i])
		}
	}
	
	// Verify independence
	clone.Enqueue(999)
	if original.Len() == clone.Len() {
		t.Error("modifying clone should not affect original")
	}
}

func TestQueueZeroValue(t *testing.T) {
	var q Queue[int]
	
	// Should be usable without calling NewQueue
	if !q.IsEmpty() {
		t.Error("zero value queue should be empty")
	}
	
	q.Enqueue(1)
	val, ok := q.Dequeue()
	if !ok || val != 1 {
		t.Error("zero value queue should be usable")
	}
}

func TestQueueRingBufferWraparound(t *testing.T) {
	q := NewQueue[int]()
	
	// Fill and empty the queue multiple times to test ring buffer wrapping
	for iteration := 0; iteration < 3; iteration++ {
		for i := 0; i < 10; i++ {
			q.Enqueue(i)
		}
		
		for i := 0; i < 10; i++ {
			val, ok := q.Dequeue()
			if !ok {
				t.Fatalf("Dequeue failed at iteration %d, index %d", iteration, i)
			}
			if val != i {
				t.Errorf("iteration %d: Dequeue = %d, want %d", iteration, val, i)
			}
		}
		
		if !q.IsEmpty() {
			t.Errorf("iteration %d: queue should be empty", iteration)
		}
	}
}

func TestQueueGrowth(t *testing.T) {
	q := NewQueue[int]()
	n := 1000
	
	// Enqueue many elements to trigger growth
	for i := 0; i < n; i++ {
		q.Enqueue(i)
	}
	
	if q.Len() != n {
		t.Errorf("len = %d, want %d", q.Len(), n)
	}
	
	// Verify FIFO order
	for i := 0; i < n; i++ {
		val, ok := q.Dequeue()
		if !ok {
			t.Fatalf("Dequeue failed at index %d", i)
		}
		if val != i {
			t.Errorf("Dequeue = %d, want %d", val, i)
		}
	}
}

func TestQueueGenericTypes(t *testing.T) {
	t.Run("strings", func(t *testing.T) {
		q := NewQueue[string]()
		q.Enqueue("hello")
		q.Enqueue("world")
		val, ok := q.Dequeue()
		if !ok || val != "hello" {
			t.Errorf("got %q, want %q", val, "hello")
		}
	})
	
	t.Run("struct", func(t *testing.T) {
		type person struct {
			name string
			age  int
		}
		q := NewQueue[person]()
		p := person{"Alice", 30}
		q.Enqueue(p)
		val, ok := q.Dequeue()
		if !ok || val != p {
			t.Error("queue should work with structs")
		}
	})
}

func TestQueueLargeQueue(t *testing.T) {
	q := NewQueue[int]()
	n := 10000
	
	// Enqueue many elements
	for i := 0; i < n; i++ {
		q.Enqueue(i)
	}
	
	if q.Len() != n {
		t.Errorf("len = %d, want %d", q.Len(), n)
	}
	
	// Dequeue them all
	for i := 0; i < n; i++ {
		val, ok := q.Dequeue()
		if !ok {
			t.Fatalf("Dequeue failed at iteration %d", i)
		}
		if val != i {
			t.Errorf("Dequeue = %d, want %d", val, i)
		}
	}
	
	if !q.IsEmpty() {
		t.Error("queue should be empty after dequeuing all")
	}
}

// Benchmarks
func BenchmarkEnqueue(b *testing.B) {
	q := NewQueue[int]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}
}

func BenchmarkDequeue(b *testing.B) {
	q := NewQueue[int]()
	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Dequeue()
	}
}

func BenchmarkEnqueueDequeue(b *testing.B) {
	q := NewQueue[int]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
		q.Dequeue()
	}
}

