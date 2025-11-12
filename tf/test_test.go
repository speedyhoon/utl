package tf_test

import (
	"fmt"
	"github.com/speedyhoon/utl/tf"
	"math/rand"
	"testing"
)

func TestRun(t *testing.T) {
	index := rand.Int()
	captureChan := make(chan string, 1)

	tf.Run(t, index, func(subT *testing.T) {
		captureChan <- subT.Name()
	})
	got := <-captureChan

	expected := fmt.Sprintf("TestRun/test[%d]", index)
	if got != expected {
		t.Errorf("Test name incorrect, got: %q, want: %q", got, expected)
	}
}
