// Package algorithms provides generic sorting and searching algorithms.
//
// Note: These are stateless utility functions and are safe for concurrent use.
package algorithms

// BinarySearch returns the index of target in a sorted slice using cmp comparator.
// cmp(a, b) should return -1 if a<b, 0 if a==b, 1 if a>b.
// If not found, returns -1.
func BinarySearch[T any](arr []T, target T, cmp func(a, b T) int) int {
	lo, hi := 0, len(arr)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		c := cmp(arr[mid], target)
		if c == 0 {
			return mid
		}
		if c < 0 {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return -1
}
