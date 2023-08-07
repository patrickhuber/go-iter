package iter

import (
	"github.com/patrickhuber/go-types"
	"github.com/patrickhuber/go-types/option"
	"golang.org/x/exp/constraints"
)

type RangeOption int

// RangeWith emulates the '[' '(' and ')' ']' syntax used to show a range is inclusive or exclusive
func RangeWith[T constraints.Integer](includeBegin bool, begin T, end T, includeEnd bool) Iterator[T] {
	return &rangeIterator[T]{
		begin:          begin,
		beginInclusive: includeBegin,
		end:            end,
		endInclusive:   includeEnd,
		current:        option.None[T](),
	}
}

func Range[T constraints.Integer](begin T, end T) Iterator[T] {
	return RangeWith[T](true, begin, end, false)
}

type rangeIterator[T constraints.Integer] struct {
	begin          T
	beginInclusive bool
	end            T
	endInclusive   bool
	current        types.Option[T]
}

func (ri *rangeIterator[T]) Next() types.Option[T] {
	var end T = ri.end
	if ri.endInclusive {
		end += 1
	}

	// currently nil?
	if ri.current.IsNone() {
		var begin T = ri.begin
		if !ri.beginInclusive {
			begin += 1
		}
		ri.current = option.Some(begin)

		if begin >= end {
			return option.None[T]()
		}
		return ri.current
	}

	// currently past end?
	some := ri.current.Unwrap()

	// next past end?
	if some+1 >= end {
		return option.None[T]()
	}

	// return next
	ri.current = option.Some(some + 1)
	return ri.current
}
