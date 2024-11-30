package utl

import "strconv"

// Ordinal0 gives you the input number in a rank/ordinal where zero is the first position. e.g. Ordinal0(0, true) outputs "=1st"
// Ordinal0 is suitable for range loops without having to manually increment or pass i+1 to Ordinal.
func Ordinal0(position int, isEqual bool) (suffix string) {
	if isEqual {
		suffix = "="
	}

	position++
	suffix += strconv.Itoa(position)

	const first, second, third, eleventh, twelfth, thirteenth, ten, hundred = 1, 2, 3, 11, 12, 13, 10, 100

	switch position % ten {
	case first:
		if position%hundred != eleventh {
			return suffix + "st"
		}
	case second:
		if position%hundred != twelfth {
			return suffix + "nd"
		}
	case third:
		if position%hundred != thirteenth {
			return suffix + "rd"
		}
	}

	return suffix + "th"
}
