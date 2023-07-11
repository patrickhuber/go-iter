package iter_test

import (
	"reflect"
	"testing"

	"github.com/patrickhuber/go-iter"
)

func TestSlice(t *testing.T) {
	t.Run("FromSlice", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5, 6, 10}
		it := iter.FromSlice(slice)
		count := 0
		iter.ForEachIndex(it, func(i int, val int) {
			if i >= len(slice) {
				t.Fatalf("%d out of bounds %d", i, len(slice))
			}
			if val != slice[i] {
				t.Fatalf("%d does not equal %d", val, slice[i])
			}
			count++
		})
		if count != len(slice) {
			t.Fatalf("expected count %d to equal %d", count, len(slice))
		}
	})
	t.Run("ToSlice", func(t *testing.T) {
		rng := iter.Range(0, 10)
		expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		actual := iter.ToSlice(rng)
		if !reflect.DeepEqual(actual, expected) {
			t.Fatalf("expected %v to equal %v", actual, expected)
		}
	})
}
