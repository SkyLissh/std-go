package iters

import "github.com/skylissh/iters/models"

func Index[T comparable](iter models.Iterable[T], item T) int {
	index := -1

	for v := iter.Next(); v != nil; v = iter.Next() {
		index++

		if *v == item {
			return index
		}
	}

	return index
}
