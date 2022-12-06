package models

func NewIterator[T any](values []T) *Iterator[T] {
	return &Iterator[T]{values, -1, Base[T]{}}
}

type Iterator[T any] struct {
	values  []T
	current int

	Base[T]
}

func (iter *Iterator[T]) Next() *T {
	iter.current++

	if iter.current >= len(iter.values) {
		return nil
	}

	return &iter.values[iter.current]
}

func (iter *Iterator[T]) Collect() []T {
	return iter.Base.Collect(iter)
}

func (iter *Iterator[T]) ForEach(f func(T)) {
	iter.Base.ForEach(iter, f)
}

func (iter *Iterator[T]) Filter(predicate func(value T) bool) *Filter[T] {
	return iter.Base.Filter(iter, predicate)
}

func (iter *Iterator[T]) Take(n uint) *Take[T] {
	return iter.Base.Take(iter, n)
}

func (iter *Iterator[T]) Reduce(f func(acc T, value T) T) T {
	return iter.Base.Reduce(iter, f)
}

func (iter *Iterator[T]) All(predicate func(value T) bool) bool {
	return iter.Base.All(iter, predicate)
}

func (iter *Iterator[T]) Any(predicate func(value T) bool) bool {
	return iter.Base.Any(iter, predicate)
}
