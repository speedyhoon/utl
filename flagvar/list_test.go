package flagvar

import (
	"github.com/speedyhoon/utl/tf"

	"testing"
)

func TestStrList_Set(t *testing.T) {
	tests := []struct {
		s    string
		want StrList
	}{
		{s: "", want: StrList{}},
		{s: " \n\t \t ", want: StrList{}},
		{s: " \n\t\v\a\b \t\v ", want: StrList{"\a\b"}},
		{s: "empty", want: StrList{"empty"}},
		{s: " empty ", want: StrList{"empty"}},
		{s: "\n\r\n\r\n\r\n\r\nempty\r \n\r\n\r", want: StrList{"empty"}},
		{s: " \n\t \t  empty\t\t ", want: StrList{"empty"}},
		{s: " \n\t\v\a\b \t\v  empty\v\t\a\b\t ", want: StrList{"\a\b \t\v  empty\v\t\a\b"}},
		{s: "empty,none", want: StrList{"empty", "none"}},
		{s: " \n\r \t  \fempty\t\f\t ,\f \n\t \r  \fnone\r\t \f", want: StrList{"empty", "none"}},
		{s: " \n\t\v\a \t\v  empty\v\t\a\t , \n\t\v\a \t\v  none\v\t\a\t ", want: StrList{"\a \t\v  empty\v\t\a", "\a \t\v  none\v\t\a"}},
	}
	for n, tt := range tests {
		tf.Run(t, n, func(t *testing.T) {
			var l StrList
			if err := l.Set(tt.s); err != nil {
				t.Errorf("Set() error = %v, expected nil", err)
			}

			for i, w := range tt.want {
				if l[i] != w {
					t.Errorf("Set()[%d]: `%s`, expected: `%s`", i, l[i], w)
				}
			}

			if len(l) != len(tt.want) {
				t.Errorf("Set() length = %d, want length %d", len(l), len(tt.want))
			}
		})
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		l    StrList
		want string
	}{
		{l: nil, want: ""},
		{l: StrList{}, want: ""},
		{l: StrList{""}, want: ""},
		{l: StrList{"empty"}, want: "empty"},
		{l: StrList{"\"empty, none\""}, want: `"empty, none"`},
		{l: StrList{"empty", "none"}, want: "empty, none"},
		{l: StrList{`"empty"`, `"none"`}, want: `"empty", "none"`},
		{l: StrList{"\n\t\v\a\b \t\v empty\n\r \t\v\t\a\b\t\n"}, want: "\n\t\v\a\b \t\v empty\n\r \t\v\t\a\b\t\n"},
		{l: StrList{
			"\n\t\v\a\b \t\v empty\n\r \t\v\t\a\b\t\n",
			"\n\t\v\a\b \t\v none\n\r \t\v\t\a\b\t\n",
		}, want: "\n\t\v\a\b \t\v empty\n\r \t\v\t\a\b\t\n, \n\t\v\a\b \t\v none\n\r \t\v\t\a\b\t\n"},
	}
	for n, tt := range tests {
		tf.Run(t, n, func(t *testing.T) {
			got := tt.l.String()
			if got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
