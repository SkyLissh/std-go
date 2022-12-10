package cmp

import (
	"bytes"
	"strings"
	"time"

	"golang.org/x/exp/constraints"
)

// Here are the default comparison functions for the supported types.
var (
	compareInts    = By(compareType[int])
	compareStrings = By(strings.Compare)
	compareBytes   = By(bytes.Compare)
	compareFloats  = By(compareType[float64])
	compareRune    = By(compareType[rune])
	compareUint    = By(compareType[uint])
	compareBools   = By(func(value, other bool) int {
		switch {
		case value == other:
			return 0
		case value:
			return 1
		default:
			return -1

		}
	})
	compareTime = By(func(value, other time.Time) int {
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

func compareSlice[T constraints.Ordered](value, other []T) int {
	if len(value) < len(other) {
		return -1
	}

	if len(value) > len(other) {
		return 1
	}

	for i := 0; i < len(value); i++ {
		if Is(value[i]).Greater(other[i]) {
			return 1
		}

		if Is(value[i]).Less(other[i]) {
			return -1
		}
	}

	return 0
}

func compareMap[T constraints.Ordered, U constraints.Ordered](value, other map[T]U) int {
	if len(value) < len(other) {
		return -1
	}

	if len(value) > len(other) {
		return 1
	}

	for k, v := range value {
		if _, ok := other[k]; !ok {
			return -1
		}

		if Is(v).Greater(other[k]) {
			return 1
		}

		if Is(v).Less(other[k]) {
			return -1
		}
	}

	return 0
}

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
