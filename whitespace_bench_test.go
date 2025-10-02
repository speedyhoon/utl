package utl_test

import (
	"github.com/speedyhoon/utl"
	"testing"
)

// go test -bench=.*BenchmarkTrimSpace.* . -benchmem  -benchtime=30s -shuffle=on | benchtime

var out []byte

func BenchmarkTrimSpace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		out = utl.TrimSpace([]byte(" \t\x00\t \t\x00\t bird \t\x00\t \t\x00\t "))
	}
}
