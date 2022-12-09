package iters

// Iterator is the base struct for all iterators.
//
// You can use it to create your own iterators, only remember to implement
// the Next method.
type Iterator[T any] struct {
	iterable Iterable[T]
}

// Next returns the next value of the iterator, and advances the iterator.
//
// If the iterator is empty, it returns nil.
func (iter *Iterator[T]) Next() *T {
	return iter.iterable.Next()
}

// Iteratates over the iterator and calls the function f for each value.
//
// # Example
//
//	iter := itertools.AsIter([]int{1, 2, 3, 4, 5})
//	iter.ForEach(func(value int) {
//		fmt.Print("%d ", value)
//	})
//
//	// Output:
//	// 1 2 3 4 5
func (iter *Iterator[T]) ForEach(f func(T)) {
	for v := iter.Next(); v != nil; v = iter.Next() {
		f(*v)
	}
}

// Reduce the iterator to a single value, applying the function f to each value.
//
// The result of the function f is used as the accumulator for the next iteration.
//
// # Example
//
//	iter := itertools.AsIter([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
//	sum := iter.Reduce(func(acc int, value int) int {
//		return acc + value
//	})
//
//	assert.Equal(t, 55, sum)
func (iter *Iterator[T]) Reduce(f func(acc T, value T) T) T {
	acc := *iter.Next()

	iter.ForEach(func(value T) {
		acc = f(acc, value)
	})

	return acc
}

// Returns a new iterator that contains the values of the original iterator
// that match the predicate.
//
// # Example
//
//	iter := itertools.AsIter([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
//	iter = iter.Filter(func(value int) bool {
//		return value%2 == 0
//	}).Collect()
//
//	assert.Equal(t, []int{2, 4, 6, 8, 10}, iter)
func (iter *Iterator[T]) Filter(predicate func(T) bool) *Filter[T] {
	return NewFilter(iter.iterable, predicate)
}

// Returns a new iterator that yields n values of the original iterator, or
// fewer if the original iterator is exhausted.
//
// # Examples
//
// Take 5 values:
//
//	iter := itertools.AsIter([]int{1, 2, 3})
//	iter = iter.Take(2)
//
//	assert.Equal(t, 1, iter.Next())
//	assert.Equal(t, 2, iter.Next())
//	assert.Nil(t, iter.Next())
//
// Take more values than the original iterator has:
//
//	iter := itertools.AsIter([]int{1, 2, 3})
//	iter = iter.Take(10)
//
//	assert.Equal(t, 1, iter.Next())
//	assert.Equal(t, 2, iter.Next())
//	assert.Equal(t, 3, iter.Next())
//	assert.Nil(t, iter.Next())
func (iter *Iterator[T]) Take(n uint) *Take[T] {
	return NewTake(iter.iterable, n)
}

// Check if all values of the iterator match the predicate.
//
// Returns false at the first value that does not match the predicate,
// so it is not necessary to iterate over the entire iterator.
//
// # Example
//
//	iter := itertools.AsIter([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
//	allEven := iter.All(func(value int) bool {
//		return value%2 == 0
//	})
//
//	assert.False(t, allEven)
//
//	// We can continue to use the iterator after the call to All:
//	assert.Equal(t, 3, iter.Next())
func (iter *Iterator[T]) All(predicate func(T) bool) bool {
	for v := iter.Next(); v != nil; v = iter.Next() {
		if !predicate(*v) {
			return false
		}
	}

	return true
}

// Check if any value of the iterator matches the predicate.
//
// Returns true at the first value that matches the predicate,
// so it is not necessary to iterate over the entire iterator.
//
// # Example
//
//		iter := itertools.AsIter([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
//		anyEven := iter.Any(func(value int) bool {
//			return value%2 == 0
//		})
//
//	 assert.True(t, anyEven)
//
//		// We can continue to use the iterator after the call to Any:
//		assert.Equal(t, 3, iter.Next())
func (iter *Iterator[T]) Any(predicate func(T) bool) bool {
	return iter.Filter(predicate).Next() != nil
}

// Seearches for a value in the iterator that matches the predicate.
//
// Returns the first value that matches the predicate, or nil if no value
// matches the predicate.
//
// # Example
//
//	iter := itertools.AsIter([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
//	even := iter.Find(func(value int) bool {
//		return value%2 == 0
//	})
//
//	assert.Equal(t, 2, even)
//
// iter.Find(predicate) is equivalent to iter.Filter(predicate).Next()
func (iter *Iterator[T]) Find(predicate func(T) bool) *T {
	return iter.Filter(predicate).Next()
}

// Collect the values of the iterator into a slice.
//
// # Example
//
//	iter := itertools.AsIter([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
//	evens := iter.Filter(func(value int) bool {
//		return value%2 == 0
//	}).Collect()
//
//	assert.Equal(t, []int{2, 4, 6, 8, 10}, evens)
func (iter *Iterator[T]) Collect() []T {
	collect := make([]T, 0)

	iter.ForEach(func(value T) {
		collect = append(collect, value)
	})

	return collect
}
