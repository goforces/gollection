package tree

import (
	"testing"
)

// intCmp is a standard comparator for integers
func intCmp(a, b int) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

// stringCmp is a standard comparator for strings
func stringCmp(a, b string) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

func TestNew(t *testing.T) {
	tree := New[int, string](intCmp)
	if tree == nil {
		t.Fatal("New() returned nil")
	}
	if tree.root != nil {
		t.Error("New tree should have nil root")
	}
}

func TestNewBinaryTree(t *testing.T) {
	// Test deprecated function
	tree := NewBinaryTree[int, string](intCmp)
	if tree == nil {
		t.Fatal("NewBinaryTree() returned nil")
	}
}

func TestPutGet(t *testing.T) {
	tree := New[int, string](intCmp)

	// Put single value
	tree.Put(5, "five")
	val, ok := tree.Get(5)
	if !ok {
		t.Error("Get(5) should return true after Put(5)")
	}
	if val != "five" {
		t.Errorf("Get(5) = %q, want %q", val, "five")
	}

	// Put more values
	tree.Put(3, "three")
	tree.Put(7, "seven")
	tree.Put(1, "one")
	tree.Put(9, "nine")

	// Verify all values
	tests := []struct {
		key     int
		wantVal string
		wantOk  bool
	}{
		{1, "one", true},
		{3, "three", true},
		{5, "five", true},
		{7, "seven", true},
		{9, "nine", true},
		{0, "", false},
		{2, "", false},
		{10, "", false},
	}

	for _, tt := range tests {
		val, ok := tree.Get(tt.key)
		if ok != tt.wantOk {
			t.Errorf("Get(%d) ok = %v, want %v", tt.key, ok, tt.wantOk)
		}
		if ok && val != tt.wantVal {
			t.Errorf("Get(%d) = %q, want %q", tt.key, val, tt.wantVal)
		}
	}
}

func TestPutUpdate(t *testing.T) {
	tree := New[int, string](intCmp)

	// Put initial value
	tree.Put(5, "five")

	// Update value
	tree.Put(5, "FIVE")

	val, ok := tree.Get(5)
	if !ok {
		t.Error("Get(5) should return true")
	}
	if val != "FIVE" {
		t.Errorf("Get(5) = %q, want %q", val, "FIVE")
	}
}

func TestGetEmpty(t *testing.T) {
	tree := New[int, string](intCmp)

	val, ok := tree.Get(5)
	if ok {
		t.Error("Get on empty tree should return false")
	}
	if val != "" {
		t.Errorf("Get on empty tree returned %q, want zero value", val)
	}
}

func TestContains(t *testing.T) {
	tree := New[int, string](intCmp)

	tree.Put(5, "five")
	tree.Put(3, "three")
	tree.Put(7, "seven")

	tests := []struct {
		key  int
		want bool
	}{
		{3, true},
		{5, true},
		{7, true},
		{1, false},
		{10, false},
	}

	for _, tt := range tests {
		got := tree.Contains(tt.key)
		if got != tt.want {
			t.Errorf("Contains(%d) = %v, want %v", tt.key, got, tt.want)
		}
	}
}

func TestDelete(t *testing.T) {
	tree := New[int, string](intCmp)

	// Build tree
	tree.Put(5, "five")
	tree.Put(3, "three")
	tree.Put(7, "seven")
	tree.Put(1, "one")
	tree.Put(4, "four")
	tree.Put(6, "six")
	tree.Put(9, "nine")

	// Delete existing key
	deleted := tree.Delete(3)
	if !deleted {
		t.Error("Delete(3) should return true")
	}
	if tree.Contains(3) {
		t.Error("tree should not contain deleted key")
	}

	// Verify other keys still present
	for _, key := range []int{1, 4, 5, 6, 7, 9} {
		if !tree.Contains(key) {
			t.Errorf("tree missing key %d after delete", key)
		}
	}

	// Delete non-existent key
	deleted = tree.Delete(999)
	if deleted {
		t.Error("Delete(999) should return false")
	}
}

func TestDeleteLeaf(t *testing.T) {
	tree := New[int, string](intCmp)

	tree.Put(5, "five")
	tree.Put(3, "three")
	tree.Put(7, "seven")

	// Delete leaf node
	deleted := tree.Delete(3)
	if !deleted {
		t.Error("Delete leaf should return true")
	}
	if tree.Contains(3) {
		t.Error("tree should not contain deleted leaf")
	}

	// Root and other child should remain
	if !tree.Contains(5) || !tree.Contains(7) {
		t.Error("other nodes should remain after deleting leaf")
	}
}

func TestDeleteNodeWithOneChild(t *testing.T) {
	tree := New[int, string](intCmp)

	tree.Put(5, "five")
	tree.Put(3, "three")
	tree.Put(1, "one")

	// Delete node with one child (3 has left child 1)
	deleted := tree.Delete(3)
	if !deleted {
		t.Error("Delete should return true")
	}
	if tree.Contains(3) {
		t.Error("tree should not contain deleted node")
	}

	// Child should still be accessible
	if !tree.Contains(1) {
		t.Error("child of deleted node should remain")
	}
}

func TestDeleteNodeWithTwoChildren(t *testing.T) {
	tree := New[int, string](intCmp)

	tree.Put(5, "five")
	tree.Put(3, "three")
	tree.Put(7, "seven")
	tree.Put(1, "one")
	tree.Put(4, "four")
	tree.Put(6, "six")
	tree.Put(9, "nine")

	// Delete node with two children
	deleted := tree.Delete(5) // root with two children
	if !deleted {
		t.Error("Delete should return true")
	}
	if tree.Contains(5) {
		t.Error("tree should not contain deleted node")
	}

	// All other nodes should still be present
	for _, key := range []int{1, 3, 4, 6, 7, 9} {
		if !tree.Contains(key) {
			t.Errorf("tree missing key %d after delete", key)
		}
	}
}

func TestDeleteRoot(t *testing.T) {
	tree := New[int, string](intCmp)

	tree.Put(5, "five")

	deleted := tree.Delete(5)
	if !deleted {
		t.Error("Delete root should return true")
	}
	if tree.Contains(5) {
		t.Error("tree should not contain deleted root")
	}
}

func TestDeleteAll(t *testing.T) {
	tree := New[int, string](intCmp)

	keys := []int{5, 3, 7, 1, 4, 6, 9}
	for _, k := range keys {
		tree.Put(k, "value")
	}

	// Delete all keys
	for _, k := range keys {
		deleted := tree.Delete(k)
		if !deleted {
			t.Errorf("Delete(%d) should return true", k)
		}
	}

	// Tree should be empty
	for _, k := range keys {
		if tree.Contains(k) {
			t.Errorf("tree should not contain %d", k)
		}
	}
}

func TestClone(t *testing.T) {
	original := New[int, string](intCmp)

	original.Put(5, "five")
	original.Put(3, "three")
	original.Put(7, "seven")

	clone := original.Clone()

	// Verify contents match
	for _, key := range []int{3, 5, 7} {
		origVal, origOk := original.Get(key)
		cloneVal, cloneOk := clone.Get(key)

		if origOk != cloneOk {
			t.Errorf("key %d: original ok=%v, clone ok=%v", key, origOk, cloneOk)
		}
		if origVal != cloneVal {
			t.Errorf("key %d: original=%q, clone=%q", key, origVal, cloneVal)
		}
	}

	// Verify independence
	clone.Put(9, "nine")
	if original.Contains(9) {
		t.Error("modifying clone should not affect original")
	}
}

func TestStringKeys(t *testing.T) {
	tree := New[string, int](stringCmp)

	tree.Put("apple", 1)
	tree.Put("banana", 2)
	tree.Put("cherry", 3)

	val, ok := tree.Get("banana")
	if !ok || val != 2 {
		t.Errorf("Get('banana') = (%d, %v), want (2, true)", val, ok)
	}

	deleted := tree.Delete("banana")
	if !deleted || tree.Contains("banana") {
		t.Error("Delete('banana') failed")
	}
}

func TestCustomComparator(t *testing.T) {
	// Reverse comparator (descending order)
	reverseCmp := func(a, b int) int {
		if a < b {
			return 1
		}
		if a > b {
			return -1
		}
		return 0
	}

	tree := New[int, string](reverseCmp)

	tree.Put(5, "five")
	tree.Put(3, "three")
	tree.Put(7, "seven")

	// Should still work correctly with reverse comparator
	val, ok := tree.Get(5)
	if !ok || val != "five" {
		t.Error("tree should work with custom comparator")
	}
}

func TestDuplicateKeys(t *testing.T) {
	tree := New[int, string](intCmp)

	tree.Put(5, "first")
	tree.Put(5, "second")
	tree.Put(5, "third")

	// Should only have one entry with latest value
	val, ok := tree.Get(5)
	if !ok {
		t.Error("Get(5) should return true")
	}
	if val != "third" {
		t.Errorf("Get(5) = %q, want %q", val, "third")
	}
}

func TestStructKeys(t *testing.T) {
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

	tree := New[person, string](cmp)

	p1 := person{"Alice", 30}
	p2 := person{"Bob", 25}
	p3 := person{"Charlie", 35}

	tree.Put(p1, "Alice's data")
	tree.Put(p2, "Bob's data")
	tree.Put(p3, "Charlie's data")

	val, ok := tree.Get(p2)
	if !ok || val != "Bob's data" {
		t.Error("tree should work with struct keys")
	}
}

func TestPointerValues(t *testing.T) {
	tree := New[int, *string](intCmp)

	s1 := "hello"
	s2 := "world"

	tree.Put(1, &s1)
	tree.Put(2, &s2)

	val, ok := tree.Get(1)
	if !ok || val != &s1 {
		t.Error("tree should work with pointer values")
	}
}

func TestEmptyTree(t *testing.T) {
	tree := New[int, string](intCmp)

	// Get on empty
	_, ok := tree.Get(1)
	if ok {
		t.Error("Get on empty tree should return false")
	}

	// Contains on empty
	if tree.Contains(1) {
		t.Error("Contains on empty tree should return false")
	}

	// Delete on empty
	if tree.Delete(1) {
		t.Error("Delete on empty tree should return false")
	}
}

func TestLargeTree(t *testing.T) {
	tree := New[int, int](intCmp)
	n := 10000

	// Insert many elements
	for i := 0; i < n; i++ {
		tree.Put(i, i*2)
	}

	// Verify all present
	for i := 0; i < n; i++ {
		val, ok := tree.Get(i)
		if !ok {
			t.Errorf("tree missing key %d", i)
		}
		if val != i*2 {
			t.Errorf("Get(%d) = %d, want %d", i, val, i*2)
		}
	}

	// Delete half
	for i := 0; i < n; i += 2 {
		tree.Delete(i)
	}

	// Verify correct keys remain
	for i := 0; i < n; i++ {
		shouldExist := i%2 == 1
		exists := tree.Contains(i)
		if exists != shouldExist {
			t.Errorf("key %d: exists=%v, want %v", i, exists, shouldExist)
		}
	}
}

func TestBalancedInsertion(t *testing.T) {
	tree := New[int, string](intCmp)

	// Insert in order (creates unbalanced tree, but should still work)
	for i := 1; i <= 10; i++ {
		tree.Put(i, "value")
	}

	// Verify all accessible
	for i := 1; i <= 10; i++ {
		if !tree.Contains(i) {
			t.Errorf("tree missing key %d", i)
		}
	}
}

func TestReverseOrderInsertion(t *testing.T) {
	tree := New[int, string](intCmp)

	// Insert in reverse order
	for i := 10; i >= 1; i-- {
		tree.Put(i, "value")
	}

	// Verify all accessible
	for i := 1; i <= 10; i++ {
		if !tree.Contains(i) {
			t.Errorf("tree missing key %d", i)
		}
	}
}

// Benchmarks
func BenchmarkPut(b *testing.B) {
	tree := New[int, int](intCmp)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Put(i, i)
	}
}

func BenchmarkGet(b *testing.B) {
	tree := New[int, int](intCmp)
	for i := 0; i < 1000; i++ {
		tree.Put(i, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Get(i % 1000)
	}
}

func BenchmarkDelete(b *testing.B) {
	tree := New[int, int](intCmp)
	for i := 0; i < b.N; i++ {
		tree.Put(i, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Delete(i)
	}
}

func BenchmarkContains(b *testing.B) {
	tree := New[int, int](intCmp)
	for i := 0; i < 1000; i++ {
		tree.Put(i, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Contains(i % 1000)
	}
}
