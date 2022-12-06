package iters

import (
	"github.com/skylissh/iters/models"

	"golang.org/x/exp/constraints"
)

func Max[T constraints.Ordered](iter models.Iterable[T]) T {
	return iter.Reduce(func(acc, value T) T {
		if acc > value {
			return acc
		}

		return value
	})
}
