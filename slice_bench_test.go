package utl_test

import (
	"github.com/speedyhoon/utl"
	"strings"
	"testing"
)

func BenchmarkStrRemoveEmpty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		utl.StrRemoveEmpty([]string{"0", "", " ", "\t", "\n", "1"})
	}
}

func BenchmarkStrRemoveEmptyOld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StrRemoveEmptyOld([]string{"0", "", " ", "\t", "\n", "1"})
	}
}

func StrRemoveEmptyOld(s []string) string {
	var a []string
	for _, v := range s {
		if t := strings.TrimSpace(v); t != "" {
			a = append(a, t)
		}
	}
	return strings.Join(a, " ")
}
