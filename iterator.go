package iter

import "github.com/patrickhuber/go-types"

// Iterator defines a generic iterface for iterating over a sequence
type Iterator[T any] interface {
	Next() types.Option[T]
}
