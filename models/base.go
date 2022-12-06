package models

type Base[T any] struct{}

func (b *Base[T]) ForEach(iter Iterable[T], f func(T)) {
	for v := iter.Next(); v != nil; v = iter.Next() {
		f(*v)
	}
}

func (b *Base[T]) Collect(iter Iterable[T]) []T {
	var collect []T

	iter.ForEach(func(value T) {
		collect = append(collect, value)
	})

	return collect
}

func (b *Base[T]) Reduce(iter Iterable[T], f func(acc T, value T) T) T {
	acc := *iter.Next()

	iter.ForEach(func(value T) {
		acc = f(acc, value)
	})

	return acc
}

func (b *Base[T]) Filter(iter Iterable[T], predicate func(T) bool) *Filter[T] {
	return &Filter[T]{iter, predicate, Base[T]{}}
}

func (b *Base[T]) Take(iter Iterable[T], n uint) *Take[T] {
	return &Take[T]{iter, n, Base[T]{}}
}

func (b *Base[T]) All(iter Iterable[T], predicate func(T) bool) bool {
	for v := iter.Next(); v != nil; v = iter.Next() {
		if !predicate(*v) {
			return false
		}
	}

	return true
}

func (b *Base[T]) Any(iter Iterable[T], predicate func(T) bool) bool {
	return iter.Filter(predicate).Next() != nil
}
