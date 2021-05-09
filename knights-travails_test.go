package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStartIsTheTargetPosition(t *testing.T) {
	start := Square{X: 1, Y: 1}
	result := findMoves(start, start)

	assert.Equal(t, 1, len(result))
}

func TestOneMoveToTarget(t *testing.T) {
	start := Square{X: 4, Y: 5}
	target := Square{X: 5, Y: 3}
	result := findMoves(start, target)

	assert.Equal(t, 1, len(result))
}

func TestA8B7(t *testing.T) {
	start := Square{X: 1, Y: 8}
	target := Square{X: 2, Y: 7}
	result := findMoves(start, target)

	assert.Equal(t, 1, len(result))
}

func TestB1H5(t *testing.T) {
	start := Square{X: 2, Y: 1}
	target := Square{X: 8, Y: 5}
	result := findMoves(start, target)

	assert.Equal(t, 1, len(result))
}
