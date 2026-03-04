package flagvar

import "strconv"

// Uint32 implements the flag.Value interface.
type Uint32 uint32

// Set converts string s into an uint32.
func (l *Uint32) Set(s string) error {
	u64, err := strconv.ParseUint(s, 10, 32)
	if err == nil {
		*l = Uint32(u64)
	}

	return err
}

// String returns the string representation of Uint32.
func (l *Uint32) String() string {
	if l != nil {
		return strconv.FormatUint(uint64(*l), 10)
	}
	return ""
}

func (l *Uint32) Get() uint32 {
	if l != nil {
		return uint32(*l)
	}
	return 0
}
