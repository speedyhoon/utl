package flagvar

import "strings"

// Sep is the string separator used in StrList.Set.
var Sep = ","

// StrList implements the flag.Value interface.
type StrList []string

// Set slices s into all substrings separated by Sep and returns a slice of
// the substrings between those separators with leading & trailing whitespace removed.
func (l *StrList) Set(s string) error {
	for _, item := range strings.Split(s, Sep) {
		item = strings.TrimSpace(item)
		if item != "" {
			*l = append(*l, item)
		}
	}

	return nil
}

// String returns the string representation of the StrList, joining its elements with commas & a space.
func (l *StrList) String() string {
	return strings.Join(*l, ", ")
}
