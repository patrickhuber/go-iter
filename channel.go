package iter

import (
	"context"

	types "github.com/patrickhuber/go-types"
	"github.com/patrickhuber/go-types/option"
)

type ChannelOption[T any] func(*channelOption[T])

type channelOption[T any] struct {
	cx context.Context
}

// WithContext provides an context.Context for channel operations
func WithContext[T any](cx context.Context) ChannelOption[T] {
	return func(ci *channelOption[T]) {
		ci.cx = cx
	}
}

func FromChannel[T any](ch chan T, options ...ChannelOption[T]) Iterator[T] {
	co := &channelOption[T]{}
	for _, option := range options {
		option(co)
	}
	if co.cx == nil {
		co.cx = context.Background()
	}
	return &channelIterator[T]{
		ch: ch,
		cx: co.cx,
	}
}

type channelIterator[T any] struct {
	ch chan T
	cx context.Context
}

// Next implements Iterator.
func (ci *channelIterator[T]) Next() types.Option[T] {
	select {
	case v, ok := <-ci.ch:
		return option.New(v, ok)
	case <-ci.cx.Done():
		return option.None[T]()
	}
}
