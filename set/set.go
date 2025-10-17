package set

// HashSet is a generic set backed by a map.
// Zero value is ready to use.
type HashSet[T comparable] struct {
	m map[T]struct{}
}

// New creates an empty HashSet.
func New[T comparable]() *HashSet[T] { return &HashSet[T]{m: make(map[T]struct{})} }

// FromSlice builds a set from the given slice.
func FromSlice[T comparable](items []T) *HashSet[T] {
	hs := New[T]()
	for _, v := range items {
		hs.Add(v)
	}
	return hs
}

// Add inserts the value into the set. Returns true if it was not present.
func (s *HashSet[T]) Add(v T) bool {
	if s.m == nil {
		s.m = make(map[T]struct{})
	}
	_, existed := s.m[v]
	s.m[v] = struct{}{}
	return !existed
}

// Remove deletes the value from the set. Returns true if it was present.
func (s *HashSet[T]) Remove(v T) bool {
	if s.m == nil {
		return false
	}
	_, existed := s.m[v]
	delete(s.m, v)
	return existed
}

// Contains reports whether the value is in the set.
func (s *HashSet[T]) Contains(v T) bool {
	if s.m == nil {
		return false
	}
	_, ok := s.m[v]
	return ok
}

// Len returns the number of elements in the set.
func (s *HashSet[T]) Len() int { return len(s.m) }

// IsEmpty reports whether the set has no elements.
func (s *HashSet[T]) IsEmpty() bool { return len(s.m) == 0 }

// Clear removes all elements from the set in O(1) time.
func (s *HashSet[T]) Clear() {
	s.m = make(map[T]struct{})
}

// ToSlice returns all elements as a slice, in unspecified order.
func (s *HashSet[T]) ToSlice() []T {
	out := make([]T, 0, len(s.m))
	for v := range s.m {
		out = append(out, v)
	}
	return out
}

// Clone returns a shallow copy of the set.
func (s *HashSet[T]) Clone() *HashSet[T] {
	clone := New[T]()
	for v := range s.m {
		clone.m[v] = struct{}{}
	}
	return clone
}

// Union returns a new set that is the union of a and b.
func Union[T comparable](a, b *HashSet[T]) *HashSet[T] {
	res := a.Clone()
	for v := range b.m {
		res.m[v] = struct{}{}
	}
	return res
}

// Intersection returns a new set that is the intersection of a and b.
func Intersection[T comparable](a, b *HashSet[T]) *HashSet[T] {
	res := New[T]()
	// iterate over smaller set for efficiency
	var small, large *HashSet[T]
	if a.Len() < b.Len() {
		small, large = a, b
	} else {
		small, large = b, a
	}
	for v := range small.m {
		if _, ok := large.m[v]; ok {
			res.m[v] = struct{}{}
		}
	}
	return res
}

// Difference returns a new set that is a - b.
func Difference[T comparable](a, b *HashSet[T]) *HashSet[T] {
	res := New[T]()
	for v := range a.m {
		if _, ok := b.m[v]; !ok {
			res.m[v] = struct{}{}
		}
	}
	return res
}

// IsSubset reports whether a is a subset of b.
func IsSubset[T comparable](a, b *HashSet[T]) bool {
	for v := range a.m {
		if _, ok := b.m[v]; !ok {
			return false
		}
	}
	return true
}

// ForEach iterates over the set elements in unspecified order.
func (s *HashSet[T]) ForEach(fn func(v T)) {
	for v := range s.m {
		fn(v)
	}
}
