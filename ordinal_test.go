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

func TestOrdinal0(t *testing.T) {
	for i := 0; i < math.MaxUint16; i++ {
		o := Ordinal0(i, false)

		if !checkSuffix0(o) {
			t.Errorf("%v is an invalid ordinal", o)
		}

		if strings.HasPrefix(o, "=") {
			t.Errorf("%v should not be prefixed with `=`", o)
		}
	}
}

func TestOrdinal0Equals(t *testing.T) {
	for i := 0; i < math.MaxUint16; i++ {
		o := Ordinal0(i, true)

		if !strings.HasPrefix(o, "=") {
			t.Errorf("%v should be prefixed with `=`", o)
		}

		if !checkSuffix0(o) {
			t.Errorf("%v is an invalid ordinal", o)
		}
	}
}

func TestOrdinal0EqualsShort(t *testing.T) {
	list := []string{"1st", "2nd", "3rd", "4th", "5th", "6th", "7th", "8th", "9th", "10th", "11th", "12th", "13th"}
	for i := range list {
		o := Ordinal0(i, false)

		if o != list[i] {
			t.Errorf("%v should be %s", o, list[i])
		}

		o = Ordinal0(i, true)

		if o != "="+list[i] {
			t.Errorf("%v should be %s", o, list[i])
		}
	}
}

func checkSuffix0(ordinal string) bool {
	stuff := []string{"11th", "12th", "13th", "0th", "1st", "2nd", "3rd", "4th", "5th", "6th", "7th", "8th", "9th"}
	for _, th := range stuff {
		if strings.HasSuffix(ordinal, th) {
			return true
		}
	}

	return false
}
