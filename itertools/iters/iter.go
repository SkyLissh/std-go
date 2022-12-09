package iters

func NewIter[T any](values *[]T) *Iter[T] {
	iter := &Iter[T]{values, -1, Iterator[T]{}}
	iter.Iterator.iterable = iter
	return iter
}

type Iter[T any] struct {
	values  *[]T
	current int

	Iterator[T]
}

// Advances the iterator and returns the next value.
//
// You need to be careful when using this method, because it can return nil when
// the iterator is empty.
//
// # Example
//
//	iter := itertools.AsIter([]int{1, 2})
//
//	assert.Equal(t, 1, *iter.Next())
//	assert.Equal(t, 2, *iter.Next())
//	// The iterator is empty, so it returns nil
//	assert.Nil(t, iter.Next())
func (iter *Iter[T]) Next() *T {
	iter.current++

	if iter.current >= len(*iter.values) {
		return nil
	}

	return &(*iter.values)[iter.current]
}

// Returns a new iterator with the same values as the original.
//
// # Example
//
//	iter := itertools.AsIter([]int{1, 2})
//	iter2 := iter.Clone()
//
//	assert.Equal(t, iter.Next(), iter2.Next())
func (iter *Iter[T]) Clone() *Iter[T] {
	return NewIter(iter.values)
}
