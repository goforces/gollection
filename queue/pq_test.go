package queue

import (
	"testing"
)

func TestNewPriorityQueue(t *testing.T) {
	less := func(a, b int) bool { return a < b }
	pq := NewPriorityQueue(less)
	
	if pq == nil {
		t.Fatal("NewPriorityQueue() returned nil")
	}
	if !pq.IsEmpty() {
		t.Error("New priority queue should be empty")
	}
	if pq.Len() != 0 {
		t.Errorf("New priority queue length = %d, want 0", pq.Len())
	}
}

func TestPriorityQueueFromSlice(t *testing.T) {
	less := func(a, b int) bool { return a < b }
	input := []int{5, 2, 8, 1, 9}
	pq := PriorityQueueFromSlice(input, less)
	
	if pq.Len() != len(input) {
		t.Errorf("len = %d, want %d", pq.Len(), len(input))
	}
	
	// Pop all and verify they come out in sorted order
	want := []int{1, 2, 5, 8, 9}
	for i, expected := range want {
		val, ok := pq.Pop()
		if !ok {
			t.Fatalf("Pop failed at index %d", i)
		}
		if val != expected {
			t.Errorf("Pop[%d] = %d, want %d", i, val, expected)
		}
	}
}

func TestMinHeap(t *testing.T) {
	// Min heap: smaller values have higher priority
	less := func(a, b int) bool { return a < b }
	pq := NewPriorityQueue(less)
	
	// Push elements in random order
	pq.Push(5)
	pq.Push(2)
	pq.Push(8)
	pq.Push(1)
	pq.Push(9)
	pq.Push(3)
	
	// Should pop in ascending order
	want := []int{1, 2, 3, 5, 8, 9}
	for i, expected := range want {
		val, ok := pq.Pop()
		if !ok {
			t.Fatalf("Pop failed at index %d", i)
		}
		if val != expected {
			t.Errorf("Pop[%d] = %d, want %d", i, val, expected)
		}
	}
}

func TestMaxHeap(t *testing.T) {
	// Max heap: larger values have higher priority
	less := func(a, b int) bool { return a > b }
	pq := NewPriorityQueue(less)
	
	// Push elements in random order
	pq.Push(5)
	pq.Push(2)
	pq.Push(8)
	pq.Push(1)
	pq.Push(9)
	pq.Push(3)
	
	// Should pop in descending order
	want := []int{9, 8, 5, 3, 2, 1}
	for i, expected := range want {
		val, ok := pq.Pop()
		if !ok {
			t.Fatalf("Pop failed at index %d", i)
		}
		if val != expected {
			t.Errorf("Pop[%d] = %d, want %d", i, val, expected)
		}
	}
}

func TestPopEmpty(t *testing.T) {
	less := func(a, b int) bool { return a < b }
	pq := NewPriorityQueue[int](less)
	
	val, ok := pq.Pop()
	if ok {
		t.Error("Pop on empty priority queue should return false")
	}
	if val != 0 {
		t.Errorf("Pop on empty priority queue returned %d, want zero value", val)
	}
}

func TestPriorityQueuePeek(t *testing.T) {
	less := func(a, b int) bool { return a < b }
	pq := NewPriorityQueue(less)
	
	// Peek on empty
	_, ok := pq.Peek()
	if ok {
		t.Error("Peek on empty priority queue should return false")
	}
	
	// Add elements
	pq.Push(5)
	pq.Push(2)
	pq.Push(8)
	
	// Peek should return min (2) without removing it
	val, ok := pq.Peek()
	if !ok {
		t.Error("Peek should return true on non-empty priority queue")
	}
	if val != 2 {
		t.Errorf("Peek = %d, want 2", val)
	}
	
	// Verify peek doesn't remove element
	if pq.Len() != 3 {
		t.Error("Peek should not modify priority queue size")
	}
	
	// Pop should return same value
	val, _ = pq.Pop()
	if val != 2 {
		t.Errorf("Pop after Peek = %d, want 2", val)
	}
}

func TestPriorityQueueLen(t *testing.T) {
	less := func(a, b int) bool { return a < b }
	pq := NewPriorityQueue(less)
	
	lengths := []int{0, 1, 2, 3, 2, 1, 0}
	ops := []func(){
		func() {},
		func() { pq.Push(1) },
		func() { pq.Push(2) },
		func() { pq.Push(3) },
		func() { pq.Pop() },
		func() { pq.Pop() },
		func() { pq.Pop() },
	}
	
	for i, op := range ops {
		op()
		if pq.Len() != lengths[i] {
			t.Errorf("after op[%d], len = %d, want %d", i, pq.Len(), lengths[i])
		}
	}
}

func TestPriorityQueueIsEmpty(t *testing.T) {
	less := func(a, b int) bool { return a < b }
	pq := NewPriorityQueue(less)
	
	if !pq.IsEmpty() {
		t.Error("new priority queue should be empty")
	}
	
	pq.Push(1)
	if pq.IsEmpty() {
		t.Error("priority queue with element should not be empty")
	}
	
	pq.Pop()
	if !pq.IsEmpty() {
		t.Error("priority queue after popping all should be empty")
	}
}

func TestPriorityQueueClear(t *testing.T) {
	less := func(a, b int) bool { return a < b }
	pq := PriorityQueueFromSlice([]int{1, 2, 3, 4, 5}, less)
	pq.Clear()
	
	if !pq.IsEmpty() {
		t.Error("Clear() should make priority queue empty")
	}
	if pq.Len() != 0 {
		t.Errorf("after Clear(), len = %d, want 0", pq.Len())
	}
	
	// Should be able to use after clear
	pq.Push(10)
	val, ok := pq.Pop()
	if !ok || val != 10 {
		t.Error("priority queue should be usable after Clear()")
	}
}

func TestPriorityQueueClone(t *testing.T) {
	less := func(a, b int) bool { return a < b }
	original := PriorityQueueFromSlice([]int{3, 1, 4, 1, 5}, less)
	clone := original.Clone()
	
	// Verify contents match
	if original.Len() != clone.Len() {
		t.Fatalf("clone len = %d, want %d", clone.Len(), original.Len())
	}
	
	// Pop from both and compare
	for original.Len() > 0 {
		origVal, origOk := original.Pop()
		cloneVal, cloneOk := clone.Pop()
		
		if origOk != cloneOk {
			t.Error("clone and original should have same ok values")
		}
		if origVal != cloneVal {
			t.Errorf("clone value = %d, original value = %d", cloneVal, origVal)
		}
	}
}

func TestCloneIndependence(t *testing.T) {
	less := func(a, b int) bool { return a < b }
	original := PriorityQueueFromSlice([]int{1, 2, 3}, less)
	clone := original.Clone()
	
	// Modify clone
	clone.Push(999)
	
	// Original should be unaffected
	if original.Len() == clone.Len() {
		t.Error("modifying clone should not affect original")
	}
}

func TestCustomComparator(t *testing.T) {
	type task struct {
		name     string
		priority int
	}
	
	// Higher priority value = higher priority
	less := func(a, b task) bool { return a.priority > b.priority }
	pq := NewPriorityQueue(less)
	
	pq.Push(task{"low", 1})
	pq.Push(task{"high", 10})
	pq.Push(task{"medium", 5})
	
	// Should pop in priority order
	val, _ := pq.Pop()
	if val.priority != 10 {
		t.Errorf("first pop priority = %d, want 10", val.priority)
	}
	
	val, _ = pq.Pop()
	if val.priority != 5 {
		t.Errorf("second pop priority = %d, want 5", val.priority)
	}
	
	val, _ = pq.Pop()
	if val.priority != 1 {
		t.Errorf("third pop priority = %d, want 1", val.priority)
	}
}

func TestStringPriorityQueue(t *testing.T) {
	// Lexicographic order
	less := func(a, b string) bool { return a < b }
	pq := NewPriorityQueue(less)
	
	pq.Push("zebra")
	pq.Push("apple")
	pq.Push("mango")
	pq.Push("banana")
	
	want := []string{"apple", "banana", "mango", "zebra"}
	for i, expected := range want {
		val, ok := pq.Pop()
		if !ok {
			t.Fatalf("Pop failed at index %d", i)
		}
		if val != expected {
			t.Errorf("Pop[%d] = %q, want %q", i, val, expected)
		}
	}
}

func TestDuplicateElements(t *testing.T) {
	less := func(a, b int) bool { return a < b }
	pq := NewPriorityQueue(less)
	
	pq.Push(5)
	pq.Push(2)
	pq.Push(5)
	pq.Push(2)
	pq.Push(5)
	
	want := []int{2, 2, 5, 5, 5}
	for i, expected := range want {
		val, ok := pq.Pop()
		if !ok {
			t.Fatalf("Pop failed at index %d", i)
		}
		if val != expected {
			t.Errorf("Pop[%d] = %d, want %d", i, val, expected)
		}
	}
}

func TestLargePriorityQueue(t *testing.T) {
	less := func(a, b int) bool { return a < b }
	pq := NewPriorityQueue(less)
	n := 10000
	
	// Push elements in reverse order
	for i := n - 1; i >= 0; i-- {
		pq.Push(i)
	}
	
	if pq.Len() != n {
		t.Errorf("len = %d, want %d", pq.Len(), n)
	}
	
	// Pop them all - should come out in sorted order
	for i := 0; i < n; i++ {
		val, ok := pq.Pop()
		if !ok {
			t.Fatalf("Pop failed at iteration %d", i)
		}
		if val != i {
			t.Errorf("Pop = %d, want %d", val, i)
		}
	}
	
	if !pq.IsEmpty() {
		t.Error("priority queue should be empty after popping all")
	}
}

func TestHeapProperty(t *testing.T) {
	less := func(a, b int) bool { return a < b }
	pq := NewPriorityQueue(less)
	
	// Push random values
	values := []int{15, 3, 17, 5, 25, 7, 19, 1, 9, 11}
	for _, v := range values {
		pq.Push(v)
	}
	
	// Pop all values - they should come out in sorted order
	prev := -1
	for pq.Len() > 0 {
		val, ok := pq.Pop()
		if !ok {
			t.Fatal("Pop failed")
		}
		if val < prev {
			t.Errorf("heap property violated: %d came before %d", prev, val)
		}
		prev = val
	}
}

func TestMixedPushPop(t *testing.T) {
	less := func(a, b int) bool { return a < b }
	pq := NewPriorityQueue(less)
	
	// Mix push and pop operations
	pq.Push(5)
	pq.Push(2)
	val, _ := pq.Pop() // Should be 2
	if val != 2 {
		t.Errorf("Pop = %d, want 2", val)
	}
	
	pq.Push(1)
	pq.Push(8)
	val, _ = pq.Pop() // Should be 1
	if val != 1 {
		t.Errorf("Pop = %d, want 1", val)
	}
	
	pq.Push(3)
	
	// Remaining: 5, 8, 3 -> should pop 3, 5, 8
	want := []int{3, 5, 8}
	for i, expected := range want {
		val, ok := pq.Pop()
		if !ok {
			t.Fatalf("Pop failed at index %d", i)
		}
		if val != expected {
			t.Errorf("Pop[%d] = %d, want %d", i, val, expected)
		}
	}
}

// Benchmarks
func BenchmarkPush(b *testing.B) {
	less := func(a, b int) bool { return a < b }
	pq := NewPriorityQueue(less)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pq.Push(i)
	}
}

func BenchmarkPop(b *testing.B) {
	less := func(a, b int) bool { return a < b }
	pq := NewPriorityQueue(less)
	for i := 0; i < b.N; i++ {
		pq.Push(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pq.Pop()
	}
}

func BenchmarkPushPop(b *testing.B) {
	less := func(a, b int) bool { return a < b }
	pq := NewPriorityQueue(less)
	// Pre-fill with some elements
	for i := 0; i < 100; i++ {
		pq.Push(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pq.Push(i)
		pq.Pop()
	}
}

