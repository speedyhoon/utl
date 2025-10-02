package utl

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

	ord += UtoA(position)

	const first, second, third, ten, eleventh, twelfth, thirteenth, hundred = 1, 2, 3, 10, 11, 12, 13, 100

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
