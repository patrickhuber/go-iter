package iter

import (
	types "github.com/patrickhuber/go-types"
	"github.com/patrickhuber/go-types/option"
	"github.com/patrickhuber/go-types/tuple"
)

// Zip combines the two iterators into a single iterator of type Tuple[T1, T2].
// If the iterators are of different sizes, Zip returns an iterator of the shortest size and consumes the entire longer iterator
func Zip[T1, T2 any](iter1 Iterator[T1], iter2 Iterator[T2]) Iterator[types.Tuple2[T1, T2]] {
	return &zipIterator[T1, T2]{
		iter1: iter1,
		iter2: iter2,
	}
}

type zipIterator[T1, T2 any] struct {
	iter1 Iterator[T1]
	iter2 Iterator[T2]
}

func (zi *zipIterator[T1, T2]) Next() types.Option[types.Tuple2[T1, T2]] {
	for {
		v1, ok1 := zi.iter1.Next().Deconstruct()
		v2, ok2 := zi.iter2.Next().Deconstruct()
		switch {
		case !ok1 && !ok2:
			return option.None[types.Tuple2[T1, T2]]()
		case ok1 && ok2:
			return option.Some(tuple.New2(v1, v2))
		case ok1 && !ok2:
			continue
		case !ok1 && ok2:
			continue
		}
	}
}
