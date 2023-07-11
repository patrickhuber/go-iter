package iter_test

import (
	"testing"

	"github.com/patrickhuber/go-iter"
)

func TestCount(t *testing.T) {
	expected := 10
	rng := iter.Range(0, expected)
	actual := iter.Count(rng)
	if actual != expected {
		t.Fatalf("expected count of %d but found %d", expected, actual)
	}
}
