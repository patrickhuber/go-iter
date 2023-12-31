package iter_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/patrickhuber/go-iter"
)

func TestChannel(t *testing.T) {

	t.Run("basic", func(t *testing.T) {
		ch := make(chan int)
		go func(c chan int) {
			defer close(c)
			for i := 0; i < 10; i++ {
				ch <- i
			}
		}(ch)
		it := iter.FromChannel(ch)
		iter.ForEach(it, func(i int) {
			fmt.Println(i)
		})
	})

	t.Run("cancel", func(t *testing.T) {
		ch := make(chan int)
		cx, cancel := context.WithCancel(context.Background())
		go func(c chan int, cx context.Context) {
			defer close(c)
			for i := 0; i < 10; i++ {
				select {
				case ch <- i:
				case <-cx.Done():
					return
				}
			}
		}(ch, cx)

		cancel()
		it := iter.FromChannel(ch, iter.WithContext[int](cx))
		iter.ForEach(it, func(i int) {
			fmt.Println(i)
		})
		if cx.Err() == nil {
			t.Fatalf("expected context to err")
		}
	})
}
