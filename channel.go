package iter

import (
	"context"

	types "github.com/patrickhuber/go-types"
	"github.com/patrickhuber/go-types/option"
)

type ChannelOption[T any] func(*channelIterator[T])

func WithContext[T any](cx context.Context) ChannelOption[T] {
	return func(ci *channelIterator[T]) {
		ci.cx = cx
	}
}

func FromChannel[T any](ch chan T, options ...ChannelOption[T]) Iterator[T] {
	ci := &channelIterator[T]{
		ch: ch,
	}
	for _, option := range options {
		option(ci)
	}
	if ci.cx == nil {
		ci.cx = context.Background()
	}
	return ci
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
