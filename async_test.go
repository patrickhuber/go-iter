package iter_test

import (
	"context"
	"testing"

	"github.com/patrickhuber/go-iter"
)

func TestAsync(t *testing.T) {
	max := 10
	rng := iter.Range(0, max)
	cx := context.Background()
	sum := 0
	for o := range iter.Async(rng, cx) {
		sum += o
	}
	if sum != 45 {
		t.Fatalf("expected sum to be 45 but found %d", sum)
	}
}
