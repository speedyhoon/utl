package utl

import "strconv"

// Plural returns an "s" if length != 1
func Plural(length int, single, multiple string) string {
	if length != 1 && length != -1 {
		if multiple != "" {
			return multiple
		}
		return "s"
	}
	if single != "" {
		return single
	}
	return ""
}

//Ordinal gives you the input number in a rank/ordinal format. e.g. Ordinal(3, true) outputs "=3rd"
//nolint:gomnd // No magic numbers.
func Ordinal(position uint, isEqual bool) (suffix string) {
	if isEqual {
		suffix = "="
	}

	suffix += strconv.FormatUint(uint64(position), 10)

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
