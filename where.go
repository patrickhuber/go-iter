package iter

import "github.com/patrickhuber/go-types"

func Where[T any](iterator Iterator[T], where func(t T) bool) Iterator[T] {
	return &whereIterator[T]{
		iterator: iterator,
		where:    where,
	}
}

type whereIterator[T any] struct {
	iterator Iterator[T]
	where    func(t T) bool
}

func (i *whereIterator[T]) Next() types.Option[T] {
	for {
		switch op := i.iterator.Next().(type) {
		case types.Some[T]:
			if i.where(op.Value()) {
				return op
			}
		case types.None[T]:
			return op
		}
	}
}
