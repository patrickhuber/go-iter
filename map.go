package iter

import (
	"github.com/patrickhuber/go-types"
	"github.com/patrickhuber/go-types/option"
	"github.com/patrickhuber/go-types/tuple"
)

func FromMap[TKey comparable, TValue any](m map[TKey]TValue) Iterator[types.Tuple2[TKey, TValue]] {
	return &mapIterator[TKey, TValue]{
		m: m,
	}
}

// mapIterator holds the list of keys and uses it to index into the hashmap
// I looked at using go routines but managing timeouts and etc made it more difficult
type mapIterator[TKey comparable, TValue any] struct {
	keys Iterator[TKey]
	m    map[TKey]TValue
}

func (mi *mapIterator[TKey, TValue]) Next() types.Option[types.Tuple2[TKey, TValue]] {
	// first iteration, provide the list of keys
	if mi.keys == nil {
		var keys []TKey
		for k := range mi.m {
			keys = append(keys, k)
		}
		mi.keys = FromSlice(keys)
	}
	switch next := mi.keys.Next().(type) {
	case types.Some[TKey]:
		k := next.Value()
		return option.Some(tuple.New2(k, mi.m[k]))
	default:
		return option.None[types.Tuple2[TKey, TValue]]()
	}
}
