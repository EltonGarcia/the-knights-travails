package main

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
var moves = []Move{
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
func findMoves(start, target Square) []Path {
	// Create queue of squares to be visited.
	q := make(queue)

	// Enqueue the start position.
	q.Push(newPath(start))

	// Use a map to keep track of the visited squares.
	visited := [boardSize][boardSize]bool{}

	// Array to keep the solutions found.
	var results []Path
	// Function to identify if a solution was already found.
	pathFound := func() bool { return len(results) > 0 }
	// Keep track of the shortest path distance found.
	shortestDistance := 0

	// Loop through the queue until it is empty.
	for q.IsEmpty() {
		// Pop an element, onwards called current
		current := q.Pop()

		// Check if target was already found and distance is
		// greater than the shortest distance
		if pathFound() && current.Distance() >= shortestDistance {
			// If the current distance is greater or equals to the shortest
			// distance found we don't need to continue the path.
			continue
		}

		// Check if current is the target square.
		if current.HasReachedTarget(target) {
			// Update the shortest distance the first
			// time that a solution is found.
			if !pathFound() {
				shortestDistance = current.Distance()
			}
			// Add the path to the solutions array
			results = append(results, current)
			// We are at the target position,
			// we don't need to continue with the current path.
			continue
		}

		// Loop through moves.
		for _, movement := range moves {
			// Apply each move to the current square.
			newPath, isValid := current.Move(movement, boardSize)
			if isValid {
				// TODO: evaluate visited improvements
				// Check if the result square was visited.
				position := newPath.LastPosition()
				if visited[position.X][position.Y] == false {
					// Enqueue if not.
					visited[position.X][position.Y] = true
					q.Push(newPath)
				}
			}
		}
	}

	return results
}

func main() {

}
