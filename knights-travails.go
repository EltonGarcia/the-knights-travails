package main

import "math"

// Constant for a standard board size.
const boardSize = 8

type Square struct {
	X int
	Y int
}

type Move struct {
	X int
	Y int
}

// Specifies all possible moves for a chess knight.
var knightMoves = []Move{
	{-2, -1},
	{-1, -2},
	{1, -2},
	{2, -1},
	{-2, 1},
	{-1, 2},
	{1, 2},
	{2, 1},
}

// findMoves - find the shortest sequence of valid moves
// to take a knight from the start to the target position.
func findMoves(start, target Square) Path {
	// Create queue of squares to be visited.
	queue := make(Queue, queueCapacity())

	// Enqueue the start position.
	queue.Enqueue(newPath(start))

	// Use a map to keep track of the visited squares.
	visited := map[Square]bool{}

	// Loop through the queue until it is empty.
	for !queue.IsEmpty() {
		// Pop an element, onwards called current
		current := queue.Dequeue()

		// Check if current is the target square.
		if current.HasReachedTarget(target) {
			return current
		}

		// Loop through moves.
		for _, movement := range knightMoves {
			// Apply each move to the current square.
			newPath, isValid := current.Move(movement, boardSize)
			if isValid {
				// Check if the result square was visited.
				position := newPath.LastPosition()
				if visited[position] == false {
					// Enqueue if not.
					visited[position] = true
					queue.Enqueue(newPath)
				}
			}
		}
	}

	return Path{}
}

// Calculates the queueCapacity based on the board size.
func queueCapacity() int {
	return int(math.Pow(boardSize, 2))
}
