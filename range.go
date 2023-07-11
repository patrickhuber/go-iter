package iter

import (
	"github.com/patrickhuber/go-types"
	"github.com/patrickhuber/go-types/option"
	"golang.org/x/exp/constraints"
)

func Range[T constraints.Integer](begin T, end T) Iterator[T] {
	return &rangeIterator[T]{
		begin: begin,
		end:   end,
	}
}

type rangeIterator[T constraints.Integer] struct {
	begin   T
	end     T
	current T
}

func (ri *rangeIterator[T]) Next() types.Option[T] {
	if ri.current >= ri.end {
		return option.None[T]()
	}
	res := option.Some(ri.current)
	ri.current += 1
	return res
}
