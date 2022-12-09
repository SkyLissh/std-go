package itertools

import "github.com/skylissh/std-go/itertools/iters"

// Repeats the values from the original iterator indefinitely.
//
// # Example
//
//		iter := itertools.AsIter([]int{1, 2, 3})
//		cycle := itertools.Cycle(iter)
//
//		assert.Equal(t, 1, iter.Next())
//		assert.Equal(t, 2, iter.Next())
//		assert.Equal(t, 3, iter.Next())
//		assert.Equal(t, 1, iter.Next())
//		assert.Equal(t, 2, iter.Next())
//		assert.Equal(t, 3, iter.Next())
//		assert.Equal(t, 1, iter.Next())
//	 // Endless loop...
func Cycle[T any, E iters.CloneableIter[T, E]](iter E) *iters.Cycle[T, E] {
	cycle := iters.NewCycle[T](iter)
	return cycle
}
