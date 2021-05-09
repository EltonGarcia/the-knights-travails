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

type Path struct {
	Moves []Move
	Distance int
}

// Specifies all possible moves for a chess knight.
var moves = []Move{
	{ -2, -1 },
	{ -1, -2 },
	{ 1, -2 },
	{ 2, -1 },
	{ -2, 1 },
	{ -1, 2 },
	{ 1, 2 },
	{ 2, 1 },
}


// findMoves - find the shortest sequence of valid moves
// to take a knight from the start to the target position.
func findMoves(start, target Square){
	// Enqueue the start position.

	// Use a map to keep track of the visited squares.

	// Loop through the queue until it is empty.

	// Pop an element, onwards called current

	// Check if target was already found and distance is
	// greater than the shortest distance

	// Check if current is the target square, and
	// update the shortest distance and target found flag.

	// Loop through moves.

	// Apply each move to the current square.

	// Check if the result square was visited.
	// Enqueue if not.
}

func main(){

}