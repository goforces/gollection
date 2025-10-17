package collections

import (
	"testing"
)

func TestMap(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		fn    func(int) int
		want  []int
	}{
		{
			name:  "empty slice",
			input: []int{},
			fn:    func(x int) int { return x * 2 },
			want:  []int{},
		},
		{
			name:  "double values",
			input: []int{1, 2, 3, 4, 5},
			fn:    func(x int) int { return x * 2 },
			want:  []int{2, 4, 6, 8, 10},
		},
		{
			name:  "square values",
			input: []int{1, 2, 3, 4},
			fn:    func(x int) int { return x * x },
			want:  []int{1, 4, 9, 16},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Map(tt.input, tt.fn)
			
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

func TestMapTypeConversion(t *testing.T) {
	// Test converting int to string
	input := []int{1, 2, 3}
	result := Map(input, func(x int) string {
		return string(rune('0' + x))
	})
	
	want := []string{"1", "2", "3"}
	if len(result) != len(want) {
		t.Fatalf("len = %d, want %d", len(result), len(want))
	}
	
	for i := range result {
		if result[i] != want[i] {
			t.Errorf("element[%d] = %q, want %q", i, result[i], want[i])
		}
	}
}

func TestFilter(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		pred  func(int) bool
		want  []int
	}{
		{
			name:  "empty slice",
			input: []int{},
			pred:  func(x int) bool { return x > 0 },
			want:  []int{},
		},
		{
			name:  "filter even numbers",
			input: []int{1, 2, 3, 4, 5, 6},
			pred:  func(x int) bool { return x%2 == 0 },
			want:  []int{2, 4, 6},
		},
		{
			name:  "filter greater than 3",
			input: []int{1, 2, 3, 4, 5},
			pred:  func(x int) bool { return x > 3 },
			want:  []int{4, 5},
		},
		{
			name:  "no matches",
			input: []int{1, 2, 3},
			pred:  func(x int) bool { return x > 10 },
			want:  []int{},
		},
		{
			name:  "all match",
			input: []int{1, 2, 3},
			pred:  func(x int) bool { return x > 0 },
			want:  []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Filter(tt.input, tt.pred)
			
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

func TestFilterStrings(t *testing.T) {
	input := []string{"apple", "banana", "cherry", "date", "elderberry"}
	result := Filter(input, func(s string) bool {
		return len(s) > 5
	})
	
	want := []string{"banana", "cherry", "elderberry"}
	if len(result) != len(want) {
		t.Fatalf("len = %d, want %d", len(result), len(want))
	}
	
	for i := range result {
		if result[i] != want[i] {
			t.Errorf("element[%d] = %q, want %q", i, result[i], want[i])
		}
	}
}

func TestReduce(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		init  int
		fn    func(int, int) int
		want  int
	}{
		{
			name:  "empty slice",
			input: []int{},
			init:  0,
			fn:    func(acc, x int) int { return acc + x },
			want:  0,
		},
		{
			name:  "sum",
			input: []int{1, 2, 3, 4, 5},
			init:  0,
			fn:    func(acc, x int) int { return acc + x },
			want:  15,
		},
		{
			name:  "product",
			input: []int{1, 2, 3, 4, 5},
			init:  1,
			fn:    func(acc, x int) int { return acc * x },
			want:  120,
		},
		{
			name:  "max",
			input: []int{3, 1, 4, 1, 5, 9, 2, 6},
			init:  0,
			fn: func(acc, x int) int {
				if x > acc {
					return x
				}
				return acc
			},
			want: 9,
		},
		{
			name:  "count evens",
			input: []int{1, 2, 3, 4, 5, 6},
			init:  0,
			fn: func(acc, x int) int {
				if x%2 == 0 {
					return acc + 1
				}
				return acc
			},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Reduce(tt.input, tt.init, tt.fn)
			
			if got != tt.want {
				t.Errorf("got = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestReduceTypeConversion(t *testing.T) {
	// Test reducing []int to string
	input := []int{1, 2, 3}
	result := Reduce(input, "", func(acc string, x int) string {
		if acc == "" {
			return string(rune('0' + x))
		}
		return acc + "," + string(rune('0'+x))
	})
	
	want := "1,2,3"
	if result != want {
		t.Errorf("got %q, want %q", result, want)
	}
}

func TestReduceToMap(t *testing.T) {
	// Test building a map from a slice
	type pair struct {
		key   string
		value int
	}
	
	input := []pair{
		{"a", 1},
		{"b", 2},
		{"c", 3},
	}
	
	result := Reduce(input, make(map[string]int), func(acc map[string]int, p pair) map[string]int {
		acc[p.key] = p.value
		return acc
	})
	
	want := map[string]int{"a": 1, "b": 2, "c": 3}
	
	if len(result) != len(want) {
		t.Errorf("len = %d, want %d", len(result), len(want))
	}
	
	for k, v := range want {
		if result[k] != v {
			t.Errorf("result[%q] = %d, want %d", k, result[k], v)
		}
	}
}

func TestChaining(t *testing.T) {
	// Test chaining Map, Filter, and Reduce
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	
	// Double all numbers, filter evens, sum them
	doubled := Map(input, func(x int) int { return x * 2 })
	evens := Filter(doubled, func(x int) bool { return x%4 == 0 })
	sum := Reduce(evens, 0, func(acc, x int) int { return acc + x })
	
	// doubled: 2,4,6,8,10,12,14,16,18,20
	// evens (divisible by 4): 4,8,12,16,20
	// sum: 60
	want := 60
	if sum != want {
		t.Errorf("chained result = %d, want %d", sum, want)
	}
}

func TestMapPreservesOrder(t *testing.T) {
	input := []int{5, 3, 8, 1, 9, 2}
	result := Map(input, func(x int) int { return x * 2 })
	want := []int{10, 6, 16, 2, 18, 4}
	
	for i := range result {
		if result[i] != want[i] {
			t.Errorf("element[%d] = %d, want %d (order not preserved)", i, result[i], want[i])
		}
	}
}

func TestFilterPreservesOrder(t *testing.T) {
	input := []int{5, 3, 8, 1, 9, 2, 7}
	result := Filter(input, func(x int) bool { return x > 4 })
	want := []int{5, 8, 9, 7}
	
	for i := range result {
		if result[i] != want[i] {
			t.Errorf("element[%d] = %d, want %d (order not preserved)", i, result[i], want[i])
		}
	}
}

// Benchmarks
func BenchmarkMap(b *testing.B) {
	input := make([]int, 1000)
	for i := range input {
		input[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Map(input, func(x int) int { return x * 2 })
	}
}

func BenchmarkFilter(b *testing.B) {
	input := make([]int, 1000)
	for i := range input {
		input[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Filter(input, func(x int) bool { return x%2 == 0 })
	}
}

func BenchmarkReduce(b *testing.B) {
	input := make([]int, 1000)
	for i := range input {
		input[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reduce(input, 0, func(acc, x int) int { return acc + x })
	}
}

