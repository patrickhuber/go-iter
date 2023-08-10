package iter_test

import (
	"testing"

	"github.com/patrickhuber/go-iter"
	types "github.com/patrickhuber/go-types"
)

func TestNew(t *testing.T) {
	evens := iter.New(0, 2, 4, 6, 8, 10)
	for i := 0; i <= 10; i += 2 {
		switch next := evens.Next().(type) {
		case types.Some[int]:
			if next.Value != i {
				t.Fatalf("expected %d found %d", i, next.Value)
			}
		case types.None[int]:
			t.Fatalf("expected Some[int]")
		}
	}
	if !evens.Next().IsNone() {
		t.Fatalf("expected iterator to be consumed")
	}
}
