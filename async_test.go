package iter_test

import (
	"testing"

	"github.com/patrickhuber/go-iter"
)

func TestAsync(t *testing.T) {
	t.Run("async", func(t *testing.T) {
		max := 10
		rng := iter.Range(0, max)
		sum := 0
		for o := range iter.Async(rng) {
			sum += o
		}
		if sum != 45 {
			t.Fatalf("expected sum to be 45 but found %d", sum)
		}
	})
}
