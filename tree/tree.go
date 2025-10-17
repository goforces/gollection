// Package tree provides a generic binary search tree implementation.
//
// ⚠️  NOT THREAD-SAFE
// This implementation is not safe for concurrent access.
// Wrap with external synchronization (sync.Mutex) if needed.
package tree

// BinaryTree is a basic binary search tree for ordered keys using comparator.
type BinaryTree[K any, V any] struct {
	root *node[K, V]
	cmp  func(a, b K) int
}

type node[K any, V any] struct {
	key   K
	val   V
	left  *node[K, V]
	right *node[K, V]
}

// NewBinaryTree creates an empty BST using the provided comparator.
// cmp(a,b) should return -1 if a<b, 0 if equal, 1 if a>b.
func NewBinaryTree[K any, V any](cmp func(a, b K) int) *BinaryTree[K, V] {
	return &BinaryTree[K, V]{cmp: cmp}
}

// Put inserts or replaces a key.
func (t *BinaryTree[K, V]) Put(k K, v V) {
	t.root = put(t.root, k, v, t.cmp)
}

func put[K any, V any](n *node[K, V], k K, v V, cmp func(a, b K) int) *node[K, V] {
	if n == nil {
		return &node[K, V]{key: k, val: v}
	}
	s := cmp(k, n.key)
	if s < 0 {
		n.left = put(n.left, k, v, cmp)
	} else if s > 0 {
		n.right = put(n.right, k, v, cmp)
	} else {
		n.val = v
	}
	return n
}

// Get retrieves value for key; ok is false if not present.
func (t *BinaryTree[K, V]) Get(k K) (V, bool) {
	cur := t.root
	for cur != nil {
		s := t.cmp(k, cur.key)
		if s == 0 {
			return cur.val, true
		}
		if s < 0 {
			cur = cur.left
		} else {
			cur = cur.right
		}
	}
	var zero V
	return zero, false
}

// Contains reports whether key exists.
func (t *BinaryTree[K, V]) Contains(k K) bool {
	_, ok := t.Get(k)
	return ok
}

// Delete removes the key from the tree. Returns true if the key was present.
func (t *BinaryTree[K, V]) Delete(k K) bool {
	if !t.Contains(k) {
		return false
	}
	t.root = deleteNode(t.root, k, t.cmp)
	return true
}

func deleteNode[K any, V any](n *node[K, V], k K, cmp func(a, b K) int) *node[K, V] {
	if n == nil {
		return nil
	}
	
	s := cmp(k, n.key)
	if s < 0 {
		n.left = deleteNode(n.left, k, cmp)
	} else if s > 0 {
		n.right = deleteNode(n.right, k, cmp)
	} else {
		// Found the node to delete
		// Case 1: No children or one child
		if n.left == nil {
			return n.right
		}
		if n.right == nil {
			return n.left
		}
		
		// Case 2: Two children - find in-order successor (min in right subtree)
		successor := findMin(n.right)
		n.key = successor.key
		n.val = successor.val
		n.right = deleteNode(n.right, successor.key, cmp)
	}
	return n
}

func findMin[K any, V any](n *node[K, V]) *node[K, V] {
	for n.left != nil {
		n = n.left
	}
	return n
}
