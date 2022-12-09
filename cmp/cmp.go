package cmp

import "time"

func By[T any](fns ...func(T, T) int) *Comparator[T] {
	if len(fns) == 0 {
		panic("At least one function is required")
	}

	cmp := Comparator[T]{make([]func(T, T) int, 0, len(fns)), nil}
	for _, fn := range fns {
		if fn == nil {
			panic("The provided function is nil")
		}

		cmp.add(fn)
	}

	return &cmp
}

func Is[T any](value T) *Comparator[T] {
	cast := any(value)

	t, ok := cast.(Comparable[T])
	if ok {
		method := func(value, other T) int {
			return t.Compare(other)
		}

		return By(method).Is(value)
	}

	var cmp any
	switch cast.(type) {
	case int:
		cmp = CompareInts
	case string:
		cmp = CompareStrings
	case []byte:
		cmp = CompareBytes
	case bool:
		cmp = CompareBools
	case float64:
		cmp = CompareFloats
	case rune:
		cmp = CompareRune
	case uint:
		cmp = CompareUint
	case time.Time:
		cmp = CompareTime
	default:
		panic("The provided value is not comparable")
	}

	return cmp.(*Comparator[T]).Is(value)
}
