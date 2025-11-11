// Package tf provides testing function wrappers.
package tf

import (
	"fmt"
	"testing"
)

// Run uses "test[index]" as the test name for testing.T.Run()
func Run(t *testing.T, index int, f func(*testing.T)) {
	t.Run(fmt.Sprintf("test[%d]", index), f)
}
