package utl_test

import (
	"fmt"
	"github.com/speedyhoon/utl"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogBaseX(t *testing.T) {
	tests := []struct {
		base float64
		x    float64
		want float64
	}{
		{base: 3, x: 3, want: 1},
		{base: 3, x: 9, want: 2.0000000000000004},
		{base: 3, x: 27, want: 3.0000000000000004},
		{base: 4, x: 4, want: 1},
		{base: 4, x: 16, want: 2},
		{base: 4, x: 64, want: 3},
		{base: 4, x: 256, want: 4},
		{base: 4, x: 1024, want: 5},
		{base: 4, x: 4096, want: 6},
		{base: 4, x: 16384, want: 7},
		{base: 10, x: 10, want: 1},
		{base: 10, x: 100, want: 2},
		{base: 10, x: 1000, want: 2.9999999999999996},
		{base: 10, x: 10000, want: 4},
		{base: 256, x: 256, want: 1},
		{base: 256, x: 256 * 256, want: 2},
		{base: 256, x: 256 * 256 * 256, want: 3},
		{base: 256, x: 256 * 256 * 256 * 256, want: 4},
	}
	for _, tt := range tests {
		t.Run(
			fmt.Sprintf("%f, %f", tt.base, tt.x),
			func(t *testing.T) {
				assert.Equalf(t, tt.want, utl.LogBaseX(tt.base, tt.x), "LogBaseX(%f, %f)", tt.base, tt.x)
			})
	}
}

func TestLogBaseXUint(t *testing.T) {
	tests := []struct {
		base uint
		x    uint
		want uint
	}{
		{base: 2, x: 2, want: 1},
		{base: 2, x: 4, want: 2},
		{base: 2, x: 8, want: 3},
		{base: 2, x: 16, want: 4},
		{base: 2, x: 32, want: 5},
		{base: 2, x: 64, want: 6},
		{base: 2, x: 128, want: 7},
		{base: 2, x: 256, want: 8},
		{base: 2, x: 512, want: 9},
		{base: 2, x: 1024, want: 10},
		{base: 2, x: 2048, want: 11},
		{base: 2, x: 4096, want: 12},
		{base: 3, x: 3, want: 1},
		{base: 3, x: 9, want: 3},
		{base: 3, x: 27, want: 4},
		{base: 3, x: 89, want: 5},
		{base: 4, x: 4, want: 1},
		{base: 4, x: 16, want: 2},
		{base: 4, x: 64, want: 3},
		{base: 4, x: 256, want: 4},
		{base: 4, x: 1024, want: 5},
		{base: 4, x: 4096, want: 6},
		{base: 4, x: 16384, want: 7},
		{base: 10, x: 10, want: 1},
		{base: 10, x: 100, want: 2},
		{base: 10, x: 1000, want: 3},
		{base: 10, x: 10000, want: 4},
		{base: 256, x: 256, want: 1},
		{base: 256, x: 256 * 256, want: 2},
		{base: 256, x: 256 * 256 * 256, want: 3},
		{base: 256, x: 256 * 256 * 256 * 256, want: 4},
		{base: 256, x: 256 * 256 * 256 * 256, want: 4},
	}
	for _, tt := range tests {
		t.Run(
			fmt.Sprintf("%d, %d", tt.base, tt.x),
			func(t *testing.T) {
				assert.Equal(t, tt.want, utl.LogBaseXUint(tt.base, tt.x), "LogBaseXUint(%d, %d)", tt.base, tt.x)
			})
	}
}
