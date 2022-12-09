package iters

func NewCycle[T any, E CloneableIter[T, E]](iter E) *Cycle[T, E] {
	cycle := &Cycle[T, E]{*iter.Clone(), iter, Iterator[T]{}}
	cycle.Iterator.iterable = cycle
	return cycle
}

// Repeats an iterator indefinitely.
//
// This struct is only used internally by the Cycle method from iterators.
type Cycle[T any, E CloneableIter[T, E]] struct {
	orig E
	iter Iterable[T]

	Iterator[T]
}

// Returns a new iterator that repeats the original iterator indefinitely.
// Differs from other iterators in that returns nil when the iterator is exhausted.
//
// # Example
//
//		iter := tertools.AsIter([]int{1, 2, 3})
//		cycle := iter.Cycle()
//
//		assert.Equal(t, 1, *cycle.Next())
//		assert.Equal(t, 2, *cycle.Next())
//		assert.Equal(t, 3, *cycle.Next())
//		assert.Equal(t, 1, *cycle.Next())
//		assert.Equal(t, 2, *cycle.Next())
//		assert.Equal(t, 3, *cycle.Next())
//		assert.Equal(t, 1, *cycle.Next())
//	 // ...
func (cycle *Cycle[T, E]) Next() *T {
	if v := cycle.iter.Next(); v != nil {
		return v
	}

	cycle.iter = *cycle.orig.Clone()
	return cycle.iter.Next()
}

// Returns a new iterator with the same values as the original.
// This will be also endless.
//
// # Example
//
//		iter := itertools.AsIter([]int{1, 2, 3})
//
//		cycle := iter.Cycle()
//		clonned := cycle.Clone()
//
//		assert.Equal(t, *cycle.Next(), *clonned.Next())
//		assert.Equal(t, *cycle.Next(), *clonned.Next())
//		assert.Equal(t, *cycle.Next(), *clonned.Next())
//		assert.Equal(t, *cycle.Next(), *clonned.Next())
//		assert.Equal(t, *cycle.Next(), *clonned.Next())
//	 // Endless, you get the point...
func (cycle *Cycle[T, E]) Clone() *Cycle[T, E] {
	return NewCycle[T](cycle.orig)
}
