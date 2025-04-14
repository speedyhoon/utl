package utl

// Del removes an item at the index provided.
func Del[T any](t *[]T, index uint) {
	l := Len(*t)
	if l == 0 || index >= l {
		return
	}

	del(t, l, index)
}

// del is only to be used when guaranteed not to be out of range, otherwise it will panic.
func del[T any](t *[]T, length, index uint) {
	switch index {
	case 0: // First item.
		*t = (*t)[1:]
	case length - 1: // Last item.
		*t = (*t)[:index]
	default:
		*t = append((*t)[:index], (*t)[index+1:]...)
	}
}

// DelDup removes secondary identical items duplicated in a slice.
func DelDup[T comparable](s *[]T) {
	// Loop from the first item through to the second last item.
	for i, l := 0, len(*s); i < l-1; i++ {
		// Loop from the next item (i+1) through to the last item.
		for j := i + 1; j < l; {
			if (*s)[i] == (*s)[j] {
				del(s, uint(l), uint(j))
				l--
				continue
			}
			j++
		}
	}
}

// Len returns the length of a slice as an unsigned integer.
func Len[T any](t []T) uint {
	return uint(len(t))
}
