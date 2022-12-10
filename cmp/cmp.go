// This package provides a set of functions to compare values, and a common
// interface to compare values of any type.
//
// I created this package because of the lack of a common interface to compare, or
// a way to overload the comparison operator. This package provides a common
// interface to compare values of any type, and a set of functions to compare
// values of common types.
//
// For be able to use the functions of this package your type must implement the
// Comparable[T] interface.
//
// # Example
//
// The following example shows how to create a type that implements the
// Comparable[T] interface.
//
//	type Person struct {
//		Name string
//		Age  int
//	}
//
//	func (p Person) Compare(other Person) int {
//		if p.Age < other.Age {
//			return -1
//		}
//
//		if p.Age > other.Age {
//			return 1
//		}
//
//		return strings.Compare(p.Name, other.Name)
//	}
package cmp

import (
	"time"

	"golang.org/x/exp/constraints"
)

// By returns a new Comparator[T] that uses the provided functions to compare
// values. The functions are executed in the order they are provided. If the
// first function returns a value other than 0, the comparison is complete and
// the value is returned. If the first function returns 0, the second function
// is executed and so on. If all functions return 0, the values are considered
// equal.
//
// # Example
//
// The following example shows how to create a Comparator[T] that compares
// slices by their length.
//
//		cmp := cmp.By(func(a, b []int) int {
//	 	   return len(a) - len(b)
//		})
//
//		assert.Equal(t, cmp.Is([]int{1, 2, 3}), []int{1, 2, 3})
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

// Is returns a new Comparator[T] that compares values using the provided
// value. The value must implement the Comparable[T] interface or be one of
// the following types: int, string, []byte, bool, float64, rune, uint,
// time.Time.
//
// This methos is a common way to compare values. It is equivalent to
// calling By with the appropriate comparison function.
//
// # Example
//
//	if cmp.Is(1).Equal(2) {
//		// ...
//	}
//
//	if cmp.Is("foo").Equal("bar") {
//		// ...
//	}
//
//	// If the current type does not implement the Comparable[T] interface,
//	// the function will panic.
//	assert.Panics(t, func() {
//		cmp.Is(struct{}{}).Equal(struct{}{})
//	})
//
//	// This method don't work with slices, maps for that you need to use
//	// IsSlice or IsMap.
//	assert.Panics(t, func() {
//		cmp.Is([]int{}).Equal([]int{})
//	})
func Is[T any](value T) *Comparator[T] {
	cast := any(value)

	if s, ok := cast.(Comparable[T]); ok {
		method := func(value, other T) int {
			return s.Compare(other)
		}

		return By(method).Is(value)
	}

	var cmp any
	switch cast.(type) {
	case int:
		cmp = compareInts
		return cmp.(*Comparator[T]).Is(value)
	case string:
		cmp = compareStrings
	case []byte:
		cmp = compareBytes
	case bool:
		cmp = compareBools
	case float64:
		cmp = compareFloats
	case rune:
		cmp = compareRune
	case uint:
		cmp = compareUint
	case time.Time:
		cmp = compareTime
	default:
		panic("The provided type is not comparable")
	}

	return cmp.(*Comparator[T]).Is(value)
}

// IsSlice returns a new Comparator[[]T] that compares slices using the
// provided value. The slice will be compared by length and then by each
// element.
//
// # Example
//
//	if cmp.IsSlice([]int{1, 2, 3}).Equal([]int{1, 2, 3}) {
//		// ...
//	}
func IsSlice[T constraints.Ordered](value []T) *Comparator[[]T] {
	return By(compareSlice[T]).Is(value)
}

// IsMap returns a new Comparator[map[K]V] that compares maps using the
// provided value. The map will be compared by length and then by each
// key/value pair.
//
// # Example
//
//	if cmp.IsMap(map[string]int{"foo": 1}).Equal(map[string]int{"foo": 1}) {
//		// ...
//	}
func IsMap[K constraints.Ordered, V constraints.Ordered](value map[K]V) *Comparator[map[K]V] {
	return By(compareMap[K, V]).Is(value)
}
