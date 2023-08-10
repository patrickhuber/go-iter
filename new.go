package iter

// New creates a new iterator from the list of items
func New[T any](items ...T) Iterator[T] {
	return FromSlice(items)
}
