package collections

import (
	"sort"
	"testing"
)

func TestDictionaryNew(t *testing.T) {
	d := New[string, int]()
	if d == nil {
		t.Fatal("New() returned nil")
	}
	if !d.IsEmpty() {
		t.Error("New dictionary should be empty")
	}
	if d.Len() != 0 {
		t.Errorf("New dictionary length = %d, want 0", d.Len())
	}
}

func TestNewDictionary(t *testing.T) {
	// Test deprecated function
	d := NewDictionary[string, int]()
	if d == nil {
		t.Fatal("NewDictionary() returned nil")
	}
	if !d.IsEmpty() {
		t.Error("NewDictionary should be empty")
	}
}

func TestFromMap(t *testing.T) {
	src := map[string]int{"a": 1, "b": 2, "c": 3}
	d := FromMap(src)
	
	if d.Len() != len(src) {
		t.Errorf("len = %d, want %d", d.Len(), len(src))
	}
	
	for k, v := range src {
		got, ok := d.Get(k)
		if !ok {
			t.Errorf("dictionary missing key %q", k)
		}
		if got != v {
			t.Errorf("Get(%q) = %d, want %d", k, got, v)
		}
	}
	
	// Verify it's a copy
	src["d"] = 4
	if d.Len() == len(src) {
		t.Error("FromMap should create a copy, not a reference")
	}
}

func TestPut(t *testing.T) {
	d := New[string, int]()
	
	// Put new key
	old, existed := d.Put("a", 1)
	if existed {
		t.Error("putting new key should return existed=false")
	}
	if old != 0 {
		t.Errorf("putting new key returned old value %d, want 0", old)
	}
	
	// Put existing key
	old, existed = d.Put("a", 2)
	if !existed {
		t.Error("putting existing key should return existed=true")
	}
	if old != 1 {
		t.Errorf("old value = %d, want 1", old)
	}
	
	// Verify new value
	val, ok := d.Get("a")
	if !ok || val != 2 {
		t.Errorf("Get('a') = (%d, %v), want (2, true)", val, ok)
	}
}

func TestGet(t *testing.T) {
	d := FromMap(map[string]int{"a": 1, "b": 2, "c": 3})
	
	tests := []struct {
		key     string
		wantVal int
		wantOk  bool
	}{
		{"a", 1, true},
		{"b", 2, true},
		{"c", 3, true},
		{"d", 0, false},
		{"", 0, false},
	}
	
	for _, tt := range tests {
		val, ok := d.Get(tt.key)
		if ok != tt.wantOk {
			t.Errorf("Get(%q) ok = %v, want %v", tt.key, ok, tt.wantOk)
		}
		if ok && val != tt.wantVal {
			t.Errorf("Get(%q) = %d, want %d", tt.key, val, tt.wantVal)
		}
	}
}

func TestDelete(t *testing.T) {
	d := FromMap(map[string]int{"a": 1, "b": 2, "c": 3})
	
	// Delete existing key
	old, existed := d.Delete("b")
	if !existed {
		t.Error("deleting existing key should return existed=true")
	}
	if old != 2 {
		t.Errorf("old value = %d, want 2", old)
	}
	
	// Verify it's gone
	_, ok := d.Get("b")
	if ok {
		t.Error("key should not exist after Delete")
	}
	
	// Delete non-existent key
	old, existed = d.Delete("z")
	if existed {
		t.Error("deleting non-existent key should return existed=false")
	}
	if old != 0 {
		t.Errorf("old value = %d, want 0", old)
	}
}

func TestLen(t *testing.T) {
	d := New[string, int]()
	
	lengths := []int{0, 1, 2, 3, 2, 1, 0}
	ops := []func(){
		func() {},
		func() { d.Put("a", 1) },
		func() { d.Put("b", 2) },
		func() { d.Put("c", 3) },
		func() { d.Delete("c") },
		func() { d.Delete("b") },
		func() { d.Delete("a") },
	}
	
	for i, op := range ops {
		op()
		if d.Len() != lengths[i] {
			t.Errorf("after op[%d], len = %d, want %d", i, d.Len(), lengths[i])
		}
	}
}

func TestIsEmpty(t *testing.T) {
	d := New[string, int]()
	
	if !d.IsEmpty() {
		t.Error("new dictionary should be empty")
	}
	
	d.Put("a", 1)
	if d.IsEmpty() {
		t.Error("dictionary with element should not be empty")
	}
	
	d.Delete("a")
	if !d.IsEmpty() {
		t.Error("dictionary after deleting all should be empty")
	}
}

func TestKeys(t *testing.T) {
	src := map[string]int{"a": 1, "b": 2, "c": 3}
	d := FromMap(src)
	
	keys := d.Keys()
	if len(keys) != len(src) {
		t.Fatalf("Keys len = %d, want %d", len(keys), len(src))
	}
	
	// Sort to compare (order is unspecified)
	sort.Strings(keys)
	want := []string{"a", "b", "c"}
	
	for i := range keys {
		if keys[i] != want[i] {
			t.Errorf("key[%d] = %q, want %q", i, keys[i], want[i])
		}
	}
}

func TestValues(t *testing.T) {
	src := map[string]int{"a": 1, "b": 2, "c": 3}
	d := FromMap(src)
	
	values := d.Values()
	if len(values) != len(src) {
		t.Fatalf("Values len = %d, want %d", len(values), len(src))
	}
	
	// Sort to compare (order is unspecified)
	sort.Ints(values)
	want := []int{1, 2, 3}
	
	for i := range values {
		if values[i] != want[i] {
			t.Errorf("value[%d] = %d, want %d", i, values[i], want[i])
		}
	}
}

func TestForEach(t *testing.T) {
	d := FromMap(map[string]int{"a": 1, "b": 2, "c": 3})
	
	sum := 0
	seen := make(map[string]bool)
	
	d.ForEach(func(k string, v int) {
		sum += v
		seen[k] = true
	})
	
	if sum != 6 {
		t.Errorf("sum = %d, want 6", sum)
	}
	
	for key := range map[string]bool{"a": true, "b": true, "c": true} {
		if !seen[key] {
			t.Errorf("ForEach missed key %q", key)
		}
	}
}

func TestToMap(t *testing.T) {
	src := map[string]int{"a": 1, "b": 2, "c": 3}
	d := FromMap(src)
	
	got := d.ToMap()
	if len(got) != len(src) {
		t.Fatalf("ToMap len = %d, want %d", len(got), len(src))
	}
	
	for k, v := range src {
		if got[k] != v {
			t.Errorf("ToMap()[%q] = %d, want %d", k, got[k], v)
		}
	}
	
	// Verify it's a copy
	got["d"] = 4
	if d.Len() == len(got) {
		t.Error("ToMap should return a copy, not a reference")
	}
}

func TestClone(t *testing.T) {
	original := FromMap(map[string]int{"a": 1, "b": 2, "c": 3})
	clone := original.Clone()
	
	// Verify contents match
	if original.Len() != clone.Len() {
		t.Fatalf("clone len = %d, want %d", clone.Len(), original.Len())
	}
	
	for _, key := range original.Keys() {
		origVal, _ := original.Get(key)
		cloneVal, ok := clone.Get(key)
		if !ok {
			t.Errorf("clone missing key %q", key)
		}
		if origVal != cloneVal {
			t.Errorf("clone[%q] = %d, want %d", key, cloneVal, origVal)
		}
	}
	
	// Verify independence
	clone.Put("d", 999)
	if _, ok := original.Get("d"); ok {
		t.Error("modifying clone should not affect original")
	}
}

func TestZeroValue(t *testing.T) {
	var d Dictionary[string, int]
	
	// Should be usable without calling New
	if !d.IsEmpty() {
		t.Error("zero value dictionary should be empty")
	}
	
	d.Put("a", 1)
	val, ok := d.Get("a")
	if !ok || val != 1 {
		t.Error("zero value dictionary should be usable")
	}
}

func TestGenericTypes(t *testing.T) {
	t.Run("int keys", func(t *testing.T) {
		d := New[int, string]()
		d.Put(1, "one")
		d.Put(2, "two")
		val, ok := d.Get(1)
		if !ok || val != "one" {
			t.Errorf("got %q, want %q", val, "one")
		}
	})
	
	t.Run("struct values", func(t *testing.T) {
		type person struct {
			name string
			age  int
		}
		d := New[string, person]()
		p := person{"Alice", 30}
		d.Put("alice", p)
		val, ok := d.Get("alice")
		if !ok || val != p {
			t.Error("dictionary should work with struct values")
		}
	})
	
	t.Run("pointer values", func(t *testing.T) {
		d := New[string, *int]()
		x := 42
		d.Put("answer", &x)
		val, ok := d.Get("answer")
		if !ok || val != &x {
			t.Error("dictionary should work with pointer values")
		}
	})
}

func TestLargeDictionary(t *testing.T) {
	d := New[int, int]()
	n := 10000
	
	// Add many elements
	for i := 0; i < n; i++ {
		d.Put(i, i*2)
	}
	
	if d.Len() != n {
		t.Errorf("len = %d, want %d", d.Len(), n)
	}
	
	// Verify all present
	for i := 0; i < n; i++ {
		val, ok := d.Get(i)
		if !ok {
			t.Errorf("dictionary missing key %d", i)
		}
		if val != i*2 {
			t.Errorf("Get(%d) = %d, want %d", i, val, i*2)
		}
	}
	
	// Delete all
	for i := 0; i < n; i++ {
		d.Delete(i)
	}
	
	if !d.IsEmpty() {
		t.Error("dictionary should be empty after deleting all")
	}
}

// Benchmarks
func BenchmarkPut(b *testing.B) {
	d := New[int, int]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.Put(i, i)
	}
}

func BenchmarkGet(b *testing.B) {
	d := New[int, int]()
	for i := 0; i < 1000; i++ {
		d.Put(i, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.Get(i % 1000)
	}
}

func BenchmarkDelete(b *testing.B) {
	d := New[int, int]()
	for i := 0; i < b.N; i++ {
		d.Put(i, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.Delete(i)
	}
}

