package flagvar

import "strconv"

// Uint32 implements the flag.Value interface.
type Uint32 uint32

// Set converts s into an uint32.
func (l *Uint32) Set(s string) error {
	u64, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return err
	}

	*l = Uint32(u64)
	return nil
}

// String returns the string representation of Uint32.
func (l *Uint32) String() string {
	if l != nil {
		return strconv.FormatUint(uint64(*l), 10)
	}
	return ""
}
