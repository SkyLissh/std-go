package models

type Iterable[T any] interface {
	Next() *T
	ForEach(func(T))
	Reduce(func(acc T, value T) T) T
	Filter(func(T) bool) *Filter[T]
	Take(uint) *Take[T]
}
