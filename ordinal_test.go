package utl

import (
	"math"
	"strings"
	"testing"
)

func TestOrdinal(t *testing.T) {
	var i uint16

	for ; i < math.MaxUint16; i++ {
		o := Ordinal(uint(i), false)

		if !checkSuffix(o) {
			t.Errorf("%v is an invalid ordinal", o)
		}

		if strings.HasPrefix(o, "=") {
			t.Errorf("%v should not be prefixed with `=`", o)
		}
	}
}

func TestOrdinalEquals(t *testing.T) {
	var i uint16

	for ; i < math.MaxUint16; i++ {
		o := Ordinal(uint(i), true)

		if !strings.HasPrefix(o, "=") {
			t.Errorf("%v should be prefixed with `=`", o)
		}

		if !checkSuffix(o) {
			t.Errorf("%v is an invalid ordinal", o)
		}
	}
}

func checkSuffix(ordinal string) bool {
	stuff := []string{"11th", "12th", "13th", "0th", "1st", "2nd", "3rd", "4th", "5th", "6th", "7th", "8th", "9th"}
	for _, th := range stuff {
		if strings.HasSuffix(ordinal, th) {
			return true
		}
	}
	return false
}
