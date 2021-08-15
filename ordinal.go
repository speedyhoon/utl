package utl

import "strconv"

// Ordinal0 gives you the input number in a rank/ordinal where zero is the first position. e.g. Ordinal0(0, true) outputs "=1st"
// Ordinal0 is suitable for range loops without having to manually increment or pass i+1 to Ordinal.
//nolint:gomnd // No magic numbers.
func Ordinal0(position int, isEqual bool) (suffix string) {
	if isEqual {
		suffix = "="
	}

	position++
	suffix += strconv.Itoa(position)

	switch position % 10 {
	case 1:
		if position%100 != 11 {
			return suffix + "st"
		}
	case 2:
		if position%100 != 12 {
			return suffix + "nd"
		}
	case 3:
		if position%100 != 13 {
			return suffix + "rd"
		}
	}

	return suffix + "th"
}
