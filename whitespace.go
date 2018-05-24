package utl

import (
	"strings"
	"unicode"
)

//StripWhiteSpace removes unicode whitespace characters and condenses condenses multiple spaces to a single character
func StripWhiteSpace(str string) string {
	return CollapseSpaces(strings.Map(func(r rune) rune {
		//credit: Tim Cooper //stackoverflow.com/questions/32081808/strip-all-whitespace-from-a-string
		if !unicode.IsPrint(r) || unicode.IsSpace(r) && r != ' ' {
			return -1
		}
		return r
	}, str))
}

func CollapseSpaces(a string) string {
	l := len(a)
	for {
		//Replace double spaces with a single space
		a = strings.Replace(a, "  ", " ", -1)
		if l == len(a) {
			break
		}
		l = len(a)
	}
	return a
}
