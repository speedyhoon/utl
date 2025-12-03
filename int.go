package utl

import "cmp"

// WithinRange sets x to be within a numeric range.
func WithinRange[T cmp.Ordered](x, minimum, maximum T) T {
	return min(maximum, max(minimum, x))
}
