package iter_test

import (
	"strconv"
	"testing"

	"github.com/patrickhuber/go-iter"
	"github.com/patrickhuber/go-types"
	"github.com/patrickhuber/go-types/tuple"
)

func TestMap(t *testing.T) {
	t.Run("FromMap", func(t *testing.T) {
		m := map[string]int{}
		expected := []types.Tuple2[string, int]{}
		for i := 0; i < 4; i++ {
			key := strconv.Itoa(i)
			m[key] = i
			expected = append(expected, tuple.New2(key, i))
		}
		it := iter.FromMap(m)
		actual := iter.ToSlice(it)
		if len(actual) != len(expected) {
			t.Fatalf("expected len(actual) %d to equal len(expected) %d", len(actual), len(expected))
		}
		for k := range m {
			delete(m, k)
		}
		for i := 0; i < len(expected); i++ {
			k, v := expected[i].Deconstruct()
			m[k] = v
		}
		for i := 0; i < len(actual); i++ {
			k, a := actual[i].Deconstruct()
			e, ok := m[k]
			if !ok {
				t.Fatalf("key %s not found", k)
			}
			if a != e {
				t.Fatalf("expected value %d for key %s to be %d", a, k, e)
			}
		}
	})
}
