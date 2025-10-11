package collections

// Map transforms a slice of A into a slice of B using fn.
func Map[A any, B any](in []A, fn func(A) B) []B {
	out := make([]B, len(in))
	for i, v := range in {
		out[i] = fn(v)
	}
	return out
}

// Filter returns a slice of elements that satisfy pred.
func Filter[T any](in []T, pred func(T) bool) []T {
	out := make([]T, 0, len(in))
	for _, v := range in {
		if pred(v) {
			out = append(out, v)
		}
	}
	return out
}

// Reduce reduces a slice to a single value using accumulator function.
func Reduce[T any, R any](in []T, init R, fn func(R, T) R) R {
	acc := init
	for _, v := range in {
		acc = fn(acc, v)
	}
	return acc
}
