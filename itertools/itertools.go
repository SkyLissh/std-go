package itertools

import "github.com/skylissh/std-go/itertools/iters"

func AsIter[T any](values []T) *iters.Iter[T] {
	return iters.NewIter(&values)
}
