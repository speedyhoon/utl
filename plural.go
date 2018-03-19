package utl

import "fmt"

//Plural returns an "s" if length != 1
func Plural(length int, single, multiple string) string {
	if length != 1 {
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
func Ordinal(position uint, equal bool) string {
	suffix := "th"
	switch position % 10 {
	case 1:
		if position%100 != 11 {
			suffix = "st"
		}
	case 2:
		if position%100 != 12 {
			suffix = "nd"
		}
	case 3:
		if position%100 != 13 {
			suffix = "rd"
		}
	}
	var sign string
	if equal {
		sign = "="
	}
	return fmt.Sprintf("%v%d%v", sign, position, suffix)
}
