package iter

import (
	"github.com/patrickhuber/go-types"
	"github.com/patrickhuber/go-types/option"
)

func Select[TSource, TTarget any](iterator Iterator[TSource], transform func(TSource) TTarget) Iterator[TTarget] {
	return &transformIterator[TSource, TTarget]{
		iterator:  iterator,
		transform: transform,
	}
}

type transformIterator[TSource, TTarget any] struct {
	iterator  Iterator[TSource]
	transform func(TSource) TTarget
}

func (i *transformIterator[TSource, TTarget]) Next() types.Option[TTarget] {
	switch op := i.iterator.Next().(type) {
	case types.Some[TSource]:
		return option.Some(i.transform(op.Value))
	default:
		return option.None[TTarget]()
	}
}
