package resourceid

import (
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
)

func TestNext(t *testing.T) {
	t.Parallel()
	for i := 0; i < 10000; i++ {
		t.Run(fmt.Sprintf("next-%d", i), func(t *testing.T) {})
		got := Next()
		assert.Assert(t, got != "")
	}
}
