package iters

import "github.com/skylissh/iters/models"

func AsIter[T any](values []T) *models.Iterator[T] {
	return models.NewIterator(values)
}
