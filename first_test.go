package iter_test

import (
	"testing"

	"github.com/patrickhuber/go-iter"
	"github.com/patrickhuber/go-types"
)

func TestFirst(t *testing.T) {
	sequence := iter.Range(15, 30)
	first := iter.First(sequence)
	switch op := first.(type) {
	case types.Some[int]:
		if op.Value != 15 {
			t.Fatalf("expected 15 but found %d", op.Value)
		}
	case types.None[int]:
		t.Fatalf("no first item")
	default:
		t.Fatalf("item is not an int")
	}
}

func TestFirstWhere(t *testing.T) {
	sequence := iter.Range(20, 25)
	first := iter.FirstWhere(sequence, func(i int) bool { return i%3 == 0 })
	if first.IsNone() {
		t.Fatalf("expected 21, found none")
	}
	if first.Unwrap() != 21 {
		t.Fatalf("expected 21, found %d", first.Unwrap())
	}
}
