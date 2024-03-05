package utl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsGoFileName(t *testing.T) {
	tests := []struct {
		path string
		want bool
	}{
		{path: "", want: false},
		{path: ".go", want: true},
		{path: ".Go", want: true},
		{path: ".gO", want: true},
		{path: ".GO", want: true},
		{path: "foo.go", want: true},
		{path: "foo.Go", want: true},
		{path: "foo.gO", want: true},
		{path: "foo.GO", want: true},
		{path: "foo.go.go", want: true},
		{path: "foo.go.Go", want: true},
		{path: "foo.go.gO", want: true},
		{path: "foo.go.GO", want: true},
		{path: "foo.go.go", want: true},
		{path: "foo.Go.go", want: true},
		{path: "foo.gO.go", want: true},
		{path: "foo.GO.go", want: true},
		{path: "foo.go_test", want: false},
		{path: "foo.Go_test", want: false},
		{path: "foo.gO_test", want: false},
		{path: "foo.GO_test", want: false},
		{path: ".go_test", want: false},
		{path: ".Go_test", want: false},
		{path: ".gO_test", want: false},
		{path: ".GO_test", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			assert.Equalf(t, tt.want, IsGoFileName(tt.path), "IsGoFileName(%v)", tt.path)
		})
	}
}

func TestIsGoTestFileName(t *testing.T) {
	tests := []struct {
		path string
		want bool
	}{
		{path: "", want: false},
		{path: ".go", want: false},
		{path: ".Go", want: false},
		{path: ".gO", want: false},
		{path: ".GO", want: false},
		{path: "_test.go", want: true},
		{path: "_test.Go", want: true},
		{path: "_test.gO", want: true},
		{path: "_test.GO", want: true},
		{path: "._test.go", want: true},
		{path: "._test.Go", want: true},
		{path: "._test.gO", want: true},
		{path: "._test.GO", want: true},
		{path: "foo.go", want: false},
		{path: "foo.Go", want: false},
		{path: "foo.gO", want: false},
		{path: "foo.GO", want: false},
		{path: "foo_test.go", want: true},
		{path: "foo_test.Go", want: true},
		{path: "foo_test.gO", want: true},
		{path: "foo_test.GO", want: true},
		{path: "._test_foo.go", want: false},
		{path: "._test_foo.Go", want: false},
		{path: "._test_foo.gO", want: false},
		{path: "._test_foo.GO", want: false},
		{path: "foo.go_test.go", want: true},
		{path: "foo.go_test.Go", want: true},
		{path: "foo.go_test.gO", want: true},
		{path: "foo.go_test.GO", want: true},
		{path: "foo.go_test.go", want: true},
		{path: "foo.Go_test.go", want: true},
		{path: "foo.gO_test.go", want: true},
		{path: "foo.GO_test.go", want: true},
		{path: "foo.go_test", want: false},
		{path: "foo.Go_test", want: false},
		{path: "foo.gO_test", want: false},
		{path: "foo.GO_test", want: false},
		{path: ".go_test", want: false},
		{path: ".Go_test", want: false},
		{path: ".gO_test", want: false},
		{path: ".GO_test", want: false},
		{path: ".go_test.go", want: true},
		{path: ".Go_test.go", want: true},
		{path: ".gO_test.go", want: true},
		{path: ".GO_test.go", want: true},
	}
	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			assert.Equalf(t, tt.want, IsGoTestFileName(tt.path), "IsGoTestFileName(%v)", tt.path)
		})
	}
}
