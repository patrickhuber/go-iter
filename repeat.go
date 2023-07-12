package iter

import (
	types "github.com/patrickhuber/go-types"
	"github.com/patrickhuber/go-types/option"
)

// Repeat returns an iterator that repeats the given `element` `count` times
func Repeat[T any](element T, count int) Iterator[T] {
	return &repeatIterator[T]{
		element: element,
		count:   count,
		index:   0,
	}
}

type repeatIterator[T any] struct {
	element T
	count   int
	index   int
}

// Next implements Iterator.
func (ri *repeatIterator[T]) Next() types.Option[T] {
	if ri.index >= ri.count {
		return option.None[T]()
	}
	ri.index += 1
	return option.Some(ri.element)
}
