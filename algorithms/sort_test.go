package algorithms

import (
	"math/rand"
	"testing"
)

func intLess(a, b int) bool {
	return a < b
}

func stringLess(a, b string) bool {
	return a < b
}

func isSorted(arr []int, less func(a, b int) bool) bool {
	for i := 1; i < len(arr); i++ {
		if less(arr[i], arr[i-1]) {
			return false
		}
	}
	return true
}

func TestQuickSortBasic(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{"empty", []int{}, []int{}},
		{"single", []int{1}, []int{1}},
		{"two ascending", []int{1, 2}, []int{1, 2}},
		{"two descending", []int{2, 1}, []int{1, 2}},
		{"sorted", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"reverse", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"random", []int{3, 1, 4, 1, 5, 9, 2, 6}, []int{1, 1, 2, 3, 4, 5, 6, 9}},
		{"duplicates", []int{5, 2, 5, 2, 5}, []int{2, 2, 5, 5, 5}},
		{"all same", []int{3, 3, 3, 3}, []int{3, 3, 3, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr := make([]int, len(tt.input))
			copy(arr, tt.input)
			
			QuickSort(arr, intLess)
			
			if len(arr) != len(tt.want) {
				t.Errorf("length = %d, want %d", len(arr), len(tt.want))
				return
			}
			
			for i := range arr {
				if arr[i] != tt.want[i] {
					t.Errorf("element[%d] = %d, want %d", i, arr[i], tt.want[i])
				}
			}
		})
	}
}

func TestQuickSortNegativeNumbers(t *testing.T) {
	arr := []int{3, -1, 4, -5, 2, 0, -3}
	want := []int{-5, -3, -1, 0, 2, 3, 4}
	
	QuickSort(arr, intLess)
	
	for i := range arr {
		if arr[i] != want[i] {
			t.Errorf("element[%d] = %d, want %d", i, arr[i], want[i])
		}
	}
}

func TestQuickSortStrings(t *testing.T) {
	arr := []string{"zebra", "apple", "mango", "banana", "cherry"}
	want := []string{"apple", "banana", "cherry", "mango", "zebra"}
	
	QuickSort(arr, stringLess)
	
	for i := range arr {
		if arr[i] != want[i] {
			t.Errorf("element[%d] = %q, want %q", i, arr[i], want[i])
		}
	}
}

func TestQuickSortDescending(t *testing.T) {
	arr := []int{3, 1, 4, 1, 5, 9, 2, 6}
	want := []int{9, 6, 5, 4, 3, 2, 1, 1}
	
	// Reverse comparator
	greater := func(a, b int) bool { return a > b }
	QuickSort(arr, greater)
	
	for i := range arr {
		if arr[i] != want[i] {
			t.Errorf("element[%d] = %d, want %d", i, arr[i], want[i])
		}
	}
}

func TestQuickSortStability(t *testing.T) {
	// QuickSort is not guaranteed to be stable, but it should still sort correctly
	type item struct {
		key   int
		value string
	}
	
	arr := []item{
		{3, "a"},
		{1, "b"},
		{2, "c"},
		{1, "d"},
		{3, "e"},
	}
	
	less := func(a, b item) bool { return a.key < b.key }
	QuickSort(arr, less)
	
	// Check sorted by key
	for i := 1; i < len(arr); i++ {
		if arr[i].key < arr[i-1].key {
			t.Errorf("not sorted: arr[%d].key=%d > arr[%d].key=%d", i-1, arr[i-1].key, i, arr[i].key)
		}
	}
}

func TestQuickSortCustomType(t *testing.T) {
	type person struct {
		name string
		age  int
	}
	
	arr := []person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
		{"Dave", 20},
	}
	
	less := func(a, b person) bool { return a.age < b.age }
	QuickSort(arr, less)
	
	want := []person{
		{"Dave", 20},
		{"Bob", 25},
		{"Alice", 30},
		{"Charlie", 35},
	}
	
	for i := range arr {
		if arr[i] != want[i] {
			t.Errorf("element[%d] = %v, want %v", i, arr[i], want[i])
		}
	}
}

func TestQuickSortInPlace(t *testing.T) {
	arr := []int{5, 2, 8, 1, 9}
	original := arr // same underlying array
	
	QuickSort(arr, intLess)
	
	// Verify it modified the original array
	if !isSorted(original, intLess) {
		t.Error("QuickSort should modify array in-place")
	}
}

func TestQuickSortLargeArray(t *testing.T) {
	n := 10000
	arr := make([]int, n)
	
	// Fill with random values
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(n)
	}
	
	QuickSort(arr, intLess)
	
	// Verify sorted
	if !isSorted(arr, intLess) {
		t.Error("large array not sorted correctly")
	}
}

func TestQuickSortAlreadySorted(t *testing.T) {
	n := 1000
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	
	QuickSort(arr, intLess)
	
	// Should still be sorted
	for i := 0; i < n; i++ {
		if arr[i] != i {
			t.Errorf("element[%d] = %d, want %d", i, arr[i], i)
		}
	}
}

func TestQuickSortReverseSorted(t *testing.T) {
	n := 1000
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = n - i - 1
	}
	
	QuickSort(arr, intLess)
	
	// Should be sorted
	for i := 0; i < n; i++ {
		if arr[i] != i {
			t.Errorf("element[%d] = %d, want %d", i, arr[i], i)
		}
	}
}

func TestQuickSortAllSame(t *testing.T) {
	arr := make([]int, 1000)
	for i := range arr {
		arr[i] = 42
	}
	
	QuickSort(arr, intLess)
	
	// Should all still be 42
	for i, v := range arr {
		if v != 42 {
			t.Errorf("element[%d] = %d, want 42", i, v)
		}
	}
}

func TestQuickSortManyDuplicates(t *testing.T) {
	arr := []int{}
	for i := 0; i < 100; i++ {
		arr = append(arr, 1, 2, 3, 4, 5)
	}
	
	QuickSort(arr, intLess)
	
	// Verify sorted
	if !isSorted(arr, intLess) {
		t.Error("array with many duplicates not sorted correctly")
	}
}

func TestQuickSortEdgeCases(t *testing.T) {
	t.Run("nil array", func(t *testing.T) {
		var arr []int
		QuickSort(arr, intLess) // should not panic
	})
	
	t.Run("empty array", func(t *testing.T) {
		arr := []int{}
		QuickSort(arr, intLess) // should not panic
		if len(arr) != 0 {
			t.Error("empty array modified")
		}
	})
	
	t.Run("single element", func(t *testing.T) {
		arr := []int{42}
		QuickSort(arr, intLess)
		if arr[0] != 42 {
			t.Error("single element modified")
		}
	})
}

func TestQuickSortZeroValues(t *testing.T) {
	arr := []int{0, 5, 0, 3, 0, 1, 0}
	want := []int{0, 0, 0, 0, 1, 3, 5}
	
	QuickSort(arr, intLess)
	
	for i := range arr {
		if arr[i] != want[i] {
			t.Errorf("element[%d] = %d, want %d", i, arr[i], want[i])
		}
	}
}

func TestQuickSortPointers(t *testing.T) {
	v1, v2, v3, v4 := 4, 1, 3, 2
	arr := []*int{&v1, &v2, &v3, &v4}
	
	less := func(a, b *int) bool { return *a < *b }
	QuickSort(arr, less)
	
	want := []*int{&v2, &v4, &v3, &v1}
	for i := range arr {
		if arr[i] != want[i] {
			t.Errorf("element[%d] points to wrong value", i)
		}
	}
}

func TestQuickSortFloat64(t *testing.T) {
	arr := []float64{3.14, 1.41, 2.71, 1.73, 0.0, -1.0}
	want := []float64{-1.0, 0.0, 1.41, 1.73, 2.71, 3.14}
	
	less := func(a, b float64) bool { return a < b }
	QuickSort(arr, less)
	
	for i := range arr {
		if arr[i] != want[i] {
			t.Errorf("element[%d] = %f, want %f", i, arr[i], want[i])
		}
	}
}

func TestQuickSortPartiallyOrdered(t *testing.T) {
	// Array with some ordered sections
	arr := []int{1, 2, 3, 10, 9, 8, 4, 5, 6, 7}
	
	QuickSort(arr, intLess)
	
	for i := 0; i < 10; i++ {
		if arr[i] != i+1 {
			t.Errorf("element[%d] = %d, want %d", i, arr[i], i+1)
		}
	}
}

// Benchmarks
func BenchmarkQuickSort(b *testing.B) {
	arr := make([]int, 1000)
	for i := range arr {
		arr[i] = rand.Intn(1000)
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		temp := make([]int, len(arr))
		copy(temp, arr)
		b.StartTimer()
		
		QuickSort(temp, intLess)
	}
}

func BenchmarkQuickSortSorted(b *testing.B) {
	arr := make([]int, 1000)
	for i := range arr {
		arr[i] = i
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		temp := make([]int, len(arr))
		copy(temp, arr)
		b.StartTimer()
		
		QuickSort(temp, intLess)
	}
}

func BenchmarkQuickSortReverse(b *testing.B) {
	arr := make([]int, 1000)
	for i := range arr {
		arr[i] = 1000 - i
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		temp := make([]int, len(arr))
		copy(temp, arr)
		b.StartTimer()
		
		QuickSort(temp, intLess)
	}
}

func BenchmarkQuickSortLarge(b *testing.B) {
	arr := make([]int, 100000)
	for i := range arr {
		arr[i] = rand.Intn(100000)
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		temp := make([]int, len(arr))
		copy(temp, arr)
		b.StartTimer()
		
		QuickSort(temp, intLess)
	}
}

