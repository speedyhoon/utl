package utl

import (
	"github.com/speedyhoon/utl/tf"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDel(t *testing.T) {
	type testCase struct {
		name  string
		index uint
		have  []int
		want  []int
	}
	tests := []testCase{
		{name: "nil_0", index: 0, have: nil, want: nil},
		{name: "nil_1", index: 1, have: nil, want: nil},
		{name: "empty_0", index: 0, have: []int{}, want: []int{}},
		{name: "empty_1", index: 1, have: []int{}, want: []int{}},
		{name: "one_0", index: 0, have: []int{1}, want: []int{}},
		{name: "one_1", index: 1, have: []int{1}, want: []int{1}},
		{name: "two_0", index: 0, have: []int{1, 2}, want: []int{2}},
		{name: "two_1", index: 1, have: []int{1, 2}, want: []int{1}},
		{name: "two_2", index: 2, have: []int{1, 2}, want: []int{1, 2}},
		{name: "three_0", index: 0, have: []int{1, 2, 3}, want: []int{2, 3}},
		{name: "three_1", index: 1, have: []int{1, 2, 3}, want: []int{1, 3}},
		{name: "three_2", index: 2, have: []int{1, 2, 3}, want: []int{1, 2}},
		{name: "three_3", index: 3, have: []int{1, 2, 3}, want: []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Del(&tt.have, tt.index)
			assert.Equalf(t, tt.want, tt.have, "Del(%v)", tt.have)
		})
	}
}

func TestDel_string(t *testing.T) {
	zero := make([]string, 0)
	one := []string{"one"}
	two := []string{"one", "two"}

	tests := []struct {
		s     []string
		index uint
		want  []string
	}{
		{s: nil, index: 0, want: nil},
		{s: nil, index: 1, want: nil},

		{s: zero, index: 0, want: zero},
		{s: zero, index: 1, want: zero},

		{s: one, index: 0, want: zero},
		{s: one, index: 1, want: one},

		{s: two, index: 0, want: []string{"two"}},
		{s: two, index: 1, want: one},
		{s: two, index: 2, want: two},
	}
	for i, tt := range tests {
		tf.Run(t, i, func(t *testing.T) {
			Del(&tt.s, tt.index)
			assert.Equal(t, tt.want, tt.s)
		})
	}
}

func TestDelDup(t *testing.T) {
	type testCase struct {
		name string
		have []int
		want []int
	}
	tests := []testCase{
		{name: "nil", have: nil, want: nil},
		{name: "empty", have: []int{}, want: []int{}},
		{name: "one", have: []int{1}, want: []int{1}},
		{name: "two", have: []int{1, 2}, want: []int{1, 2}},
		{name: "twos", have: []int{1, 1}, want: []int{1}},
		{name: "six1", have: []int{1, 1, 1, 1, 1, 1}, want: []int{1}},
		{name: "six1,2", have: []int{1, 2, 2, 1, 2, 2}, want: []int{1, 2}},
		{name: "six1,2,4", have: []int{1, 2, 2, 1, 4, 2}, want: []int{1, 2, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DelDup(&tt.have)
			assert.Equalf(t, tt.want, tt.have, "DelDup(%v)", tt.have)
		})
	}
}

func TestLen(t *testing.T) {
	type testCase struct {
		name string
		have []int
		want uint
	}
	tests := []testCase{
		{want: 0, have: nil, name: "nil"},
		{want: 0, have: []int{}, name: "empty"},
		{want: 1, have: []int{1}, name: "one"},
		{want: 2, have: []int{1, 2}, name: "two"},
		{want: 255, have: make([]int, 255), name: "255"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Len(tt.have), "Len(%v)", tt.have)
		})
	}
}
