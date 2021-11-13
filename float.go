package utl

import (
	"strconv"
	"strings"
)

// TrimFloat converts the float64 f to a string in the shortest possible decimal or scientific format.
func TrimFloat(f float64) string {
	decimal := strings.Replace(
		strings.TrimLeft(
			strconv.FormatFloat(f, 'f', -1, 64),
			"0", // Remove all leading zeros
		),
		"-0.", "-.", 1, // Remove leading zeros when negative.
	)

	scientific := strconv.FormatFloat(f, 'E', -1, 64)

	if len(scientific) < len(decimal) {
		return scientific
	}

	return decimal
}
