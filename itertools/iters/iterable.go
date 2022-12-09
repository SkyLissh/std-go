package iters

// Iterable is an interface that describes an iterator,
// The iterator is a data structure that allows you to iterate over a collection, lazily.
//
// You can implement this interface to create your own iterator.
type Iterable[T any] interface {
	// Returns the next value of the iterator, and advances the iterator.
	//
	// If the iterator is empty, it returns nil.
	Next() *T
}

type Cloneable[T any] interface {
	// Return a new iterator with the same values as the original iterator.
	//
	// The new iterator must be independent of the original iterator.
	Clone() *T
}

type CloneableIter[T, E any] interface {
	Iterable[T]
	Cloneable[E]
}
