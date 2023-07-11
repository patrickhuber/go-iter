package iter

import "github.com/patrickhuber/go-types"

type Iterator[T any] interface {
	Next() types.Option[T]
}
