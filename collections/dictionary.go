// Package collections provides generic data structures and functional programming utilities.
//
// ⚠️  NOT THREAD-SAFE
// This implementation is not safe for concurrent access.
// Wrap with external synchronization (sync.Mutex) if needed.
package collections

// Dictionary is a thin wrapper over a Go map with helpful methods.
// Zero value is ready to use.
type Dictionary[K comparable, V any] struct {
	m map[K]V
}

// New returns an empty Dictionary.
func New[K comparable, V any]() *Dictionary[K, V] {
	return &Dictionary[K, V]{m: make(map[K]V)}
}

// NewDictionary returns an empty Dictionary.
// Deprecated: Use New instead.
func NewDictionary[K comparable, V any]() *Dictionary[K, V] {
	return New[K, V]()
}

// FromMap creates a Dictionary from a map copy.
func FromMap[K comparable, V any](src map[K]V) *Dictionary[K, V] {
	dup := make(map[K]V, len(src))
	for k, v := range src {
		dup[k] = v
	}
	return &Dictionary[K, V]{m: dup}
}

// Get returns the value for key and whether it existed.
func (d *Dictionary[K, V]) Get(key K) (V, bool) {
	var zero V
	if d.m == nil {
		return zero, false
	}
	v, ok := d.m[key]
	return v, ok
}

// Put sets the value for key, returning the previous value and whether it existed.
func (d *Dictionary[K, V]) Put(key K, value V) (V, bool) {
	if d.m == nil {
		d.m = make(map[K]V)
	}
	old, had := d.m[key]
	d.m[key] = value
	return old, had
}

// Delete removes key from the dictionary and returns the previous value and whether it existed.
func (d *Dictionary[K, V]) Delete(key K) (V, bool) {
	var zero V
	if d.m == nil {
		return zero, false
	}
	old, had := d.m[key]
	delete(d.m, key)
	return old, had
}

// Len returns the number of entries.
func (d *Dictionary[K, V]) Len() int { return len(d.m) }

// IsEmpty reports whether the dictionary has no elements.
func (d *Dictionary[K, V]) IsEmpty() bool { return len(d.m) == 0 }

// Keys returns a slice of keys in unspecified order.
func (d *Dictionary[K, V]) Keys() []K {
	out := make([]K, 0, len(d.m))
	for k := range d.m {
		out = append(out, k)
	}
	return out
}

// Values returns a slice of values in unspecified order.
func (d *Dictionary[K, V]) Values() []V {
	out := make([]V, 0, len(d.m))
	for _, v := range d.m {
		out = append(out, v)
	}
	return out
}

// ForEach iterates key-value pairs in unspecified order.
func (d *Dictionary[K, V]) ForEach(fn func(K, V)) {
	for k, v := range d.m {
		fn(k, v)
	}
}

// ToMap returns a shallow copy of the underlying map.
func (d *Dictionary[K, V]) ToMap() map[K]V {
	dup := make(map[K]V, len(d.m))
	for k, v := range d.m {
		dup[k] = v
	}
	return dup
}

// Clone returns a shallow copy of the dictionary.
func (d *Dictionary[K, V]) Clone() *Dictionary[K, V] {
	clone := New[K, V]()
	for k, v := range d.m {
		clone.m[k] = v
	}
	return clone
}
