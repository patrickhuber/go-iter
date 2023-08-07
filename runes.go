package iter

import (
	"unicode/utf8"

	types "github.com/patrickhuber/go-types"
	"github.com/patrickhuber/go-types/option"
)

type runeIterator struct {
	str      string
	position int
}

func (i *runeIterator) Next() types.Option[rune] {
	if i.position >= len(i.str) {
		return option.None[rune]()
	}
	v, width := utf8.DecodeRuneInString(i.str[i.position:])
	i.position += width
	return option.Some(v)
}

func Runes(str string) Iterator[rune] {
	return &runeIterator{
		str:      str,
		position: 0,
	}
}
