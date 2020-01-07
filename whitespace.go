package utl

import (
	"strings"
	"unicode"
)

//StripWhiteSpace removes unicode whitespace characters and condenses multiple spaces to a single character
func StripWhiteSpace(str string) string {
	return CollapseSpaces(strings.Map(func(r rune) rune {
		//credit: Tim Cooper //stackoverflow.com/questions/32081808/strip-all-whitespace-from-a-string
		if !unicode.IsPrint(r) || unicode.IsSpace(r) && r != ' ' {
			return -1
		}
		return r
	}, str))
}

func CollapseSpaces(str string) string {
	for i := 0; i != len(str); {
		i = len(str)
		//Replace double spaces with a single space
		str = strings.Replace(str, "  ", " ", -1)
	}
	return str
}
