package utl

import "strconv"

// UtoA returns the string representation of `u` in base 10.
// UtoA is equivalent to strconv.Itoa() but for unsigned integers, shorthand for strconv.FormatUint(uint64(u), 10).
func UtoA(u uint) string {
	return strconv.FormatUint(uint64(u), 10)
}
