package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPathMustIncludeStartSquare(t *testing.T) {
	start := Square{X: 1, Y: 1}
	path := newPath(start)
	assert.Equal(t, 1, path.Distance())
	assert.Equal(t, start, path.LastPosition())
}

func TestMoveMustNotModifyCurrentPath(t *testing.T) {
	start := Square{X: 1, Y: 1}
	move := Move{X: 1, Y: 1}
	path := newPath(start)
	_, isValid := path.Move(move, 8)
	assert.True(t, isValid)
	assert.Equal(t, 1, path.Distance())
	assert.Equal(t, start, path.LastPosition())
}

func TestInvalidMoveMustNotModifyCurrentPath(t *testing.T) {
	start := Square{X: 1, Y: 1}
	move := Move{X: -1, Y: -1}
	path := newPath(start)
	_, isValid := path.Move(move, 8)
	assert.False(t, isValid)
	assert.Equal(t, 1, path.Distance())
	assert.Equal(t, start, path.LastPosition())
}

func TestMoveMustCreateNewPath(t *testing.T) {
	start := Square{X: 1, Y: 1}
	move := Move{X: 1, Y: 1}
	path := newPath(start)
	newPath, isValid := path.Move(move, 8)
	assert.True(t, isValid)
	assert.Equal(t, 2, newPath.Distance())
	assert.Equal(t, Square{X: 2, Y: 2}, newPath.LastPosition())
}

func TestTargetReached(t *testing.T) {
	start := Square{X: 1, Y: 1}
	move := Move{X: 1, Y: 1}
	path := newPath(start)
	newPath, _ := path.Move(move, 8)
	targetReached := newPath.HasReachedTarget(Square{X: 2, Y: 2})
	assert.True(t, targetReached)
}

func TestTargetNotReached(t *testing.T) {
	start := Square{X: 1, Y: 1}
	move := Move{X: 1, Y: 1}
	path := newPath(start)
	newPath, _ := path.Move(move, 8)
	targetReached := newPath.HasReachedTarget(Square{X: 3, Y: 3})
	assert.False(t, targetReached)
}

