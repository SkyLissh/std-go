package cmp

// Comparable is the interface that must be implemented by any type that
// is to be compared using the Is method.
type Comparable[T any] interface {
	// Compares this object with the specified object for order.
	//
	// Returns a negative integer, zero, or a positive integer as
	// this object is less than, equal to, or greater than the specified object.
	Compare(other T) int
}
