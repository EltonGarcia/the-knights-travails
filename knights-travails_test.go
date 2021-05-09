package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStartIsTheTargetPosition(t *testing.T) {
	start := Square{X: 1, Y: 1}
	result := findMoves(start, start)

	assert.Equal(t, start, result.Moves[0])
}

func TestOneMoveToTarget(t *testing.T) {
	start := Square{X: 4, Y: 5}
	target := Square{X: 5, Y: 3}
	result := findMoves(start, target)

	assert.Equal(t, start, result.Moves[0])
	assert.Equal(t, target, result.Moves[1])
}

func TestA8B7(t *testing.T) {
	start := Square{X: 1, Y: 8}
	target := Square{X: 2, Y: 7}
	result := findMoves(start, target)

	assert.Equal(t, start, result.Moves[0])
	assert.Equal(t, Square{X: 2, Y: 6}, result.Moves[1])
	assert.Equal(t, Square{X: 1, Y: 4}, result.Moves[2])
	assert.Equal(t, Square{X: 3, Y: 5}, result.Moves[3])
	assert.Equal(t, target, result.Moves[4])
}

func TestB1H5(t *testing.T) {
	start := Square{X: 2, Y: 1}
	target := Square{X: 8, Y: 5}
	result := findMoves(start, target)

	assert.Equal(t, start, result.Moves[0])
	assert.Equal(t, Square{X: 3, Y: 3}, result.Moves[1])
	assert.Equal(t, Square{X: 5, Y: 2}, result.Moves[2])
	assert.Equal(t, Square{X: 7, Y: 3}, result.Moves[3])
	assert.Equal(t, target, result.Moves[4])
}
