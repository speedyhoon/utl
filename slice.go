package utl

import "strings"

// StrRemoveEmpty removes items that only contain whitespace.
func StrRemoveEmpty(s []string) string {
	for i := 0; i < len(s); {
		if s[i] = strings.TrimSpace(s[i]); s[i] != "" {
			i++
			continue
		}

		s = RemoveItem(s, uint(i))
	}

	return strings.Join(s, " ")
}

// RemoveItem removes items that only contain whitespace.
func RemoveItem(s []string, n uint) []string {
	l := uint(len(s))
	if l == 0 {
		return s
	}

	l--
	if n > l {
		return s
	}

	switch n {
	case 0:
		s = s[1:]
	case l:
		s = s[:l]
	default:
		s = append(s[:n], s[n+1:]...)
	}

	return s
}

// FromEnd returns the item index from the end of the slice, where 0 returns the last index and 1 the second last etc.
func FromEnd(list []string, indexFromEnd uint) string {
	l := uint(len(list))
	if l == 0 || indexFromEnd >= l {
		return ""
	}

	return list[l-1-indexFromEnd]
}

// LoopAround increments or decrements `index` by `direction`.
// If the result is outside the range of the list `qty`, then the index is inverted.
func LoopAround(qty int, index *uint8, direction int8) {
	n := int(*index) + int(direction)
	switch {
	case n >= qty:
		*index = 0
	case n < 0:
		*index = uint8(qty)
	default:
		*index = uint8(n)
	}
}
