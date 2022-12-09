package cmp

import (
	"bytes"
	"strings"
	"time"

	"golang.org/x/exp/constraints"
)

var (
	// DefaultFns is the default set of functions used to compare values.
	CompareInts    = By(compareType[int])
	CompareStrings = By(strings.Compare)
	CompareBytes   = By(bytes.Compare)
	CompareFloats  = By(compareType[float64])
	CompareRune    = By(compareType[rune])
	CompareUint    = By(compareType[uint])
	CompareBools   = By(func(value, other bool) int {
		switch {
		case value == other:
			return 0
		case value:
			return 1
		default:
			return -1

		}
	})
	CompareTime = By(func(value, other time.Time) int {
		switch {
		case value.Equal(other):
			return 0
		case value.After(other):
			return 1
		default:
			return -1
		}
	})
)

func compareType[T constraints.Ordered](value, other T) int {
	switch {
	case value < other:
		return -1
	case value > other:
		return 1
	default:
		return 0
	}
}
