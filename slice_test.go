package utl_test

import (
	"github.com/speedyhoon/utl"
	"github.com/speedyhoon/utl/tf"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrRemoveEmpty(t *testing.T) {
	tests := []struct {
		input    []string
		expected string
	}{
		{nil, ""},
		{[]string{}, ""},
		{[]string{" "}, ""},
		{[]string{"\t"}, ""},
		{[]string{"\n"}, ""},
		{[]string{"0"}, "0"},
		{[]string{"a", ""}, "a"},
		{[]string{"a", " "}, "a"},
		{[]string{"a", "\t"}, "a"},
		{[]string{"a", "\n"}, "a"},
		{[]string{"", "a"}, "a"},
		{[]string{" ", "a"}, "a"},
		{[]string{"\t", "a"}, "a"},
		{[]string{"\n", "a"}, "a"},
		{[]string{"0", "", " ", "\t", "\n", "1"}, "0 1"},
	}
	for i, tt := range tests {
		tf.Run(t, i, func(t *testing.T) {
			assert.Equal(t, tt.expected, utl.StrRemoveEmpty(tt.input))
		})
	}
}

func TestRemoveItem(t *testing.T) {
	tests := []struct {
		slice    []string
		n        uint
		expected []string
	}{
		{nil, 0, nil},
		{[]string{}, 0, []string{}},
		{[]string{"a"}, 0, []string{}},
		{[]string{"a"}, 1, []string{"a"}},
		{[]string{"a"}, 55, []string{"a"}},
		{[]string{"a", "b"}, 0, []string{"b"}},
		{[]string{"a", "b"}, 1, []string{"a"}},
		{[]string{"a", "b"}, 2, []string{"a", "b"}},
		{[]string{"0", "1", "2", "3", "4", "5"}, 0, []string{"1", "2", "3", "4", "5"}},
		{[]string{"0", "1", "2", "3", "4", "5"}, 1, []string{"0", "2", "3", "4", "5"}},
		{[]string{"0", "1", "2", "3", "4", "5"}, 2, []string{"0", "1", "3", "4", "5"}},
		{[]string{"0", "1", "2", "3", "4", "5"}, 3, []string{"0", "1", "2", "4", "5"}},
		{[]string{"0", "1", "2", "3", "4", "5"}, 4, []string{"0", "1", "2", "3", "5"}},
		{[]string{"0", "1", "2", "3", "4", "5"}, 5, []string{"0", "1", "2", "3", "4"}},
		{[]string{"0", "1", "2", "3", "4", "5"}, 6, []string{"0", "1", "2", "3", "4", "5"}},
		{[]string{"0", "1", "2", "3", "4", "5"}, 125, []string{"0", "1", "2", "3", "4", "5"}},
	}
	for i, tt := range tests {
		tf.Run(t, i, func(t *testing.T) {
			assert.Equal(t, tt.expected, utl.RemoveItem(tt.slice, tt.n))
		})
	}
}

func TestFromEnd(t *testing.T) {
	tests := []struct {
		input        []string
		indexFromEnd uint
		want         string
	}{
		{nil, 3, ""},
		{[]string{}, 3, ""},
		{[]string{"0"}, 0, "0"},
		{[]string{"0"}, 1, ""},
		{[]string{"0"}, 99, ""},
		{[]string{"0", "1"}, 0, "1"},
		{[]string{"0", "1"}, 1, "0"},
		{[]string{"0", "1"}, 99, ""},
		{[]string{"0", "1", "2"}, 0, "2"},
		{[]string{"0", "1", "2"}, 1, "1"},
		{[]string{"0", "1", "2"}, 2, "0"},
		{[]string{"0", "1", "2"}, 99, ""},
		{[]string{"0", "1", "2", "3", "4", "5", "6"}, 0, "6"},
		{[]string{"0", "1", "2", "3", "4", "5", "6"}, 1, "5"},
		{[]string{"0", "1", "2", "3", "4", "5", "6"}, 2, "4"},
		{[]string{"0", "1", "2", "3", "4", "5", "6"}, 3, "3"},
		{[]string{"0", "1", "2", "3", "4", "5", "6"}, 4, "2"},
		{[]string{"0", "1", "2", "3", "4", "5", "6"}, 5, "1"},
		{[]string{"0", "1", "2", "3", "4", "5", "6"}, 6, "0"},
		{[]string{"0", "1", "2", "3", "4", "5", "6"}, 99, ""},
	}
	for i, tt := range tests {
		tf.Run(t, i, func(t *testing.T) {
			assert.Equal(t, tt.want, utl.FromEnd(tt.input, tt.indexFromEnd))
		})
	}
}
