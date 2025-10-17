package queue

import (
	"testing"
)

func TestNewDeque(t *testing.T) {
	d := NewDeque[int]()
	if d == nil {
		t.Fatal("NewDeque() returned nil")
	}
	if !d.IsEmpty() {
		t.Error("New deque should be empty")
	}
	if d.Len() != 0 {
		t.Errorf("New deque length = %d, want 0", d.Len())
	}
}

func TestDequeFromSlice(t *testing.T) {
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
			d := DequeFromSlice(tt.input)
			got := d.ToSlice()
			
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

func TestPushFrontPopFront(t *testing.T) {
	d := NewDeque[int]()
	
	d.PushFront(1)
	d.PushFront(2)
	d.PushFront(3)
	
	if d.Len() != 3 {
		t.Errorf("after 3 pushes, len = %d, want 3", d.Len())
	}
	
	// PopFront should return in reverse order (3, 2, 1)
	tests := []struct {
		wantVal int
		wantOk  bool
	}{
		{3, true},
		{2, true},
		{1, true},
		{0, false},
	}
	
	for i, tt := range tests {
		val, ok := d.PopFront()
		if ok != tt.wantOk {
			t.Errorf("PopFront[%d] ok = %v, want %v", i, ok, tt.wantOk)
		}
		if ok && val != tt.wantVal {
			t.Errorf("PopFront[%d] val = %d, want %d", i, val, tt.wantVal)
		}
	}
}

func TestPushBackPopBack(t *testing.T) {
	d := NewDeque[int]()
	
	d.PushBack(1)
	d.PushBack(2)
	d.PushBack(3)
	
	if d.Len() != 3 {
		t.Errorf("after 3 pushes, len = %d, want 3", d.Len())
	}
	
	// PopBack should return in reverse order (3, 2, 1)
	tests := []struct {
		wantVal int
		wantOk  bool
	}{
		{3, true},
		{2, true},
		{1, true},
		{0, false},
	}
	
	for i, tt := range tests {
		val, ok := d.PopBack()
		if ok != tt.wantOk {
			t.Errorf("PopBack[%d] ok = %v, want %v", i, ok, tt.wantOk)
		}
		if ok && val != tt.wantVal {
			t.Errorf("PopBack[%d] val = %d, want %d", i, val, tt.wantVal)
		}
	}
}

func TestMixedOperations(t *testing.T) {
	d := NewDeque[int]()
	
	// Build: [3, 2, 1, 4, 5]
	d.PushBack(1)    // [1]
	d.PushFront(2)   // [2, 1]
	d.PushFront(3)   // [3, 2, 1]
	d.PushBack(4)    // [3, 2, 1, 4]
	d.PushBack(5)    // [3, 2, 1, 4, 5]
	
	want := []int{3, 2, 1, 4, 5}
	got := d.ToSlice()
	
	if len(got) != len(want) {
		t.Fatalf("len = %d, want %d", len(got), len(want))
	}
	
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("element[%d] = %d, want %d", i, got[i], want[i])
		}
	}
}

func TestPeekFront(t *testing.T) {
	d := NewDeque[int]()
	
	// Peek on empty
	_, ok := d.PeekFront()
	if ok {
		t.Error("PeekFront on empty deque should return false")
	}
	
	// Peek with elements
	d.PushBack(1)
	d.PushBack(2)
	
	val, ok := d.PeekFront()
	if !ok {
		t.Error("PeekFront should return true on non-empty deque")
	}
	if val != 1 {
		t.Errorf("PeekFront = %d, want 1", val)
	}
	
	// Verify peek doesn't remove element
	if d.Len() != 2 {
		t.Error("PeekFront should not modify deque size")
	}
}

func TestPeekBack(t *testing.T) {
	d := NewDeque[int]()
	
	// Peek on empty
	_, ok := d.PeekBack()
	if ok {
		t.Error("PeekBack on empty deque should return false")
	}
	
	// Peek with elements
	d.PushBack(1)
	d.PushBack(2)
	
	val, ok := d.PeekBack()
	if !ok {
		t.Error("PeekBack should return true on non-empty deque")
	}
	if val != 2 {
		t.Errorf("PeekBack = %d, want 2", val)
	}
	
	// Verify peek doesn't remove element
	if d.Len() != 2 {
		t.Error("PeekBack should not modify deque size")
	}
}

func TestDequeLen(t *testing.T) {
	d := NewDeque[int]()
	
	lengths := []int{0, 1, 2, 3, 4, 3, 2}
	ops := []func(){
		func() {},
		func() { d.PushBack(1) },
		func() { d.PushBack(2) },
		func() { d.PushFront(3) },
		func() { d.PushBack(4) },
		func() { d.PopFront() },
		func() { d.PopBack() },
	}
	
	for i, op := range ops {
		op()
		if d.Len() != lengths[i] {
			t.Errorf("after op[%d], len = %d, want %d", i, d.Len(), lengths[i])
		}
	}
}

func TestDequeIsEmpty(t *testing.T) {
	d := NewDeque[int]()
	
	if !d.IsEmpty() {
		t.Error("new deque should be empty")
	}
	
	d.PushBack(1)
	if d.IsEmpty() {
		t.Error("deque with element should not be empty")
	}
	
	d.PopBack()
	if !d.IsEmpty() {
		t.Error("deque after popping all should be empty")
	}
}

func TestDequeClear(t *testing.T) {
	d := DequeFromSlice([]int{1, 2, 3, 4, 5})
	d.Clear()
	
	if !d.IsEmpty() {
		t.Error("Clear() should make deque empty")
	}
	if d.Len() != 0 {
		t.Errorf("after Clear(), len = %d, want 0", d.Len())
	}
	
	// Should be able to use after clear
	d.PushBack(10)
	val, ok := d.PopFront()
	if !ok || val != 10 {
		t.Error("deque should be usable after Clear()")
	}
}

func TestDequeToSlice(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	d := DequeFromSlice(input)
	
	got := d.ToSlice()
	if len(got) != len(input) {
		t.Fatalf("ToSlice len = %d, want %d", len(got), len(input))
	}
	
	for i := range got {
		if got[i] != input[i] {
			t.Errorf("element[%d] = %d, want %d", i, got[i], input[i])
		}
	}
}

func TestDequeClone(t *testing.T) {
	original := DequeFromSlice([]int{1, 2, 3})
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
	clone.PushBack(999)
	if original.Len() == clone.Len() {
		t.Error("modifying clone should not affect original")
	}
}

func TestDequeZeroValue(t *testing.T) {
	var d Deque[int]
	
	// Should be usable without calling NewDeque
	if !d.IsEmpty() {
		t.Error("zero value deque should be empty")
	}
	
	d.PushBack(1)
	val, ok := d.PopFront()
	if !ok || val != 1 {
		t.Error("zero value deque should be usable")
	}
}

func TestDequeRingBufferWraparound(t *testing.T) {
	d := NewDeque[int]()
	
	// Alternate PushFront and PopFront to exercise wrapping
	for i := 0; i < 100; i++ {
		d.PushFront(i)
		if i%2 == 0 {
			d.PopBack()
		}
	}
	
	// Just verify it didn't crash and has reasonable size
	if d.Len() < 0 || d.Len() > 100 {
		t.Errorf("unexpected length: %d", d.Len())
	}
}

func TestDequeGrowth(t *testing.T) {
	d := NewDeque[int]()
	n := 1000
	
	// Push many elements from front
	for i := 0; i < n; i++ {
		d.PushFront(i)
	}
	
	if d.Len() != n {
		t.Errorf("len = %d, want %d", d.Len(), n)
	}
	
	// Verify order (should be n-1, n-2, ..., 1, 0)
	for i := n - 1; i >= 0; i-- {
		val, ok := d.PopFront()
		if !ok {
			t.Fatalf("PopFront failed at index %d", i)
		}
		if val != i {
			t.Errorf("PopFront = %d, want %d", val, i)
		}
	}
}

func TestAsStack(t *testing.T) {
	// Use deque as a stack (PushBack/PopBack)
	d := NewDeque[int]()
	
	d.PushBack(1)
	d.PushBack(2)
	d.PushBack(3)
	
	val, _ := d.PopBack()
	if val != 3 {
		t.Errorf("PopBack = %d, want 3", val)
	}
	val, _ = d.PopBack()
	if val != 2 {
		t.Errorf("PopBack = %d, want 2", val)
	}
	val, _ = d.PopBack()
	if val != 1 {
		t.Errorf("PopBack = %d, want 1", val)
	}
}

func TestAsQueue(t *testing.T) {
	// Use deque as a queue (PushBack/PopFront)
	d := NewDeque[int]()
	
	d.PushBack(1)
	d.PushBack(2)
	d.PushBack(3)
	
	val, _ := d.PopFront()
	if val != 1 {
		t.Errorf("PopFront = %d, want 1", val)
	}
	val, _ = d.PopFront()
	if val != 2 {
		t.Errorf("PopFront = %d, want 2", val)
	}
	val, _ = d.PopFront()
	if val != 3 {
		t.Errorf("PopFront = %d, want 3", val)
	}
}

func TestDequeLargeDeque(t *testing.T) {
	d := NewDeque[int]()
	n := 10000
	
	// Push many elements
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			d.PushBack(i)
		} else {
			d.PushFront(i)
		}
	}
	
	if d.Len() != n {
		t.Errorf("len = %d, want %d", d.Len(), n)
	}
	
	// Pop them all
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			_, ok := d.PopFront()
			if !ok {
				t.Fatalf("PopFront failed at iteration %d", i)
			}
		} else {
			_, ok := d.PopBack()
			if !ok {
				t.Fatalf("PopBack failed at iteration %d", i)
			}
		}
	}
	
	if !d.IsEmpty() {
		t.Error("deque should be empty after popping all")
	}
}

// Benchmarks
func BenchmarkPushFront(b *testing.B) {
	d := NewDeque[int]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.PushFront(i)
	}
}

func BenchmarkPushBack(b *testing.B) {
	d := NewDeque[int]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.PushBack(i)
	}
}

func BenchmarkPopFront(b *testing.B) {
	d := NewDeque[int]()
	for i := 0; i < b.N; i++ {
		d.PushBack(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.PopFront()
	}
}

func BenchmarkPopBack(b *testing.B) {
	d := NewDeque[int]()
	for i := 0; i < b.N; i++ {
		d.PushBack(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.PopBack()
	}
}

