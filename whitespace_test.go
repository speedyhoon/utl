package utl

import (
	"github.com/go-openapi/testify/v2/assert"
	"github.com/speedyhoon/tf"
	"testing"
)

func TestTrimSpace(t *testing.T) {
	tests := []struct {
		input []byte
		want  []byte
	}{
		{
			input: nil,
			want:  nil,
		}, {
			input: []byte(""),
			want:  []byte(""),
		}, {
			input: []byte("bird"),
			want:  []byte("bird"),
		}, {
			input: []byte(" \t\x00\t \t\x00\t bird \t\x00\t \t\x00\t "),
			want:  []byte("bird"),
		},
	}
	for i, tt := range tests {
		tf.Run(t, i, func(t *testing.T) {
			assert.Equalf(t, tt.want, TrimSpace(tt.input), "TrimSpace(%v)", tt.input)
		})
	}
}
