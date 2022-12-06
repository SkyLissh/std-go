package iters

import "github.com/skylissh/iters/models"

func Map[T, E any](iter models.Iterable[T], f func(value T) E) *models.Map[T, E] {
	return models.NewMap(iter, f)
}
