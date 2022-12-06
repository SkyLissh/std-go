package iters

import "github.com/skylissh/iters/models"

func Find[T comparable](iter models.Iterable[T], value T) *T {
	return iter.Filter(func(v T) bool {
		return value == v
	}).Next()
}
