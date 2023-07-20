package iter_test

import (
	"testing"

	"github.com/patrickhuber/go-iter"
	"github.com/patrickhuber/go-types"
)

func TestZip(t *testing.T) {
	t.Run("equal_length", func(t *testing.T) {
		first := []int{1, 2, 3, 4, 5}
		second := []string{"0", "2", "4", "6", "8"}
		result := iter.Zip(iter.FromSlice(first), iter.FromSlice(second))
		iter.ForEachIndex(result, func(index int, item types.Tuple2[int, string]) {
			if first[index] != item.Value1() {
				t.Fatalf("expected %d to equal %d", first[index], item.Value1())
			}
			if second[index] != item.Value2() {
				t.Fatalf("expected %s to equal %s", second[index], item.Value2())
			}
		})
	})
}
