package models

type Take[T any] struct {
	iter Iterable[T]
	n    uint

	Base[T]
}

func (take *Take[T]) Next() *T {
	if take.n != 0 {
		take.n -= 1
		return take.iter.Next()
	}

	return nil
}

func (take *Take[T]) Collect() []T {
	return take.Base.Collect(take)
}

func (take *Take[T]) ForEach(f func(T)) {
	take.Base.ForEach(take, f)
}

func (take *Take[T]) Filter(predicate func(value T) bool) *Filter[T] {
	return take.Base.Filter(take, predicate)
}

func (take *Take[T]) Take(n uint) *Take[T] {
	return take.Base.Take(take, n)
}

func (take *Take[T]) Reduce(f func(acc T, value T) T) T {
	return take.Base.Reduce(take, f)
}

func (take *Take[T]) All(predicate func(value T) bool) bool {
	return take.Base.All(take, predicate)
}

func (take *Take[T]) Any(predicate func(value T) bool) bool {
	return take.Base.Any(take, predicate)
}
