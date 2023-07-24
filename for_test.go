package iter_test

import (
	"testing"

	"github.com/patrickhuber/go-iter"
)

func TestFor(t *testing.T) {
	count := 0
	max := 10
	rng := iter.Range(0, max)
	for op := rng.Next(); op.IsSome(); op = rng.Next() {
		count++
	}
	if count != max {
		t.Fatalf("expected %d to equal %d", count, max)
	}

}
