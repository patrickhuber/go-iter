package iter

func Count[T any](iterator Iterator[T]) int {
	count := 0
	ForEach(iterator, func(t T) {
		count++
	})
	return count
}
