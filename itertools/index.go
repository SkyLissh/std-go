package itertools

import "github.com/skylissh/std-go/itertools/iters"

func Index[T comparable](iter iters.Iterable[T], item T) int {
	index := -1

	for v := iter.Next(); v != nil; v = iter.Next() {
		index++

		if *v == item {
			return index
		}
	}

	return index
}
