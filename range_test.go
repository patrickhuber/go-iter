package iter_test

import (
	"testing"

	"github.com/patrickhuber/go-iter"
)

func TestRange(t *testing.T) {
	run := func(t *testing.T, includeStart bool, start, end int, includeEnd bool) {
		var rng iter.Iterator[int]
		if includeStart && !includeEnd {
			rng = iter.Range(start, end)
		} else {
			rng = iter.RangeWith(includeStart, start, end, includeEnd)
		}
		index := start
		if !includeStart {
			index += 1
		}
		for op := rng.Next(); op.IsSome(); op = rng.Next() {
			if op.Unwrap() != index {
				t.Fatalf("expected %d to equal %d", index, op.Unwrap())
			}
			index++
		}
		index -= 1
		if index != end && includeEnd {
			t.Fatalf("expected %d to equal %d", index, end)
		}
		if index != end-1 && !includeEnd {
			t.Fatalf("expected %d to equal %d", index, end-1)
		}
	}
	type test struct {
		name         string
		includeStart bool
		start        int
		end          int
		includeEnd   bool
	}
	tests := []test{
		{"[0..10)", true, 0, 10, false},
		{"[15..30)", true, 15, 30, false},
		{"[15..30]", true, 15, 30, true},
		{"(15..30)", false, 15, 30, false},
		{"(15..30]", false, 15, 30, true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			run(t, test.includeStart, test.start, test.end, test.includeEnd)
		})
	}
}
