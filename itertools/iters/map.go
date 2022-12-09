package iters

// NewMap returns a new iterator that maps the values of another iterator,
// using the given function.
//
// This function is only intended to be used by the top level Map method.
func NewMap[T, E any](iter Iterable[T], f func(T) E) *Map[T, E] {
	m := &Map[T, E]{iter, f, Iterator[E]{}}
	m.Iterator.iterable = m
	return m
}

// Map is an iterator that maps the values of another iterator,
// using the given function.
//
// This struct is not intended to be used directly, is created by
// the top level Map method.
type Map[T, E any] struct {
	iter Iterable[T]
	f    func(T) E

	Iterator[E]
}

// Advances the iterator and returns the next value, mapped by the given function.
//
// If there are no more values, nil is returned.
//
// # Example
//
//		iter := NewIterator([]int{1, 2, 3})
//		mapping = iter.Map(func(i int) string {
//			return strconv.Itoa(i)
//		})
//
//		assert.Equal(t, "1", *mapping.Next())
//		assert.Equal(t, "2", *mapping.Next())
//		assert.Equal(t, "3", *mappign.Next())
//	 // Once the iterator is exhausted, it will always return nil.
//		assert.Nil(t, mapping.Next())
func (m Map[T, E]) Next() *E {
	next := m.iter.Next()

	if next == nil {
		return nil
	}

	result := m.f(*next)
	return &result
}

// Returns a new iterator with the same values as the original.
//
// This means that the values of the cloned iterator are the already mapped values.
//
// # Example
//
//		iter := NewIterator([]int{1, 2, 3})
//
//		numbers = iter.Map(func(i int) string {
//			return strconv.Itoa(i)
//		})
//
//		clonned := numbers.Clone()
//
//		assert.Equal(t, *numbers.Next(), *clonned.Next())
//		assert.Equal(t, *numbers.Next(), *clonned.Next())
//		assert.Equal(t, *numbers.Next(), *clonned.Next())
//	 	// Once the iterator is exhausted, it will always return nil.
//		assert.Nil(t, numbers.Next())
//		assert.Nil(t, clonned.Next())
func (m Map[T, E]) Clone() Iterable[E] {
	return NewMap(m.iter, m.f)
}
