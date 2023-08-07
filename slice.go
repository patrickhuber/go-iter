package iter

import (
	"github.com/patrickhuber/go-types"
	"github.com/patrickhuber/go-types/option"
)

func FromSlice[T any](slice []T) Iterator[T] {
	return &sliceIterator[T]{
		slice: slice,
		index: 0,
	}
}

type sliceIterator[T any] struct {
	slice []T
	index int
}

func (i *sliceIterator[T]) Next() types.Option[T] {
	if i.index >= len(i.slice) {
		return option.None[T]()
	}
	ret := option.Some(i.slice[i.index])
	i.index++
	return ret
}

func ToSlice[T any](iterator Iterator[T]) []T {
	var slice []T
	for {
		switch op := iterator.Next().(type) {
		case types.Some[T]:
			slice = append(slice, op.Value)
		case types.None[T]:
			return slice
		}
	}
}
