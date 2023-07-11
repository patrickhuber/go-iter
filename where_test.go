package iter_test

import (
	"reflect"
	"testing"

	"github.com/patrickhuber/go-iter"
)

func TestWhere(t *testing.T) {
	rng := iter.Range(0, 10)
	isEven := func(i int) bool { return i%2 == 0 }
	even := iter.Where(rng, isEven)
	expected := []int{0, 2, 4, 6, 8}
	slice := iter.ToSlice(even)
	if !reflect.DeepEqual(expected, slice) {
		t.Fatalf("expected %v to equal %v", slice, expected)
	}
}
