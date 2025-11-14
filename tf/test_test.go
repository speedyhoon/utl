package tf_test

import (
	"fmt"
	"github.com/speedyhoon/utl/tf"
	"math/rand"
	"testing"
)

func TestRun(t *testing.T) {
	index := rand.Int()
	nameChan := make(chan string, 1)

	tf.Run(t, index, func(tt *testing.T) {
		nameChan <- tt.Name()
	})
	got := <-nameChan

	expected := fmt.Sprintf("TestRun/test[%d]", index)
	if got != expected {
		t.Errorf("Test name incorrect, got: %q, want: %q", got, expected)
	}
}
