package iter

import (
	types "github.com/patrickhuber/go-types"
	"github.com/patrickhuber/go-types/option"
)

func First[T any](i Iterator[T]) types.Option[T] {
	return i.Next()
}

func FirstWhere[T any](i Iterator[T], condition func(t T) bool) types.Option[T] {
	for op := i.Next(); op.IsSome(); op = i.Next() {
		if condition(op.Unwrap()) {
			return op
		}
	}
	return option.None[T]()
}
