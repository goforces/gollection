package set

import (
	"sort"
	"testing"
)

func TestNew(t *testing.T) {
	s := New[int]()
	if s == nil {
		t.Fatal("New() returned nil")
	}
	if !s.IsEmpty() {
		t.Error("New set should be empty")
	}
	if s.Len() != 0 {
		t.Errorf("New set length = %d, want 0", s.Len())
	}
}

func TestFromSlice(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		wantLen  int
		wantVals []int
	}{
		{"empty slice", []int{}, 0, []int{}},
		{"single element", []int{1}, 1, []int{1}},
		{"unique elements", []int{1, 2, 3}, 3, []int{1, 2, 3}},
		{"duplicate elements", []int{1, 2, 2, 3, 1}, 3, []int{1, 2, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := FromSlice(tt.input)
			if s.Len() != tt.wantLen {
				t.Errorf("len = %d, want %d", s.Len(), tt.wantLen)
			}
			
			for _, val := range tt.wantVals {
				if !s.Contains(val) {
					t.Errorf("set missing value %d", val)
				}
			}
		})
	}
}

func TestAdd(t *testing.T) {
	s := New[int]()
	
	// Add new element
	added := s.Add(1)
	if !added {
		t.Error("adding new element should return true")
	}
	if !s.Contains(1) {
		t.Error("set should contain added element")
	}
	
	// Add duplicate
	added = s.Add(1)
	if added {
		t.Error("adding duplicate should return false")
	}
	if s.Len() != 1 {
		t.Errorf("len = %d, want 1", s.Len())
	}
}

func TestRemove(t *testing.T) {
	s := FromSlice([]int{1, 2, 3})
	
	// Remove existing element
	removed := s.Remove(2)
	if !removed {
		t.Error("removing existing element should return true")
	}
	if s.Contains(2) {
		t.Error("set should not contain removed element")
	}
	
	// Remove non-existent element
	removed = s.Remove(999)
	if removed {
		t.Error("removing non-existent element should return false")
	}
	
	// Remove from empty
	s.Clear()
	removed = s.Remove(1)
	if removed {
		t.Error("removing from empty set should return false")
	}
}

func TestContains(t *testing.T) {
	s := FromSlice([]int{1, 2, 3})
	
	tests := []struct {
		val  int
		want bool
	}{
		{1, true},
		{2, true},
		{3, true},
		{4, false},
		{0, false},
	}
	
	for _, tt := range tests {
		got := s.Contains(tt.val)
		if got != tt.want {
			t.Errorf("Contains(%d) = %v, want %v", tt.val, got, tt.want)
		}
	}
}

func TestLen(t *testing.T) {
	s := New[int]()
	
	lengths := []int{0, 1, 2, 3, 2, 1, 0}
	ops := []func(){
		func() {},
		func() { s.Add(1) },
		func() { s.Add(2) },
		func() { s.Add(3) },
		func() { s.Remove(2) },
		func() { s.Remove(3) },
		func() { s.Remove(1) },
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
		t.Error("new set should be empty")
	}
	
	s.Add(1)
	if s.IsEmpty() {
		t.Error("set with element should not be empty")
	}
	
	s.Remove(1)
	if !s.IsEmpty() {
		t.Error("set after removing all should be empty")
	}
}

func TestClear(t *testing.T) {
	s := FromSlice([]int{1, 2, 3, 4, 5})
	s.Clear()
	
	if !s.IsEmpty() {
		t.Error("Clear() should make set empty")
	}
	if s.Len() != 0 {
		t.Errorf("after Clear(), len = %d, want 0", s.Len())
	}
	
	// Should be able to use after clear
	s.Add(10)
	if !s.Contains(10) {
		t.Error("set should be usable after Clear()")
	}
}

func TestToSlice(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	s := FromSlice(input)
	
	got := s.ToSlice()
	if len(got) != len(input) {
		t.Fatalf("ToSlice len = %d, want %d", len(got), len(input))
	}
	
	// Sort to compare (order is unspecified)
	sort.Ints(got)
	sort.Ints(input)
	
	for i := range got {
		if got[i] != input[i] {
			t.Errorf("element[%d] = %d, want %d", i, got[i], input[i])
		}
	}
}

func TestClone(t *testing.T) {
	original := FromSlice([]int{1, 2, 3})
	clone := original.Clone()
	
	// Verify contents match
	if original.Len() != clone.Len() {
		t.Fatalf("clone len = %d, want %d", clone.Len(), original.Len())
	}
	
	for _, val := range original.ToSlice() {
		if !clone.Contains(val) {
			t.Errorf("clone missing value %d", val)
		}
	}
	
	// Verify independence
	clone.Add(999)
	if original.Contains(999) {
		t.Error("modifying clone should not affect original")
	}
}

func TestUnion(t *testing.T) {
	tests := []struct {
		name string
		a    []int
		b    []int
		want []int
	}{
		{"both empty", []int{}, []int{}, []int{}},
		{"a empty", []int{}, []int{1, 2}, []int{1, 2}},
		{"b empty", []int{1, 2}, []int{}, []int{1, 2}},
		{"no overlap", []int{1, 2}, []int{3, 4}, []int{1, 2, 3, 4}},
		{"some overlap", []int{1, 2, 3}, []int{2, 3, 4}, []int{1, 2, 3, 4}},
		{"complete overlap", []int{1, 2, 3}, []int{1, 2, 3}, []int{1, 2, 3}},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := FromSlice(tt.a)
			b := FromSlice(tt.b)
			result := Union(a, b)
			
			if result.Len() != len(tt.want) {
				t.Errorf("Union len = %d, want %d", result.Len(), len(tt.want))
			}
			
			for _, val := range tt.want {
				if !result.Contains(val) {
					t.Errorf("Union missing value %d", val)
				}
			}
		})
	}
}

func TestIntersection(t *testing.T) {
	tests := []struct {
		name string
		a    []int
		b    []int
		want []int
	}{
		{"both empty", []int{}, []int{}, []int{}},
		{"a empty", []int{}, []int{1, 2}, []int{}},
		{"b empty", []int{1, 2}, []int{}, []int{}},
		{"no overlap", []int{1, 2}, []int{3, 4}, []int{}},
		{"some overlap", []int{1, 2, 3}, []int{2, 3, 4}, []int{2, 3}},
		{"complete overlap", []int{1, 2, 3}, []int{1, 2, 3}, []int{1, 2, 3}},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := FromSlice(tt.a)
			b := FromSlice(tt.b)
			result := Intersection(a, b)
			
			if result.Len() != len(tt.want) {
				t.Errorf("Intersection len = %d, want %d", result.Len(), len(tt.want))
			}
			
			for _, val := range tt.want {
				if !result.Contains(val) {
					t.Errorf("Intersection missing value %d", val)
				}
			}
		})
	}
}

func TestDifference(t *testing.T) {
	tests := []struct {
		name string
		a    []int
		b    []int
		want []int
	}{
		{"both empty", []int{}, []int{}, []int{}},
		{"a empty", []int{}, []int{1, 2}, []int{}},
		{"b empty", []int{1, 2}, []int{}, []int{1, 2}},
		{"no overlap", []int{1, 2}, []int{3, 4}, []int{1, 2}},
		{"some overlap", []int{1, 2, 3}, []int{2, 3, 4}, []int{1}},
		{"complete overlap", []int{1, 2, 3}, []int{1, 2, 3}, []int{}},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := FromSlice(tt.a)
			b := FromSlice(tt.b)
			result := Difference(a, b)
			
			if result.Len() != len(tt.want) {
				t.Errorf("Difference len = %d, want %d", result.Len(), len(tt.want))
			}
			
			for _, val := range tt.want {
				if !result.Contains(val) {
					t.Errorf("Difference missing value %d", val)
				}
			}
		})
	}
}

func TestIsSubset(t *testing.T) {
	tests := []struct {
		name string
		a    []int
		b    []int
		want bool
	}{
		{"both empty", []int{}, []int{}, true},
		{"a empty", []int{}, []int{1, 2}, true},
		{"b empty, a not", []int{1}, []int{}, false},
		{"a subset of b", []int{1, 2}, []int{1, 2, 3}, true},
		{"equal sets", []int{1, 2, 3}, []int{1, 2, 3}, true},
		{"not subset", []int{1, 2, 4}, []int{1, 2, 3}, false},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := FromSlice(tt.a)
			b := FromSlice(tt.b)
			got := IsSubset(a, b)
			
			if got != tt.want {
				t.Errorf("IsSubset(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestForEach(t *testing.T) {
	s := FromSlice([]int{1, 2, 3, 4, 5})
	
	sum := 0
	s.ForEach(func(v int) {
		sum += v
	})
	
	want := 15
	if sum != want {
		t.Errorf("sum = %d, want %d", sum, want)
	}
}

func TestZeroValue(t *testing.T) {
	var s HashSet[int]
	
	// Should be usable without calling New
	if !s.IsEmpty() {
		t.Error("zero value set should be empty")
	}
	
	s.Add(1)
	if !s.Contains(1) {
		t.Error("zero value set should be usable")
	}
}

func TestGenericTypes(t *testing.T) {
	t.Run("strings", func(t *testing.T) {
		s := New[string]()
		s.Add("hello")
		s.Add("world")
		if !s.Contains("hello") {
			t.Error("set should contain 'hello'")
		}
	})
	
	t.Run("runes", func(t *testing.T) {
		s := New[rune]()
		s.Add('a')
		s.Add('b')
		if !s.Contains('a') {
			t.Error("set should contain 'a'")
		}
	})
}

func TestLargeSet(t *testing.T) {
	s := New[int]()
	n := 10000
	
	// Add many elements
	for i := 0; i < n; i++ {
		s.Add(i)
	}
	
	if s.Len() != n {
		t.Errorf("len = %d, want %d", s.Len(), n)
	}
	
	// Verify all present
	for i := 0; i < n; i++ {
		if !s.Contains(i) {
			t.Errorf("set missing value %d", i)
		}
	}
	
	// Remove all
	for i := 0; i < n; i++ {
		s.Remove(i)
	}
	
	if !s.IsEmpty() {
		t.Error("set should be empty after removing all")
	}
}

// Benchmarks
func BenchmarkAdd(b *testing.B) {
	s := New[int]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Add(i)
	}
}

func BenchmarkContains(b *testing.B) {
	s := New[int]()
	for i := 0; i < 1000; i++ {
		s.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Contains(i % 1000)
	}
}

func BenchmarkRemove(b *testing.B) {
	s := New[int]()
	for i := 0; i < b.N; i++ {
		s.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Remove(i)
	}
}

func BenchmarkUnion(b *testing.B) {
	a := New[int]()
	b_ := New[int]()
	for i := 0; i < 1000; i++ {
		a.Add(i)
		b_.Add(i + 500)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Union(a, b_)
	}
}

