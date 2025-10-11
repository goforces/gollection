package algorithms

// QuickSort sorts the slice in-place using provided less comparator.
func QuickSort[T any](arr []T, less func(a, b T) bool) {
	if len(arr) < 2 {
		return
	}
	quickSort(arr, 0, len(arr)-1, less)
}

func quickSort[T any](a []T, lo, hi int, less func(a, b T) bool) {
	for lo < hi {
		p := partition(a, lo, hi, less)
		if p-lo < hi-p {
			quickSort(a, lo, p-1, less)
			lo = p + 1
		} else {
			quickSort(a, p+1, hi, less)
			hi = p - 1
		}
	}
}

func partition[T any](a []T, lo, hi int, less func(a, b T) bool) int {
	pivot := a[hi]
	i := lo
	for j := lo; j < hi; j++ {
		if less(a[j], pivot) {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[i], a[hi] = a[hi], a[i]
	return i
}
