package cmp

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInt(t *testing.T) {
	assert.True(t, Is(1).Equal(1))
	assert.True(t, Is(1).NotEqual(2))
	assert.True(t, Is(1).Greater(0))
	assert.True(t, Is(1).GreaterEqual(1))
	assert.True(t, Is(1).Less(2))
	assert.True(t, Is(1).LessEqual(1))
}

func TestString(t *testing.T) {
	assert.True(t, Is("a").Equal("a"))
	assert.True(t, Is("a").NotEqual("b"))
	assert.True(t, Is("a").Greater("A"))
	assert.True(t, Is("a").GreaterEqual("a"))
	assert.True(t, Is("a").Less("b"))
	assert.True(t, Is("a").LessEqual("a"))
}

func TestBool(t *testing.T) {
	assert.True(t, Is(true).Equal(true))
	assert.True(t, Is(true).NotEqual(false))
	assert.True(t, Is(true).Greater(false))
	assert.True(t, Is(true).GreaterEqual(true))
	assert.True(t, Is(false).Less(true))
	assert.True(t, Is(true).LessEqual(true))
}

func TestFloat(t *testing.T) {
	assert.True(t, Is(1.0).Equal(1.0))
	assert.True(t, Is(1.0).NotEqual(2.0))
	assert.True(t, Is(1.0).Greater(0.0))
	assert.True(t, Is(1.0).GreaterEqual(1.0))
	assert.True(t, Is(1.0).Less(2.0))
	assert.True(t, Is(1.0).LessEqual(1.0))
}

func TestUInt(t *testing.T) {
	assert.True(t, Is(uint(1)).Equal(uint(1)))
	assert.True(t, Is(uint(1)).NotEqual(uint(2)))
	assert.True(t, Is(uint(1)).Greater(uint(0)))
	assert.True(t, Is(uint(1)).GreaterEqual(uint(1)))
	assert.True(t, Is(uint(1)).Less(uint(2)))
	assert.True(t, Is(uint(1)).LessEqual(uint(1)))
}

func TestTime(t *testing.T) {
	assert.True(t, Is(time.Now()).Equal(time.Now()))
	assert.True(t, Is(time.Now()).NotEqual(time.Now().Add(time.Second)))
	assert.True(t, Is(time.Now()).Greater(time.Now().Add(-time.Second)))
	assert.True(t, Is(time.Now()).GreaterEqual(time.Now()))
	assert.True(t, Is(time.Now()).Less(time.Now().Add(time.Second)))
	assert.True(t, Is(time.Now()).LessEqual(time.Now()))
}

func TestByte(t *testing.T) {
	assert.True(t, Is([]byte{1}).Equal([]byte{1}))
	assert.True(t, Is([]byte{1}).NotEqual([]byte{2}))
	assert.True(t, Is([]byte{1}).Greater([]byte{0}))
	assert.True(t, Is([]byte{1}).GreaterEqual([]byte{1}))
	assert.True(t, Is([]byte{1}).Less([]byte{2}))
	assert.True(t, Is([]byte{1}).LessEqual([]byte{1}))
}
