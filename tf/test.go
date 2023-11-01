package tf

import (
	"fmt"
	"testing"
)

// Run uses "test[n]" as the test name for testing.T.Run()
func Run(t *testing.T, n int, f func(*testing.T)) {
	t.Run(fmt.Sprintf("test[%d]", n), f)
}
