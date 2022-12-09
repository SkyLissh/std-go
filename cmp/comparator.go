package cmp

type Comparator[T any] struct {
	fns   []func(T, T) int
	value *T
}

func (c *Comparator[T]) Is(value T) *Comparator[T] {
	return &Comparator[T]{c.fns, &value}
}

func (c *Comparator[T]) Equal(value T) bool {
	return c.compare(*c.value, value) == 0
}

func (c *Comparator[T]) NotEqual(value T) bool {
	return c.compare(*c.value, value) != 0
}

func (c *Comparator[T]) Less(value T) bool {
	return c.compare(*c.value, value) < 0
}

func (c *Comparator[T]) LessEqual(value T) bool {
	return c.compare(*c.value, value) <= 0
}

func (c *Comparator[T]) Greater(value T) bool {
	return c.compare(*c.value, value) > 0
}

func (c *Comparator[T]) GreaterEqual(value T) bool {
	return c.compare(*c.value, value) >= 0
}

func (c *Comparator[T]) compare(value, other T) int {
	for _, fn := range c.fns {
		if result := fn(value, other); result != 0 {
			return result
		}
	}

	return 0
}
func (c *Comparator[T]) add(fn func(T, T) int) *Comparator[T] {
	c.fns = append(c.fns, fn)

	return c
}
