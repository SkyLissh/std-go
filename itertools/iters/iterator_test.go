package iters_test

import (
	"testing"

	"github.com/skylissh/std-go/itertools/iters"
	"github.com/stretchr/testify/assert"
)

var _iter = iters.NewIter(&[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

func TestClone(t *testing.T) {
	iter := _iter.Clone()

	iter.ForEach(func(value int) {
		assert.Equal(t, value, *_iter.Next())
	})

	assert.Nil(t, _iter.Next())
}

func TestNext(t *testing.T) {
	iter := _iter.Clone()
	expect := 1

	for value := iter.Next(); value != nil; value = iter.Next() {
		assert.Equal(t, expect, *value)
		expect++
	}

	assert.Nil(t, iter.Next())
}

func TestForEach(t *testing.T) {
	iter := _iter.Clone()
	expect := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	iter.ForEach(func(value int) {
		assert.Equal(t, expect[0], value)
		expect = expect[1:]
	})
}

func TestCollect(t *testing.T) {
	iter := _iter.Clone()
	expect := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	assert.Equal(t, expect, iter.Collect())
}

func TestFilter(t *testing.T) {
	iter := _iter.Clone()
	expect := []int{2, 4, 6, 8, 10}

	assert.Equal(t, expect, iter.Filter(func(value int) bool {
		return value%2 == 0
	}).Collect())
}

func TestReduce(t *testing.T) {
	iter := _iter.Clone()
	expect := 55

	assert.Equal(t, expect, iter.Reduce(func(acc, value int) int {
		return acc + value
	}))
}

func TestTake(t *testing.T) {
	iter := _iter.Clone()
	expect := []int{1, 2, 3, 4, 5}

	assert.Equal(t, expect, iter.Take(5).Collect())
}

func TestTakeMoreThanAvailable(t *testing.T) {
	iter := _iter.Clone()
	expect := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	assert.Equal(t, expect, iter.Take(15).Collect())
}

func TestAllTrue(t *testing.T) {
	iter := _iter.Clone()

	assert.True(t, iter.All(func(value int) bool {
		return value > 0
	}))
}

func TestAllFalse(t *testing.T) {
	iter := _iter.Clone()

	assert.False(t, iter.All(func(value int) bool {
		return value > 5
	}))
}

func TestAnyTrue(t *testing.T) {
	iter := _iter.Clone()

	assert.True(t, iter.Any(func(value int) bool {
		return value > 5
	}))
}

func TestAnyFalse(t *testing.T) {
	iter := _iter.Clone()

	assert.False(t, iter.Any(func(value int) bool {
		return value > 10
	}))
}

func TestFind(t *testing.T) {
	iter := _iter.Clone()

	assert.Equal(t, 5, *iter.Find(func(num int) bool { return num == 5 }))
}
