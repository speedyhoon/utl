package utl

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"testing"
	"time"
)

func TestNameOf(t *testing.T) {
	type test struct {
		name string
		arg  any
		want string
	}
	tests := []test{
		{name: "", arg: nil},
		{name: "empty", arg: "", want: ""},
		{name: "string", arg: "zebra", want: "zebra"},
		{name: "package name", arg: "example.com/zoo/lion", want: "lion"},
		{name: "time.Now", arg: time.Now, want: "time.Now"},
		{name: "utl.Cwd", arg: Cwd, want: "utl.Cwd"},
		{name: "utl.Exec", arg: Exec, want: "utl.Exec"},
		{name: "utl.UtoA", arg: UtoA, want: "utl.UtoA"},
		{name: "utl.Ordinal", arg: Ordinal, want: "utl.Ordinal"},
		{name: "utl.LogBaseXUint", arg: LogBaseXUint, want: "utl.LogBaseXUint"},
		{name: "strconv.FormatFloat", arg: strconv.FormatFloat, want: "strconv.FormatFloat"},
		{name: "http.AllowQuerySemicolons", arg: http.AllowQuerySemicolons, want: "http.AllowQuerySemicolons"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NameOfFunc(tt.arg), "NameOfFunc(%v)", tt.arg)
		})
	}

	// Test these inputs cause a panic.
	tests = []test{
		{name: "zero", arg: 0},
		{name: "type test struct", arg: test{}},
		{name: "type time.Time struct", arg: time.Time{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Panics(t, func() { NameOfFunc(tt.arg) }, tt.arg)
		})
	}
}
