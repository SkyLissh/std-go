package itertools

import (
	"github.com/skylissh/std-go/itertools/iters"
)

func Map[T, E any](iter iters.Iterable[T], f func(value T) E) *iters.Map[T, E] {
	return iters.NewMap(iter, f)
}
