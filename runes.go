package iter

import "context"

func Runes(str string, options ...ChannelOption[rune]) Iterator[rune] {

	co := &channelOption[rune]{}
	for _, option := range options {
		option(co)
	}
	if co.cx == nil {
		co.cx = context.Background()
	}

	ch := make(chan rune)
	go func(ch chan rune, cx context.Context) {
		defer close(ch)
		for _, r := range str {
			select {
			case ch <- r:
			case <-cx.Done():
				return
			}
		}
	}(ch, co.cx)
	return FromChannel(ch, options...)
}
