package iter

import "github.com/patrickhuber/go-types"

func ForEach[T any](iterator Iterator[T], action func(t T)) {
	for {
		switch op := iterator.Next().(type) {
		case types.Some[T]:
			action(op.Value())
		case types.None[T]:
			return
		}
	}
}

func ForEachIndex[T any](iterator Iterator[T], action func(i int, t T)) {
	index := 0
	for {
		switch op := iterator.Next().(type) {
		case types.Some[T]:
			action(index, op.Value())
			index++
		case types.None[T]:
			return
		}
	}
}
