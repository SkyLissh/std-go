package iters

// NewFilter returns a new iterator that filters the values of another iterator,
// based on a predicate.
//
// The predicate is a function that returns true if the value should be included,
// and false if the value should be excluded.
//
// This function is only intended to be used by the Filter method.
func NewFilter[T any](iter Iterable[T], predicate func(T) bool) *Filter[T] {
	filter := &Filter[T]{iter, predicate, Iterator[T]{}}
	filter.Iterator.iterable = filter
	return filter
}

// Filter is an iterator that filters the values of another iterator,
// based on a predicate.
//
// The predicate is a function that returns true if the value should be included,
// and false if the value should be excluded.
//
// This struct is not intended to be used directly, is created by the Filter method.
type Filter[T any] struct {
	iter      Iterable[T]
	predicate func(T) bool
	Iterator[T]
}

// Advances the iterator and returns the next value, that matches the predicate.
//
// If there are no more values, nil is returned.
//
// # Example
//
//	iter := itertools.AsIter([]int{1, 2, 3, 4, 5})
//	filter := iter.Filter(func(v int) bool { return v%2 == 0 })
//
//	assert.Equal(t, 2, *filter.Next())
//	assert.Equal(t, 4, *filter.Next())
//	assert.Nil(t, filter.Next())
func (filter *Filter[T]) Next() *T {
	var next *T

	for v := filter.iter.Next(); v != nil; v = filter.iter.Next() {

		p := filter.predicate(*v)

		if p {
			next = v
			break
		}
	}

	return next
}

// Returns a new iterator with the same values as the original.
//
// This means that the values of the cloned iterator are also filtered.
//
// # Example
//
//	iter := itertools.AsIter([]int{1, 2, 3, 4, 5})
//	iter2 := iter.Filter(func(v int) bool { return v%2 == 0 }).Clone()
//
//	assert.Equal(t, 2, *iter2.Next())
//	assert.Equal(t, 4, *iter2.Next())
//	assert.Nil(t, iter2.Next())
func (filter *Filter[T]) Clone() Iterable[T] {
	return NewFilter(filter.iter, filter.predicate)
}
