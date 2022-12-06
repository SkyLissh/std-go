package models

func NewMap[T, E any](iter Iterable[T], f func(T) E) *Map[T, E] {
	return &Map[T, E]{iter, f, Base[E]{}}
}

type Map[T, E any] struct {
	iter Iterable[T]
	f    func(T) E

	Base[E]
}

func (m Map[T, E]) Next() *E {
	next := m.iter.Next()

	if next == nil {
		return nil
	}

	result := m.f(*next)
	return &result
}

func (iter *Map[T, E]) Collect() []E {
	return iter.Base.Collect(iter)
}

func (iter *Map[T, E]) ForEach(f func(E)) {
	iter.Base.ForEach(iter, f)
}

func (iter *Map[T, E]) Filter(predicate func(value E) bool) *Filter[E] {
	return iter.Base.Filter(iter, predicate)
}

func (iter *Map[T, E]) Take(n uint) *Take[E] {
	return iter.Base.Take(iter, n)
}

func (iter *Map[T, E]) Reduce(f func(acc E, value E) E) E {
	return iter.Base.Reduce(iter, f)
}

func (iter *Map[T, E]) All(predicate func(value E) bool) bool {
	return iter.Base.All(iter, predicate)
}

func (iter *Map[T, E]) Any(predicate func(value E) bool) bool {
	return iter.Base.Any(iter, predicate)
}
