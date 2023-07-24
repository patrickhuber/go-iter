package iter

import "context"

// Async returns a channel for the given iterator over the context.
// If the context is canceled, the Async function will return immediately.
func Async[T any](it Iterator[T], cx context.Context) chan T {
	ch := make(chan T)
	go func(it Iterator[T], ch chan T, cx context.Context) {
		defer close(ch)
		for o := it.Next(); o.IsSome(); o = it.Next() {
			select {
			case ch <- o.Unwrap():
			case <-cx.Done():
				return
			}
		}
	}(it, ch, cx)
	return ch
}
