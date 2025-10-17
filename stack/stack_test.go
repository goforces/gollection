package stack

import (
	"testing"
)

func TestNew(t *testing.T) {
	s := New[int]()
	if s == nil {
		t.Fatal("New() returned nil")
	}
	if !s.IsEmpty() {
		t.Error("New stack should be empty")
	}
	if s.Len() != 0 {
		t.Errorf("New stack length = %d, want 0", s.Len())
	}
}

func TestFromSlice(t *testing.T) {
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
			s := FromSlice(tt.input)
			got := s.ToSlice()
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

func TestPushPop(t *testing.T) {
	s := New[int]()
	
	// Test push
	s.Push(1)
	s.Push(2)
	s.Push(3)
	
	if s.Len() != 3 {
		t.Errorf("after 3 pushes, len = %d, want 3", s.Len())
	}
	
	// Test pop
	tests := []struct {
		wantVal int
		wantOk  bool
	}{
		{3, true},
		{2, true},
		{1, true},
		{0, false}, // empty
	}
	
	for i, tt := range tests {
		val, ok := s.Pop()
		if ok != tt.wantOk {
			t.Errorf("pop[%d] ok = %v, want %v", i, ok, tt.wantOk)
		}
		if ok && val != tt.wantVal {
			t.Errorf("pop[%d] val = %d, want %d", i, val, tt.wantVal)
		}
	}
}

func TestPopEmpty(t *testing.T) {
	s := New[string]()
	val, ok := s.Pop()
	if ok {
		t.Error("Pop on empty stack should return false")
	}
	if val != "" {
		t.Errorf("Pop on empty stack returned %q, want zero value", val)
	}
}

func TestPeek(t *testing.T) {
	s := New[int]()
	
	// Peek on empty
	_, ok := s.Peek()
	if ok {
		t.Error("Peek on empty stack should return false")
	}
	
	// Peek with elements
	s.Push(1)
	s.Push(2)
	
	val, ok := s.Peek()
	if !ok {
		t.Error("Peek should return true on non-empty stack")
	}
	if val != 2 {
		t.Errorf("Peek = %d, want 2", val)
	}
	
	// Verify peek doesn't remove element
	if s.Len() != 2 {
		t.Error("Peek should not modify stack size")
	}
}

func TestLen(t *testing.T) {
	s := New[int]()
	
	lengths := []int{0, 1, 2, 3, 2, 1, 0}
	ops := []func(){
		func() {},
		func() { s.Push(1) },
		func() { s.Push(2) },
		func() { s.Push(3) },
		func() { s.Pop() },
		func() { s.Pop() },
		func() { s.Pop() },
	}
	
	for i, op := range ops {
		op()
		if s.Len() != lengths[i] {
			t.Errorf("after op[%d], len = %d, want %d", i, s.Len(), lengths[i])
		}
	}
}

func TestIsEmpty(t *testing.T) {
	s := New[int]()
	
	if !s.IsEmpty() {
		t.Error("new stack should be empty")
	}
	
	s.Push(1)
	if s.IsEmpty() {
		t.Error("stack with element should not be empty")
	}
	
	s.Pop()
	if !s.IsEmpty() {
		t.Error("stack after popping all should be empty")
	}
}

func TestClear(t *testing.T) {
	s := FromSlice([]int{1, 2, 3, 4, 5})
	s.Clear()
	
	if !s.IsEmpty() {
		t.Error("Clear() should make stack empty")
	}
	if s.Len() != 0 {
		t.Errorf("after Clear(), len = %d, want 0", s.Len())
	}
	
	// Should be able to use after clear
	s.Push(10)
	val, ok := s.Pop()
	if !ok || val != 10 {
		t.Error("stack should be usable after Clear()")
	}
}

func TestToSlice(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	s := FromSlice(input)
	
	got := s.ToSlice()
	if len(got) != len(input) {
		t.Fatalf("ToSlice len = %d, want %d", len(got), len(input))
	}
	
	for i := range got {
		if got[i] != input[i] {
			t.Errorf("element[%d] = %d, want %d", i, got[i], input[i])
		}
	}
	
	// Verify it's a copy
	got[0] = 999
	if s.ToSlice()[0] == 999 {
		t.Error("ToSlice should return a copy, not a reference")
	}
}

func TestClone(t *testing.T) {
	original := FromSlice([]int{1, 2, 3})
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
	clone.Push(999)
	if original.Len() == clone.Len() {
		t.Error("modifying clone should not affect original")
	}
}

func TestZeroValue(t *testing.T) {
	var s Stack[int]
	
	// Should be usable without calling New
	if !s.IsEmpty() {
		t.Error("zero value stack should be empty")
	}
	
	s.Push(1)
	val, ok := s.Pop()
	if !ok || val != 1 {
		t.Error("zero value stack should be usable")
	}
}

func TestGenericTypes(t *testing.T) {
	t.Run("strings", func(t *testing.T) {
		s := New[string]()
		s.Push("hello")
		s.Push("world")
		val, ok := s.Pop()
		if !ok || val != "world" {
			t.Errorf("got %q, want %q", val, "world")
		}
	})
	
	t.Run("struct", func(t *testing.T) {
		type person struct {
			name string
			age  int
		}
		s := New[person]()
		p := person{"Alice", 30}
		s.Push(p)
		val, ok := s.Pop()
		if !ok || val != p {
			t.Error("stack should work with structs")
		}
	})
	
	t.Run("pointers", func(t *testing.T) {
		s := New[*int]()
		x := 42
		s.Push(&x)
		val, ok := s.Pop()
		if !ok || val != &x {
			t.Error("stack should work with pointers")
		}
	})
}

func TestLargeStack(t *testing.T) {
	s := New[int]()
	n := 10000
	
	// Push many elements
	for i := 0; i < n; i++ {
		s.Push(i)
	}
	
	if s.Len() != n {
		t.Errorf("len = %d, want %d", s.Len(), n)
	}
	
	// Pop them all
	for i := n - 1; i >= 0; i-- {
		val, ok := s.Pop()
		if !ok {
			t.Fatalf("Pop failed at iteration %d", i)
		}
		if val != i {
			t.Errorf("Pop = %d, want %d", val, i)
		}
	}
	
	if !s.IsEmpty() {
		t.Error("stack should be empty after popping all")
	}
}

// Benchmarks
func BenchmarkPush(b *testing.B) {
	s := New[int]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
}

func BenchmarkPop(b *testing.B) {
	s := New[int]()
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Pop()
	}
}

func BenchmarkPushPop(b *testing.B) {
	s := New[int]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Push(i)
		s.Pop()
	}
}

