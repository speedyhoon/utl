package utl

import "math"

// LogBaseX calculates the logarithm of value `x` with a specified base `base` (which is the inverse function of exponentiation with base `base`).
func LogBaseX(base, x float64) float64 {
	return math.Log(x) / math.Log(base)
}

// LogBaseXUint calculates the LogBaseX for unsigned integers rounded up to the nearest integer.
func LogBaseXUint(base, x uint) uint {
	return uint(math.Ceil(math.Log(float64(x)) / math.Log(float64(base))))
}
