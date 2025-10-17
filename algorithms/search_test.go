package algorithms

import (
	"testing"
)

func intCmp(a, b int) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

func stringCmp(a, b string) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

func TestBinarySearchFound(t *testing.T) {
	tests := []struct {
		name   string
		arr    []int
		target int
		want   int
	}{
		{"single element", []int{5}, 5, 0},
		{"first element", []int{1, 2, 3, 4, 5}, 1, 0},
		{"last element", []int{1, 2, 3, 4, 5}, 5, 4},
		{"middle element", []int{1, 2, 3, 4, 5}, 3, 2},
		{"even length array", []int{1, 2, 3, 4, 5, 6}, 4, 3},
		{"larger array", []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}, 13, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BinarySearch(tt.arr, tt.target, intCmp)
			if got != tt.want {
				t.Errorf("BinarySearch(%v, %d) = %d, want %d", tt.arr, tt.target, got, tt.want)
			}
		})
	}
}

func TestBinarySearchNotFound(t *testing.T) {
	tests := []struct {
		name   string
		arr    []int
		target int
	}{
		{"empty array", []int{}, 5},
		{"smaller than first", []int{1, 2, 3, 4, 5}, 0},
		{"larger than last", []int{1, 2, 3, 4, 5}, 6},
		{"between elements", []int{1, 3, 5, 7, 9}, 4},
		{"single element mismatch", []int{5}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BinarySearch(tt.arr, tt.target, intCmp)
			if got != -1 {
				t.Errorf("BinarySearch(%v, %d) = %d, want -1", tt.arr, tt.target, got)
			}
		})
	}
}

func TestBinarySearchStrings(t *testing.T) {
	arr := []string{"apple", "banana", "cherry", "date", "elderberry"}
	
	tests := []struct {
		target string
		want   int
	}{
		{"apple", 0},
		{"banana", 1},
		{"cherry", 2},
		{"date", 3},
		{"elderberry", 4},
		{"apricot", -1},
		{"fig", -1},
		{"", -1},
	}

	for _, tt := range tests {
		got := BinarySearch(arr, tt.target, stringCmp)
		if got != tt.want {
			t.Errorf("BinarySearch(%v, %q) = %d, want %d", arr, tt.target, got, tt.want)
		}
	}
}

func TestBinarySearchDuplicates(t *testing.T) {
	// When there are duplicates, binary search returns *an* index (not necessarily the first)
	arr := []int{1, 2, 2, 2, 3, 4, 5}
	
	idx := BinarySearch(arr, 2, intCmp)
	if idx < 1 || idx > 3 {
		t.Errorf("BinarySearch with duplicates should return a valid index, got %d", idx)
	}
	if arr[idx] != 2 {
		t.Errorf("BinarySearch returned index %d with value %d, want 2", idx, arr[idx])
	}
}

func TestBinarySearchAllSame(t *testing.T) {
	arr := []int{5, 5, 5, 5, 5}
	
	idx := BinarySearch(arr, 5, intCmp)
	if idx < 0 || idx >= len(arr) {
		t.Errorf("BinarySearch should find element, got index %d", idx)
	}
	if arr[idx] != 5 {
		t.Errorf("BinarySearch returned wrong value")
	}
	
	notFound := BinarySearch(arr, 3, intCmp)
	if notFound != -1 {
		t.Errorf("BinarySearch should return -1 for missing element, got %d", notFound)
	}
}

func TestBinarySearchNegativeNumbers(t *testing.T) {
	arr := []int{-10, -5, -3, 0, 2, 5, 10}
	
	tests := []struct {
		target int
		want   int
	}{
		{-10, 0},
		{-5, 1},
		{0, 3},
		{10, 6},
		{-7, -1},
		{3, -1},
	}

	for _, tt := range tests {
		got := BinarySearch(arr, tt.target, intCmp)
		if got != tt.want {
			t.Errorf("BinarySearch(%v, %d) = %d, want %d", arr, tt.target, got, tt.want)
		}
	}
}

func TestBinarySearchSingleElement(t *testing.T) {
	// Found
	idx := BinarySearch([]int{42}, 42, intCmp)
	if idx != 0 {
		t.Errorf("BinarySearch([42], 42) = %d, want 0", idx)
	}
	
	// Not found
	idx = BinarySearch([]int{42}, 1, intCmp)
	if idx != -1 {
		t.Errorf("BinarySearch([42], 1) = %d, want -1", idx)
	}
}

func TestBinarySearchTwoElements(t *testing.T) {
	arr := []int{1, 2}
	
	tests := []struct {
		target int
		want   int
	}{
		{1, 0},
		{2, 1},
		{0, -1},
		{3, -1},
	}

	for _, tt := range tests {
		got := BinarySearch(arr, tt.target, intCmp)
		if got != tt.want {
			t.Errorf("BinarySearch(%v, %d) = %d, want %d", arr, tt.target, got, tt.want)
		}
	}
}

func TestBinarySearchCustomType(t *testing.T) {
	type person struct {
		name string
		age  int
	}
	
	cmp := func(a, b person) int {
		if a.age < b.age {
			return -1
		}
		if a.age > b.age {
			return 1
		}
		return 0
	}
	
	arr := []person{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 35},
		{"Dave", 40},
	}
	
	target := person{"", 35}
	idx := BinarySearch(arr, target, cmp)
	if idx != 2 {
		t.Errorf("BinarySearch returned %d, want 2", idx)
	}
	
	notFound := person{"", 32}
	idx = BinarySearch(arr, notFound, cmp)
	if idx != -1 {
		t.Errorf("BinarySearch should return -1, got %d", idx)
	}
}

func TestBinarySearchLargeArray(t *testing.T) {
	n := 100000
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i * 2 // 0, 2, 4, 6, ...
	}
	
	// Search for existing elements
	for i := 0; i < 1000; i++ {
		target := i * 2
		idx := BinarySearch(arr, target, intCmp)
		if idx != i {
			t.Errorf("BinarySearch failed for %d: got index %d, want %d", target, idx, i)
		}
	}
	
	// Search for non-existing elements (odd numbers)
	for i := 0; i < 1000; i++ {
		target := i*2 + 1
		idx := BinarySearch(arr, target, intCmp)
		if idx != -1 {
			t.Errorf("BinarySearch should return -1 for %d, got %d", target, idx)
		}
	}
}

func TestBinarySearchEdgeCases(t *testing.T) {
	t.Run("empty array", func(t *testing.T) {
		idx := BinarySearch([]int{}, 5, intCmp)
		if idx != -1 {
			t.Errorf("BinarySearch on empty array should return -1, got %d", idx)
		}
	})
	
	t.Run("nil array", func(t *testing.T) {
		var arr []int
		idx := BinarySearch(arr, 5, intCmp)
		if idx != -1 {
			t.Errorf("BinarySearch on nil array should return -1, got %d", idx)
		}
	})
	
	t.Run("zero values", func(t *testing.T) {
		arr := []int{0, 0, 0, 0}
		idx := BinarySearch(arr, 0, intCmp)
		if idx < 0 || idx >= len(arr) {
			t.Errorf("BinarySearch should find 0, got index %d", idx)
		}
	})
}

func TestBinarySearchBoundaries(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	
	// Test first element
	idx := BinarySearch(arr, 1, intCmp)
	if idx != 0 {
		t.Errorf("First element: got %d, want 0", idx)
	}
	
	// Test last element
	idx = BinarySearch(arr, 10, intCmp)
	if idx != 9 {
		t.Errorf("Last element: got %d, want 9", idx)
	}
	
	// Test before first
	idx = BinarySearch(arr, 0, intCmp)
	if idx != -1 {
		t.Errorf("Before first: got %d, want -1", idx)
	}
	
	// Test after last
	idx = BinarySearch(arr, 11, intCmp)
	if idx != -1 {
		t.Errorf("After last: got %d, want -1", idx)
	}
}

// Benchmarks
func BenchmarkBinarySearch(b *testing.B) {
	arr := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		arr[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BinarySearch(arr, 500, intCmp)
	}
}

func BenchmarkBinarySearchLarge(b *testing.B) {
	arr := make([]int, 1000000)
	for i := 0; i < 1000000; i++ {
		arr[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BinarySearch(arr, 500000, intCmp)
	}
}

func BenchmarkBinarySearchNotFound(b *testing.B) {
	arr := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		arr[i] = i * 2
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BinarySearch(arr, 999, intCmp) // odd number, won't be found
	}
}

