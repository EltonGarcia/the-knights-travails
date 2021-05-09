package main

// Path stores the moves made for the knight on the board.
type Path struct {
	Moves []Square
}

// Create a new Path from a Square.
func newPath(s Square) Path {
	return Path{
		Moves: []Square{s},
	}
}

// Calculates the Distance (number of moves taken) in the path.
func (p *Path) Distance() int {
	return len(p.Moves)
}

// LastPosition gets the Path current position.
func (p *Path) LastPosition() Square {
	return p.Moves[p.Distance()-1]
}

// Check if the path has reached the target position.
func (p *Path) HasReachedTarget(target Square) bool {
	position := p.LastPosition()
	//TODO: evaluate square comparison improvements
	return position.X == target.X && position.Y == target.Y
}

// Executes a new movement in the path.
func (p *Path) Move(m Move, boardSize int) (Path, bool) {
	position := p.LastPosition()
	square := Square{X: position.X + m.X, Y: position.Y + m.Y}
	if !isSquareValid(square, boardSize) {
		return Path{}, false
	}
	newPath := Path{
		Moves: append(p.Moves, square),
	}
	return newPath, true
}

// Check if the given square is valid on the board.
func isSquareValid(s Square, boardSize int) bool {
	return s.X >= 1 &&
		s.Y >= 1 &&
		s.X <= boardSize &&
		s.Y <= boardSize
}
