package iter

import (
	"context"

	"github.com/patrickhuber/go-types"
	"github.com/patrickhuber/go-types/option"
	"github.com/patrickhuber/go-types/tuple"
)

// FromMapAsync returns a call to FromChannel with a channel created by iterating over the map in a go routine
// The context passed in is used to pass the WithContext chanel option to the FromChannel call
func FromMapAsync[TKey comparable, TValue any](m map[TKey]TValue, cx context.Context) Iterator[types.Tuple2[TKey, TValue]] {
	ch := make(chan types.Tuple2[TKey, TValue])
	go func(ch chan types.Tuple2[TKey, TValue], cx context.Context) {
		defer close(ch)
		for k, v := range m {
			select {
			case ch <- tuple.New2(k, v):
			case <-cx.Done():
				return
			}
		}
	}(ch, cx)
	return FromChannel(ch, WithContext[types.Tuple2[TKey, TValue]](cx))
}

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
		// size to the map but set len to zero
		keys := make([]TKey, 0, len(mi.m))
		for k := range mi.m {
			keys = append(keys, k)
		}
		mi.keys = FromSlice(keys)
	}

	none := option.None[types.Tuple2[TKey, TValue]]()
	for {
		switch next := mi.keys.Next().(type) {
		case types.Some[TKey]:
			k := next.Value()
			v, ok := mi.m[k]
			if !ok {
				continue
			}
			return option.Some(tuple.New2(k, v))
		case types.None[TKey]:
			return none
		}
	}
}
