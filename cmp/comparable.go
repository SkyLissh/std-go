package cmp

// Comparable is the common interface implemented by all types that can be
// compared.
//
// You need to implement this interface if you want to use the Is method.
type Comparable[T any] interface {
	// Compares this object with the specified object for order.
	//
	// Returns a negative integer, zero, or a positive integer as
	// this object is less than, equal to, or greater than the specified object.
	Compare(other T) int
}
