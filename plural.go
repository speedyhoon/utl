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

// Ordinal gives you the input number in a rank/ordinal format. e.g. Ordinal(3, true) outputs "=3rd"
func Ordinal(position uint, isEqual bool) (ord string) {
	if isEqual {
		ord = "="
	}

	ord += strconv.FormatUint(uint64(position), 10)

	const first, second, third, eleventh, twelfth, thirteenth, ten, hundred = 1, 2, 3, 11, 12, 13, 10, 100

	switch position % ten {
	case first:
		if position%hundred != eleventh {
			return ord + "st"
		}
	case second:
		if position%hundred != twelfth {
			return ord + "nd"
		}
	case third:
		if position%hundred != thirteenth {
			return ord + "rd"
		}
	}

	return ord + "th"
}
