package iter_test

import (
	"testing"

	"github.com/patrickhuber/go-iter"
)

func TestRange(t *testing.T) {
	t.Run("0..10", func(t *testing.T) {
		rng := iter.Range(0, 10)
		count := 0
		iter.ForEachIndex(rng, func(index int, i int) {
			if count != index {
				t.Fatalf("expected %d to equal %d", count, index)
			}
			if count != i {
				t.Fatalf("expected %d to equal %d", count, i)
			}
			count++
		})
	})
}
