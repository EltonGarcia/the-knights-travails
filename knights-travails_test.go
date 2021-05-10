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
	assert.Equal(t, "E3", formatPath(result[0]))
}

func TestA8B7(t *testing.T) {
	start := Square{X: 1, Y: 8}
	target := Square{X: 2, Y: 7}
	result := findMoves(start, target)

	assert.Equal(t, 10, len(result))
	assert.Equal(t, "B6 A4 C5 B7", formatPath(result[0]))
	assert.Equal(t, "B6 C4 A5 B7", formatPath(result[1]))
	assert.Equal(t, "B6 C4 D6 B7", formatPath(result[2]))
	assert.Equal(t, "B6 C8 D6 B7", formatPath(result[3]))
	assert.Equal(t, "B6 D7 C5 B7", formatPath(result[4]))
	assert.Equal(t, "C7 A6 C5 B7", formatPath(result[5]))
	assert.Equal(t, "C7 B5 D6 B7", formatPath(result[6]))
	assert.Equal(t, "C7 E6 C5 B7", formatPath(result[7]))
	assert.Equal(t, "C7 E6 D8 B7", formatPath(result[8]))
	assert.Equal(t, "C7 E8 D6 B7", formatPath(result[9]))
}

func TestA8A7(t *testing.T) {
	start := Square{X: 1, Y: 8}
	target := Square{X: 1, Y: 7}
	result := findMoves(start, target)

	assert.Equal(t, 2, len(result))
	assert.Equal(t, "B6 C8 A7", formatPath(result[0]))
	assert.Equal(t, "C7 B5 A7", formatPath(result[1]))
}

func TestB1H5(t *testing.T) {
	start := Square{X: 2, Y: 1}
	target := Square{X: 8, Y: 5}
	result := findMoves(start, target)

	assert.Equal(t, 9, len(result))
	assert.Equal(t, "C3 E2 F4 H5", formatPath(result[0]))
	assert.Equal(t, "C3 E2 G3 H5", formatPath(result[1]))
	assert.Equal(t, "C3 D5 F4 H5", formatPath(result[2]))
	assert.Equal(t, "C3 D5 F6 H5", formatPath(result[3]))
	assert.Equal(t, "C3 E4 G3 H5", formatPath(result[4]))
	assert.Equal(t, "C3 E4 F6 H5", formatPath(result[5]))
	assert.Equal(t, "D2 F1 G3 H5", formatPath(result[6]))
	assert.Equal(t, "D2 E4 G3 H5", formatPath(result[7]))
	assert.Equal(t, "D2 E4 F6 H5", formatPath(result[8]))
}

func TestA1H8(t *testing.T) {
	start := Square{X: 1, Y: 1}
	target := Square{X: 8, Y: 8}
	result := findMoves(start, target)

	assert.Equal(t, 108, len(result))
}