package iters

// NewTake returns a new iterator that takes the first n values of another iterator.
//
// This function is only intended to be used by the Take method.
func NewTake[T any](iter Iterable[T], n uint) *Take[T] {
	take := &Take[T]{iter, n, Iterator[T]{}}
	take.Iterator.iterable = take
	return take
}

// Take is an iterator that takes the first n values of another iterator.
//
// This struct is not intended to be used directly, is created by the Take method.
type Take[T any] struct {
	iter Iterable[T]
	n    uint

	Iterator[T]
}

// Advances the iterator and returns the next value.
//
// If there are no more values, nil is returned.
//
// # Example
//
//	iter := NewIterator([]int{1, 2, 3, 4, 5})
//	iter = iter.Take(3)
//
//	assert.Equal(t, 1, *iter.Next())
//	assert.Equal(t, 2, *iter.Next())
//	assert.Equal(t, 3, *iter.Next())
//	assert.Nil(t, iter.Next())
func (take *Take[T]) Next() *T {
	if take.n != 0 {
		take.n -= 1
		return take.iter.Next()
	}

	return nil
}

// Returns a new iterator with the same values as the original.
//
// This means that the values of the cloned iterator are the already taken values.
//
// # Example
//
//	iter := NewIterator([]int{1, 2, 3, 4, 5})
//	iter = iter.Take(3)
//
//	assert.Equal(t, 1, *iter.Next())
//	assert.Equal(t, 2, *iter.Next())
//	assert.Equal(t, 3, *iter.Next())
//	assert.Nil(t, iter.Next())
//
//	// Clone the Take iterator with the already taken values.
//	iter = iter.Clone()
//
//	assert.Equal(t, 1, *iter.Next())
//	assert.Equal(t, 2, *iter.Next())
//	assert.Equal(t, 3, *iter.Next())
//	assert.Nil(t, iter.Next())
func (take *Take[T]) Clone() Iterable[T] {
	return NewTake(take.iter, take.n)
}
