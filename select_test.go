package iter_test

import (
	"strconv"
	"testing"

	"github.com/patrickhuber/go-iter"
)

func TestSelect(t *testing.T) {
	high := 10
	rng := iter.Range(0, high)
	strRng := iter.Select(rng, strconv.Itoa)
	expected := make([]string, 0, high)
	for i := 0; i < high; i++ {
		expected = append(expected, strconv.Itoa(i))
	}
	iter.ForEachIndex(strRng, func(index int, s string) {
		if index >= len(expected) {
			t.Fatalf("index %d exceeded bounds of expected values %d", index, len(expected))
		}
		if expected[index] != s {
			t.Fatalf("expected '%s' to equal '%s'", expected[index], s)
		}
	})
}
