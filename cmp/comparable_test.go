package cmp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIs(t *testing.T) {
	assert.True(t, Is(1).Equal(1))
	assert.True(t, Is(1).Greater(0))
	assert.True(t, Is(1).Less(2))
}
