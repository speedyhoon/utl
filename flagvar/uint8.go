package flagvar

import (
	"strconv"
)

// Uint8 implements the flag.Value interface.
type Uint8 uint8

// Set converts s into an uint8.
func (l *Uint8) Set(s string) error {
	u64, err := strconv.ParseUint(s, 10, 8)
	if err != nil {
		return err
	}

	*l = Uint8(u64)
	return nil
}

// String returns the string representation of Uint8.
func (l *Uint8) String() string {
	if l != nil {
		return strconv.FormatUint(uint64(*l), 10)
	}
	return ""
}
