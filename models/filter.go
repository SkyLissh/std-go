package models

func NewFilter[T any](iter Iterable[T], predicate func(T) bool) *Filter[T] {
	return &Filter[T]{iter, predicate, Base[T]{}}
}

type Filter[T any] struct {
	iter      Iterable[T]
	predicate func(T) bool
	Base[T]
}

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

func (filter *Filter[T]) Collect() []T {
	return filter.Base.Collect(filter)
}

func (filter *Filter[T]) ForEach(f func(T)) {
	filter.Base.ForEach(filter, f)
}

func (filter *Filter[T]) Filter(predicate func(value T) bool) *Filter[T] {
	return filter.Base.Filter(filter, predicate)
}

func (filter *Filter[T]) Take(n uint) *Take[T] {
	return filter.Base.Take(filter, n)
}

func (filter *Filter[T]) Reduce(f func(acc T, value T) T) T {
	return filter.Base.Reduce(filter, f)
}

func (filter *Filter[T]) All(predicate func(value T) bool) bool {
	return filter.Base.All(filter, predicate)
}

func (filter *Filter[T]) Any(predicate func(value T) bool) bool {
	return filter.Base.Any(filter, predicate)
}
