package cmp

// Comparator is a helper to compare values.
//
// This is the return type of the Is method., and is only intended to be used
// with the Is method.
type Comparator[T any] struct {
	fns   []func(T, T) int
	value *T
}

// By returns a new Comparator[T] that compares values using the provided
// functions, but is initialized with nil as the value to compare. To set the
// value to compare, use the Is method.
//
// # Example
//
//		// Compare two strings ignoring the case.
//		lowers := cmp.By(func(value, other string) int {
//			return strings.Compare(strings.ToLower(value), strings.ToLower(other))
//		})
//
//	 assert.True(t, lowers.Is("FOO").Equal("foo"))
func (c *Comparator[T]) Is(value T) *Comparator[T] {
	return &Comparator[T]{c.fns, &value}
}

// Equal returns true if the value is equal to the value to compare.
//
// # Example
//
//	assert.True(t, cmp.Is(1).Equal(1))
func (c *Comparator[T]) Equal(value T) bool {
	return c.compare(*c.value, value) == 0
}

// NotEqual returns true if the value is not equal to the value to compare.
//
// # Example
//
//	assert.True(t, cmp.Is(1).NotEqual(2))
func (c *Comparator[T]) NotEqual(value T) bool {
	return c.compare(*c.value, value) != 0
}

// Less returns true if the value is less than the value to compare.
//
// # Example
//
//	assert.True(t, cmp.Is(1).Less(2))
func (c *Comparator[T]) Less(value T) bool {
	return c.compare(*c.value, value) < 0
}

// LessEqual returns true if the value is less than or equal to the value to
// compare.
//
// # Example
//
//	assert.True(t, cmp.Is(1).LessEqual(1))
//	assert.True(t, cmp.Is(1).LessEqual(2))
func (c *Comparator[T]) LessEqual(value T) bool {
	return c.compare(*c.value, value) <= 0
}

// Greater returns true if the value is greater than the value to compare.
//
// # Example
//
//	assert.True(t, cmp.Is(2).Greater(1))
func (c *Comparator[T]) Greater(value T) bool {
	return c.compare(*c.value, value) > 0
}

// GreaterEqual returns true if the value is greater than or equal to the value
// to compare.
//
// # Example
//
//	assert.True(t, cmp.Is(1).GreaterEqual(1))
//	assert.True(t, cmp.Is(2).GreaterEqual(1))
func (c *Comparator[T]) GreaterEqual(value T) bool {
	return c.compare(*c.value, value) >= 0
}

// This method is used to compare the value to the value to compare.
//
// It will call each of the comparison functions in the order they were added
// to the Comparator[T], and return the result of the first comparison function
// that returns a non-zero value. If all comparison functions return 0, then 0
// is returned.
func (c *Comparator[T]) compare(value, other T) int {
	for _, fn := range c.fns {
		if result := fn(value, other); result != 0 {
			return result
		}
	}

	return 0
}

// This method is used to add a comparison function to the Comparator[T].
//
// It is only intended to be used by the By method.
func (c *Comparator[T]) add(fn func(T, T) int) *Comparator[T] {
	c.fns = append(c.fns, fn)

	return c
}
