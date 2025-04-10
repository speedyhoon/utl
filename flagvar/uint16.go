package flagvar

import (
	"strconv"
)

// Uint16 implements the flag.Value interface.
type Uint16 uint16

// Set converts s into an uint16.
func (l *Uint16) Set(s string) error {
	u64, err := strconv.ParseUint(s, 10, 16)
	if err != nil {
		return err
	}

	*l = Uint16(u64)
	return nil
}

// String returns the string representation of Uint16.
func (l *Uint16) String() string {
	if l != nil {
		return strconv.FormatUint(uint64(*l), 10)
	}
	return ""
}
